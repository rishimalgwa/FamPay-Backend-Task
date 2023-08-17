package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rishimalgwa/FamPay-Backend-Task/api/db"
	"github.com/rishimalgwa/FamPay-Backend-Task/api/migrations"
	"github.com/rishimalgwa/FamPay-Backend-Task/api/router"
	"github.com/rishimalgwa/FamPay-Backend-Task/api/utils"
	"github.com/rishimalgwa/FamPay-Backend-Task/config"
	"github.com/robfig/cron"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

var c *cron.Cron // Declare a global variable for the cron job

func main() {
	// Set global configuration
	utils.ImportEnv()

	// load configs
	config.LoadConfig()

	// Init Validators
	utils.InitValidators()

	// Create Fiber
	app := fiber.New(fiber.Config{})

	app.Get("/", healthCheck)
	app.Get("/health", healthCheck)

	app.Use(logger.New(logger.Config{Next: func(c *fiber.Ctx) bool {
		return strings.HasPrefix(c.Path(), "api")
	}}))

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "",
		AllowHeaders: "*",
	}))

	//Connect and migrate the db
	//if viper.GetBool("MIGRATE") {
	migrations.Migrate()
	//}

	// Initialize DB
	db.InitServices()

	// Mount Routes
	router.MountRoutes(app)

	// Get Port
	port := utils.GetPort()

	utils.InitKeyManager()

	// Create a new cron job
	c = cron.New()

	// Define a route to start the cron job
	app.Get("/startCron", func(ctx *fiber.Ctx) error {
		// Add a cron job to run every 10 seconds
		c.AddFunc("@every 10s", func() {
			// Place your code here that you want to run every 10 seconds
			fmt.Println("Cron job executed at:", time.Now())
			err := utils.GetYouTubeVideos("F1")
			if err != nil {
				log.Println(err)
			}
		})
		c.Start() // Start the cron job
		return ctx.SendString("Cron job started")
	})

	// Define a route to stop the cron job
	app.Get("/stopCron", func(ctx *fiber.Ctx) error {
		c.Stop() // Stop the cron job
		return ctx.SendString("Cron job stopped")
	})

	// Start Fiber
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

}
