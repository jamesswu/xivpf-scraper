package scraper

import (
	"fmt"
	"log"
	"strings"
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
		listing := &ffxiv.Listing{Party: []*ffxiv.Slot{}}
		e.Unmarshal(listing)

		listing.DataCenter = e.Attr("data-centre")
		description := e.ChildText(".left .description")
		description = strings.TrimSpace(strings.Replace(description, listing.Tags, "", -1))
		listing.Description = description

		e.ForEach(".party .slot", func(s int, p *colly.HTMLElement) {
			slot := ffxiv.NewSlot()
			class := p.Attr("class")

			if strings.Contains(class, "dps") {
				slot.Roles.Roles = append(slot.Roles.Roles, ffxiv.Dps)
			}
			if strings.Contains(class, "healer") {
				slot.Roles.Roles = append(slot.Roles.Roles, ffxiv.Healer)
			}
			if strings.Contains(class, "tank") {
				slot.Roles.Roles = append(slot.Roles.Roles, ffxiv.Tank)
			}
			if strings.Contains(class, "empty") {
				slot.Roles.Roles = append(slot.Roles.Roles, ffxiv.Empty)
			}
			if strings.Contains(class, "filled") {
				slot.Filled = true
				slot.Job = ffxiv.GetJob(p.Attr("title"))
			}
			listing.Party = append(listing.Party, slot)
		})
		listings.Add(listing)
	})
	c.Visit("https://xivpf.com/listings")
	s.Listings = listings
	return nil
}
