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
	apiKey = viper.GetString("YT_API_KEY1")
	if apiKey == "" {
		log.Fatal("YOUTUBE_API_KEY environment variable not set")
	}

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
		return err
	}

	parts := [2]string{"id", "snippet"}
	var searchResponse *youtube.SearchListResponse
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
	for _, v := range searchResponse.Items {

		saveVideos(db.GetDB(), v, number)
	}
	return nil
}

func saveVideos(db *gorm.DB, video *youtube.SearchResult, num int) (*models.Video, error) {

	toBeSaved := &models.Video{
		Name:         video.Snippet.Title,
		Description:  video.Snippet.Description,
		PublishedAt:  video.Snippet.PublishedAt,
		ThumbnailUrl: video.Snippet.Thumbnails.Default.Url,
		ChannelId:    video.Snippet.PublishedAt,
		ChannelName:  video.Snippet.PublishedAt,
	}
	tx := db.Begin()

	// if num == 0 {
	// 	if err := tx.Exec("DELETE FROM videos").Error; err != nil {
	// 		tx.Rollback()
	// 		return nil, err
	// 	}
	// }

	if err := tx.Create(&toBeSaved).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Exec(`
		UPDATE videos
		SET search_doc_weights = setweight(to_tsvector(name), 'A') || setweight(to_tsvector(COALESCE(description, '')), 'B')
	`).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Exec(`DROP INDEX IF EXISTS search__weights_idx`).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Exec(`
		CREATE INDEX search__weights_idx ON videos USING GIN(search_doc_weights gin_trgm_ops)
	`).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return toBeSaved, nil
}
