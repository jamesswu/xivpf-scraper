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
	app.Get("/:duty", func(c *fiber.Ctx) error {
		listings := scraper.Listings.GetListings(scraper.Listings, c.Params("duty"))
		for _, l := range listings.Listings {
			fmt.Println(l.Creator)
		}
		return c.SendString(c.Params("duty"))
	})
	app.Listen(":3000")
}
