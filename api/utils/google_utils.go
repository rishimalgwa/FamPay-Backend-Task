package utils

import (
	"context"
	"log"
	"time"

	"github.com/rishimalgwa/FamPay-Backend-Task/api/db"
	"github.com/rishimalgwa/FamPay-Backend-Task/pkg/models"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"gorm.io/gorm"
)

// func FetchVideos() {
// 	// impl multi api keys
// 	API_KEY := viper.GetString("YT_API_KEY1")
// 	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?key=%s&q=%s&type=video&order=date&part=snippet&publishedAfter=%s", API_KEY, searchQuery, publishedAfter)
// }

var (
	number      int
	youtubeNext string
	baseDate    = time.Date(2022, time.January, 1, 12, 0, 0, 0, time.UTC)
	apiKey      = ""
)

func GetYouTubeVideos(query string) error {
	// Load YouTube API key from configuration
	apiKey := viper.GetString("YT_API_KEY1")
	if apiKey == "" {
		log.Fatal("YOUTUBE_API_KEY environment variable not set")
	}

	// Create a background context and initialize YouTube service
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
		return err
	}

	// Define API request parts
	parts := [2]string{"id", "snippet"}
	var searchResponse *youtube.SearchListResponse

	// Build API request based on the current state
	if number == 0 {
		searchResponse, err = service.Search.List(parts[:]).
			Q(query).
			MaxResults(50).PublishedAfter(baseDate.Format(time.RFC3339)).
			Type("video").
			Do()
	} else {
		searchResponse, err = service.Search.List(parts[:]).
			Q(query).
			MaxResults(50).PageToken(youtubeNext).
			Type("video").
			Do()
	}
	if err != nil {
		log.Fatalf("Error calling YouTube Data API: %v", err)
		return err
	}

	number++
	if searchResponse.NextPageToken != "" {
		youtubeNext = searchResponse.NextPageToken
	}

	// Iterate over search results and save videos
	for _, v := range searchResponse.Items {
		if err := saveVideos(db.GetDB(), v); err != nil {
			log.Printf("Error saving video: %v", err)
		}
	}

	return nil
}

func saveVideos(db *gorm.DB, video *youtube.SearchResult) error {
	// Prepare data to be saved
	toBeSaved := &models.Video{
		Name:         video.Snippet.Title,
		Description:  video.Snippet.Description,
		PublishedAt:  video.Snippet.PublishedAt,
		ThumbnailUrl: video.Snippet.Thumbnails.Default.Url,
		ChannelId:    video.Snippet.ChannelId,
		ChannelName:  video.Snippet.PublishedAt,
	}

	// Begin a transaction
	tx := db.Begin()
	if err := tx.Create(&toBeSaved).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update search weights
	if err := tx.Exec(`
		UPDATE videos
		SET search_weights = setweight(to_tsvector(name), 'A') || setweight(to_tsvector(COALESCE(description, '')), 'B')
	`).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Drop existing index
	if err := tx.Exec(`DROP INDEX IF EXISTS search__weights_idx`).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create new GIN index
	if err := tx.Exec(`
		CREATE INDEX search__weights_idx ON videos USING GIN(search_weights)
	`).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
