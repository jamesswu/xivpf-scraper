package main

import (
	"os"
	"xiv-scraper/internals/router"
	"xiv-scraper/internals/scheduler"
	"xiv-scraper/internals/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// init env
	err := utils.LoadEnv()
	if err != nil {
		panic(err)
	}
	// init db
	err = utils.InitDB()
	if err != nil {
		panic(err)
	}
	// defer closing db
	defer utils.CloseDB()

	// run cronjob
	go scheduler.RunCronJob()

	// create app
	app := fiber.New()

	// middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	router.AddListingGroup(app)

	// start server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
}
