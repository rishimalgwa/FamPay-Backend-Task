package migrations

import (
	"log"

	"github.com/rishimalgwa/FamPay-Backend-Task/api/db"
	"github.com/rishimalgwa/FamPay-Backend-Task/pkg/models"
)

func Migrate() {
	database := db.GetDB()
	database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	err := database.AutoMigrate(&models.Video{})
	if err != nil {
		log.Fatalln("Cannot Migrate: ", err)
		return
	}
	database.Raw("CREATE INDEX search__weights_idx ON videos USING GIN(search_weights);")
	// Mark Migrations as complete
	log.Println("Migrations Completed")
	return
}
