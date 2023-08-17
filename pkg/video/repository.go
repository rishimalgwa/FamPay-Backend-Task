package video

import (
	"github.com/rishimalgwa/FamPay-Backend-Task/pkg/models"
	"github.com/rishimalgwa/FamPay-Backend-Task/pkg/paginate"
)

type Repository interface {
	GetAllVideos(p *paginate.Pagination) (*paginate.Pagination, error)
	SearchVideos(pagination *paginate.Pagination) ([]models.Video, error)
}
