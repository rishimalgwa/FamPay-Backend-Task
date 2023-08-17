package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rishimalgwa/FamPay-Backend-Task/api/db"
	"github.com/rishimalgwa/FamPay-Backend-Task/api/views"
	"github.com/rishimalgwa/FamPay-Backend-Task/pkg/paginate"
)

type VideoController struct {
}

// GetAllVideos Get all videos from db
func (v *VideoController) GetAllVideos(c *fiber.Ctx) error {
	// Get all orders
	pagination := new(paginate.Pagination)
	if err := c.QueryParser(pagination); err != nil {
		return views.InvalidQuery(c, err)
	}

	p, err := db.VideoSvc.GetAllVideos(pagination)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.OK(c, p)
}

// SearchVideo search videos for given query
func (v *VideoController) SearchVideo(c *fiber.Ctx) error {
	// Get all orders
	pagination := new(paginate.Pagination)
	if err := c.QueryParser(pagination); err != nil {
		return views.InvalidQuery(c, err)
	}

	p, _, err := db.VideoSvc.SearchVideos(pagination)
	if err != nil {
		return views.InternalServerError(c, err)
	}
	return views.OK(c, p)
}
