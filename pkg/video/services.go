package video

import (
	"github.com/rishimalgwa/FamPay-Backend-Task/pkg/paginate"
)

type Service interface {
	GetAllVideos(p *paginate.Pagination) (*paginate.Pagination, error)
	SearchVideos(pagination *paginate.Pagination) (*paginate.Pagination, error)
}

type videoSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &videoSvc{repo: r}
}
func (c *videoSvc) GetAllVideos(p *paginate.Pagination) (*paginate.Pagination, error) {
	return c.repo.GetAllVideos(p)
}
func (c *videoSvc) SearchVideos(pagination *paginate.Pagination) (*paginate.Pagination, error) {
	return c.repo.SearchVideos(pagination)
}
