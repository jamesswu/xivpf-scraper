package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"xiv-scraper/internals/ffxiv"
	"xiv-scraper/internals/scraper"

	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

func Cleanup(c *mongo.Collection) {
	res, err := c.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("deleted %v documents\n", res.DeletedCount)
	}
}

func PostListing(c *mongo.Collection, s *scraper.Scraper) {
	listings := s.Listings.GetUltimateListings(s.Listings)
	for _, l := range listings.Listings {
		listing := ffxiv.Listing{
			DataCenter:  l.DataCenter,
			Duty:        l.Duty,
			Tags:        l.Tags,
			Description: l.Description,
			Creator:     l.Creator,
			World:       l.World,
			Expires:     l.Expires,
			Updated:     l.Updated,
			Party:       l.Party,
		}
		_, insertErr := c.InsertOne(context.TODO(), listing)
		if insertErr != nil {
			fmt.Println("insert error")
		}
	}
	fmt.Println("added new documents")
}

func GetListings(c *mongo.Collection, duty string) []primitive.M {
	d := ffxiv.DutyHandler(duty)
	filter := bson.M{"duty": d}
	cursor, err := c.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var listings []bson.M
	if err = cursor.All(context.TODO(), &listings); err != nil {
		panic(err)
	}
	return listings
}

func RunCronJob(c *mongo.Collection) {

	sch := gocron.NewScheduler(time.UTC)
	sch.Every(3).Minutes().Do(func() {
		scraper := scraper.New("https://xivpf.com/listings")
		fmt.Println("starting scraper ...")
		fmt.Println("scraping ...")
		err := scraper.Scrape()
		if err != nil {
			fmt.Printf("scraper error: %f\n", err)
		}

		// clean database
		Cleanup(c)
		// update database with new listings
		PostListing(c, scraper)

	})
	sch.StartBlocking()
}
