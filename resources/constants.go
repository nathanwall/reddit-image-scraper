package resources

import (
	"os"
)

type ImageData struct {
	Image string
	Description string
}

type Result struct {
	Success bool
	ErrorMessage string
	Count int
}

func setConstants() map[string]interface{} {
	var config Config = loadConfig()
	var PWD string

	if config.UsePWD {
		PWD, _ = os.Getwd()
	} else {
		PWD = ""
	}

	constants := map[string]interface{} {
		"redditURL": config.RedditURL,
		"imageDownloadPath": PWD + config.ImageDownloadPath,
		"imageTypes": config.ImageTypes,
		"HTMLElement": config.HTMLElement,
		"imageAttribute": config.ImageAttribute,
		"imageDescription": config.ImageDescription,
		"imageCount": config.ImageCount,
	}

	return constants
}

var Constants = setConstants()