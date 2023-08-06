package main

import (
	"context"
	"fmt"
	"os"
	"xiv-scraper/internals/ffxiv"
	"xiv-scraper/internals/scraper"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Println("you must set your 'MONGODB_URI' environment variable")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// create new client and connect to server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// s := gocron.NewScheduler(time.UTC)
	// _, er := s.Every(3).Minutes().Do(func() {
	scraper := scraper.New("https://xivpf.com/listings")
	fmt.Println("starting scraper ...")
	fmt.Println("scraping ...")
	errr := scraper.Scrape()
	if errr != nil {
		fmt.Printf("scraper error: %f\n", err)
	}
	listings := scraper.Listings.GetUltimateListings(scraper.Listings)
	coll := client.Database("xivpf").Collection("Listings")
	// collection2 := client.Database("xivpf").Collection("Party")
	// docs := []interface{}{}
	for _, l := range listings.Listings {
		listing := ffxiv.Listing{
			DataCenter:  l.DataCenter,
			Duty:        l.Duty,
			Description: l.Description,
			Creator:     l.Creator,
			World:       l.World,
			Expires:     l.Expires,
			Updated:     l.Updated,
			Party:       l.Party,
		}
		result, insertErr := coll.InsertOne(context.TODO(), listing)
		if insertErr != nil {
			fmt.Println("insert error")
		} else {
			fmt.Println(result)
		}
	}
	// result, err := coll.InsertMany(context.TODO(), docs)

	// clean database

	// update database with new listings

	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
	// app.Get("/:duty", func(c *fiber.Ctx) error {
	// 	listings := scraper.Listings.GetListings(scraper.Listings, c.Params("duty"))
	// 	return c.JSON(listings)
	// })
	app.Listen(":3000")
}
