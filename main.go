package main

import (
	"fmt"
	"xiv-scraper/internals/scraper"

	"github.com/gofiber/fiber/v2"
)

func main() {
	scraper := scraper.New("https://xivpf.com/listings")
	fmt.Println("starting scraper ...")
	fmt.Println("scraping ...")
	err := scraper.Scrape()
	if err != nil {
		fmt.Printf("scraper error: %f\n", err)
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
	app.Get("/listings", func(c *fiber.Ctx) error {
		return c.SendString("scraping")
	})
	app.Listen(":3000")
}
