package migrations

import (
	"log"

	"github.com/rishimalgwa/FamPay-Backend-Task/api/db"
	"github.com/rishimalgwa/FamPay-Backend-Task/pkg/models"
)

func Migrate() {
	database := db.GetDB()
	database.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	err := database.AutoMigrate(&models.Video{})
	if err != nil {
		log.Fatalln("Cannot Migrate: ", err)
		return
	}
	// Mark Migrations as complete
	log.Println("Migrations Completed")
	return
}
