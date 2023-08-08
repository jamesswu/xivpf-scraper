# xiv-scraper

inspiration taken from
https://github.com/Veraticus/trappingway
and
https://github.com/epitaque/trappingway

### what it does

runs a scheduled CronJob that scrapes the party finder listings from https://xivpf.com/listings then storing the data in a MongoDB Cluster for processing and retrieval.

### how to deploy

need MONGODB_URI in an `.env` file
navigate to directory and > `go run main.go`
