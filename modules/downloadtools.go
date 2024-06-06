package modules

import (
	"io"
	"net/http"
	"os"
	"reddit-image-scraper/resources"
	"strings"
)

// Get the image extension
func getImageExtension(link string) string {
	// Get the last element of the link and split it by "."
	return strings.Split(link, ".")[len(strings.Split(link, "."))-1]
}

// Generate image name based
func generateImageName(description string, extension string) string {
	// Get the second last element of the description and replace all slashes with underscores
	imageName := strings.ReplaceAll(strings.Split(description, "/")[len(strings.Split(description, "/"))-2], "/", "_")
	// Append the extension to the image name
	imageName = imageName + "." + extension
	return imageName
}

// Validate file type
func validateFileType(imageLink string) bool {
	for _, imageType := range resources.Constants["imageTypes"].([]string) {
		if strings.HasSuffix(imageLink, imageType) {
			return true
		}
	}
	return false
}

// Create image_downloads directory if it does not exist
func createImageDownloadDirectory() resources.Result {
	result := resources.Result{Success: true, ErrorMessage: ""}
	_, err := os.Stat(resources.Constants["imageDownloadPath"].(string))
	if os.IsNotExist(err) {
		err := os.Mkdir(resources.Constants["imageDownloadPath"].(string), 0755)
		if err != nil {
			result.Success = false
			result.ErrorMessage = "Error creating image download directory: " + err.Error()
		}
	}
	return result
}

// Download images
func DownloadImages(imageData []resources.ImageData ) resources.Result {
	result := resources.Result{Success: true, ErrorMessage: ""}
	for _, data := range imageData {
		imageLink := data.Image
		imageName := generateImageName(data.Description, getImageExtension(imageLink))

		// Skip if image is not of the correct type
		validFile := validateFileType(imageLink)
		if !validFile {
			continue
		}
		
		// Download image
		response, err := http.Get(imageLink)
		if err != nil {
			result.Success = false
			result.ErrorMessage = "Error downloading image: " + err.Error() + " for image: " + imageName + " at link: " + imageLink
		}
		defer response.Body.Close()

		// Skip if image is empty
		if response.ContentLength == 0 {
			continue
		}

		// Create image_downloads directory if it does not exist
		createImageDownloadDirectory()

		// Create file
		file, err := os.Create(resources.Constants["imageDownloadPath"].(string) + imageName)
		if err != nil {
			result.Success = false
			result.ErrorMessage = "Error creating file: " + err.Error() + " for image: " + imageName + " at link: " + imageLink
		}
		defer file.Close()

		// Copy image data to file
		_, err = io.Copy(file, response.Body)
		if err != nil {
			result.Success = false
			result.ErrorMessage = "Error copying data to file: " + err.Error() + " for image: " + imageName + " at link: " + imageLink
		}
		result.Count++
		if result.Count == resources.Constants["imageCount"].(int) {
			break
		}
	}
	return result
}
