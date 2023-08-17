package video

import (
	"github.com/rishimalgwa/FamPay-Backend-Task/pkg/models"
	"github.com/rishimalgwa/FamPay-Backend-Task/pkg/paginate"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) GetAllVideos(pagination *paginate.Pagination) (*paginate.Pagination, error) {
	var videos []*models.Video
	r.DB.Scopes(paginate.Paginate(videos, pagination, r.DB)).Find(&videos)
	// check length if exceeds the pagination limit
	if len(videos) > pagination.Limit {
		newTrue := true
		pagination.IfNext = &newTrue
		// trim the extra record
		pagination.Rows = videos[:pagination.Limit]
	} else {
		newFalse := false
		pagination.IfNext = &newFalse
		pagination.Rows = videos
	}
	return pagination, nil
}
func (r *repo) SearchVideos(pagination *paginate.Pagination) ([]models.Video, int, error) {
	var videos []models.Video
	var count int64

	// Build the base query
	baseQuery := r.DB.Model(&models.Video{}).Where("search_doc_weights @@ plainto_tsquery(?)", pagination.Query)

	// Apply sorting
	baseQuery = baseQuery.Order(pagination.Sort)

	// Calculate offset based on pagination
	//offset := (pagination.Page - 1) * pagination.Limit

	// Apply pagination using the Paginate function
	r.DB.Scopes(paginate.Paginate(&videos, pagination, baseQuery))

	// Get total count for pagination
	if err := baseQuery.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	pageCount := int((count + int64(pagination.Limit) - 1) / int64(pagination.Limit))
	return videos, pageCount, nil
}
