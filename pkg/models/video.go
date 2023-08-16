package models

import (
	"time"
)

type Video struct {
	BaseModel
	Name         string    `json:"name"`
	PublishedAt  string    `json:"published_at"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	ChannelId    string    `json:"channel_id"`
	Description  string    `json:"description"`
	ChannelName  string    `json:"channel_name"`
	CreatedDate  time.Time `json:"created_date"`
}
