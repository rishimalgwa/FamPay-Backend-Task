package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
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
	apiKey      = os.Getenv("YT_KEY1")
)

func GetYouTubeVideos(query string) ([]*youtube.SearchResult, error) {

	if apiKey == "" {
		log.Fatal("YOUTUBE_API_KEY environment variable not set")
	}

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}

	parts := [2]string{"id", "snippet"}
	var searchResponse *youtube.SearchListResponse
	if number == 0 {
		searchResponse, err = service.Search.List(parts[:]).
			Q(query).
			MaxResults(50).PublishedAfter(baseDate.UTC().String()).
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
		return nil, err
	}
	fmt.Println(searchResponse)
	return searchResponse.Items, nil
}
