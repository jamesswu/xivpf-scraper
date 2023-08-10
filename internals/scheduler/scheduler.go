package scheduler

import (
	"context"
	"fmt"
	"log"
	"time"
	"xiv-scraper/internals/ffxiv"
	"xiv-scraper/internals/scraper"
	"xiv-scraper/internals/utils"

	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson"
)

func Cleanup() {
	coll := utils.GetDBCollection("Listings")

	res, err := coll.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("deleted %v documents\n", res.DeletedCount)
	}
}

func PostListing(s *scraper.Scraper) {
	coll := utils.GetDBCollection("Listings")
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
		_, insertErr := coll.InsertOne(context.TODO(), listing)
		if insertErr != nil {
			fmt.Println("insert error")
		}
	}
	fmt.Println("added new documents")
}

func RunCronJob() {

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
		Cleanup()
		// update database with new listings
		PostListing(scraper)

	})
	sch.StartAsync()
}
