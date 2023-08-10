package router

import (
	"context"
	"xiv-scraper/internals/ffxiv"
	"xiv-scraper/internals/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddListingGroup(app *fiber.App) {
	listingsGroup := app.Group("/listings")
	listingsGroup.Get("/:duty", getListings)
}

func getListings(c *fiber.Ctx) error {
	coll := utils.GetDBCollection("Listings")
	duty := c.Params("duty")

	if duty == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is required"})
	}

	d := ffxiv.DutyHandler(duty)

	listings := make([]ffxiv.Listing, 0)
	filter := bson.M{"duty": d}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	for cursor.Next(c.Context()) {
		listing := ffxiv.Listing{}
		err := cursor.Decode(&listing)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		listings = append(listings, listing)
	}
	return c.Status(200).JSON(fiber.Map{"data": listings})
}
