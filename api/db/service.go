package db

import "github.com/rishimalgwa/FamPay-Backend-Task/pkg/video"

var (
	VideoSvc video.Service = nil
)

func InitServices() {
	db := GetDB()

	videoRepo := video.NewPostgresRepo(db)
	VideoSvc = video.NewService(videoRepo)
}
