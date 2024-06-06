# Reddit Image Scraper
Scrape images from Reddit. This script will download images from a subreddit and save them to a folder on your computer. I have this one set to Cool Guides by default.

This project uses old.reddit.com to scrape images.

## Installation
Clone the repository.

## Usage
`go run .` or `./reddit-image-scraper` after building the binary file.

## Configuration
Configuration is done through the `config.yaml` file.
- `redditURL` - The URL of the subreddit you want to scrape images from. By default, it is set to r/coolguides and looks for top posts of the day.
- `imageDownloadPath` - The path to the folder you want to save the images to. By default, it is set to the `image_downloads` folder in the project directory.
- `usePWD` - Use the working directory for the image download path for the `imageDownloadPath`. By default, it is set to true.
- `imageTypes` - The types of images you want to download. By default, it is set to jpg, jpeg, and png.
- `HTMLElement` - The HTML element that contains the image URL. By default, it is set to `div.thing`.
- `imageAttribute` - The attribute that contains the image URL. By default, it is set to `data-url`.
- `imageDescription` - The attribute that contains the image description, which is used for the image name. By default, it is set to `data-permalink`.
- `imageCount` - The number of images you want to download. By default, it is set to `-1` (max).

## Build
`go build` to build an executable binary file.