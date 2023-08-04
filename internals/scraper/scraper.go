package scraper

import (
	"fmt"
	"log"
	"xiv-scraper/internals/ffxiv"

	"github.com/gocolly/colly"
)

type Scraper struct {
	Url      string
	Listings *ffxiv.Listings
}

func New(url string) *Scraper {
	return &Scraper{
		Url:      url,
		Listings: &ffxiv.Listings{Listings: []*ffxiv.Listing{}},
	}
}

func (s *Scraper) Scrape() error {
	listings := &ffxiv.Listings{}

	c := colly.NewCollector()

	// logging and error handling
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// Find and visit all links
	c.OnHTML(".listing", func(e *colly.HTMLElement) {
		e.Unmarshal()
	})
	c.Visit("https://xivpf.com/listings")

}
