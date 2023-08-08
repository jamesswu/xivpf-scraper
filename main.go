package main

import (
	"context"
	"fmt"
	"os"
	db "xiv-scraper/internals/utils"

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
	coll := client.Database("xivpf").Collection("Listings")

	go db.RunCronJob(coll)
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
	app.Get("/:duty", func(c *fiber.Ctx) error {
		listings := db.GetListings(coll, c.Params("duty"))
		return c.JSON(&listings)

	})
	app.Listen(":3000")
}
