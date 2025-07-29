package config

import (
	"os"
	"log"
	"gopkg.in/yaml.v3"
)

type config struct {
	Title string `yaml:"title"`
	Favicon string `yaml:"favicon"`
	Theme string `yaml:"theme"`
	BaseUrl string `yaml:"baseurl"`
	Mode string `yaml:"mode"`
}

var Config config

func Load() {
	data, err := os.ReadFile("data/config.yaml")
	if err != nil {
		log.Fatalf("error reading config file : %v", err)
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalf("error unmarshaling file yaml : %v", err)
	}
}
