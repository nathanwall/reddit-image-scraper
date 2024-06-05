package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Main function
func main() {
	// Create a new collector
	collector := colly.NewCollector()

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
	collector.OnHTML(Constants["HTMLElement"].(string), func(e *colly.HTMLElement) {
		var imageData []ImageData = []ImageData{
			{
				image: e.Attr(Constants["imageAttribute"].(string)), 
				description: e.Attr(Constants["imageDescription"].(string)),
			},
		}

		// Download images. If there is an error, log the error
		success := downloadImages(imageData)
		if !success.success {
			fmt.Println("Error downloading image:", success.errorMessage)
		}
	})

	// When scraping is finished, log the message
	collector.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Visit the reddit URL
	collector.Visit(Constants["redditURL"].(string))
}