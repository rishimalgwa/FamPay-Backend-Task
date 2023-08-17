package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Video struct {
	BaseModel
	Name             string `json:"name"`
	PublishedAt      string `json:"published_at"`
	ThumbnailUrl     string `json:"thumbnail_url"`
	ChannelId        string `json:"channel_id"`
	Description      string `json:"description"`
	ChannelName      string `json:"channel_name"`
	SearchDocWeights string `gorm:"column:search_doc_weights"`
}

func (v *Video) BeforeCreate(_ *gorm.DB) error {
	v.ID = uuid.New()
	return nil
}
