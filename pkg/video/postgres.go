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
func (r *repo) SearchVideos(pagination *paginate.Pagination) (*paginate.Pagination, error) {
	var videos []models.Video
	// Apply pagination using the Paginate function
	err := r.DB.Scopes(paginate.WithWhere(&videos, "search_weights @@ plainto_tsquery(?)", pagination.Query, pagination, r.DB)).Scan(&videos).Error
	if err != nil {
		return nil, err
	}
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
