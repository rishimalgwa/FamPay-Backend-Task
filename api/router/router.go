package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rishimalgwa/FamPay-Backend-Task/api/controllers"
)

func MountRoutes(c *fiber.App) {
	thisController := controllers.VideoController{}
	videoGroup := c.Group("/videos")
	videoGroup.Get("/get", thisController.GetAllVideos)
	videoGroup.Get("/search", thisController.SearchVideo)
}
