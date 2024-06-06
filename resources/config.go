package resources

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
 	RedditURL string `yaml:"redditURL"`
 	ImageDownloadPath string `yaml:"imageDownloadPath"`
 	UsePWD bool `yaml:"usePWD"`
 	ImageTypes []string `yaml:"imageTypes"`
 	HTMLElement string `yaml:"HTMLElement"`
 	ImageAttribute string `yaml:"imageAttribute"`
 	ImageDescription string `yaml:"imageDescription"`
	ImageCount int `yaml:"imageCount"`
}

type ConfigWrapper struct {
	Config Config `yaml:"config"`
}

func loadConfig() Config {
	var pwd, _ = os.Getwd()
	file, err := os.ReadFile(pwd + "/config.yaml")
  if err != nil {
  	log.Fatal(err)
  }

  var configWrapper ConfigWrapper
  err = yaml.Unmarshal(file, &configWrapper)
  if err != nil {
  	log.Fatal(err)
  }

	return configWrapper.Config
}
