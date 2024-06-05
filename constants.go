package main

import (
	"os"
)

var PWD, _ = os.Getwd()

var Constants = map[string]interface{} {
	"redditURL": "https://old.reddit.com/r/coolguides/top/?t=day",
	"imageDownloadPath": PWD + "/image_downloads/",
	"imageTypes": []string {".jpg", ".jpeg", ".png"},
	"HTMLElement": "div.thing",
	"imageAttribute": "data-url",
	"imageDescription": "data-permalink",
}

type ImageData struct {
	image string
	description string
}

type Result struct {
	success bool
	errorMessage string
}