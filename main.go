package main

import (
	"fmt"
	"reddit-image-scraper/modules"
	"reddit-image-scraper/resources"

	"github.com/gocolly/colly"
)

// Main function
func main() {
	// Create a new collector
	collector := colly.NewCollector()
	var success resources.Result
	var imageData []resources.ImageData

	// On every request, visit the URL
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// On error, log the error
	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// On response, log the response
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	// When HTML is received, get the image and description and download the image
	collector.OnHTML(resources.Constants["HTMLElement"].(string), func(e *colly.HTMLElement) {
		imageData = append(imageData, resources.ImageData{
				Image: e.Attr(resources.Constants["imageAttribute"].(string)),
				Description: e.Attr(resources.Constants["imageDescription"].(string)),
			})
	})

	// When scraping is finished, log the message
	collector.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Visit the reddit URL
	collector.Visit(resources.Constants["redditURL"].(string))

	success = modules.DownloadImages(imageData)
	if !success.Success {
		fmt.Println("Error downloading images:", success.ErrorMessage)
	}	else {
		fmt.Println("Downloaded", success.Count, "images")
	}
}