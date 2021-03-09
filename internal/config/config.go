package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Host     string   `json:"host"`
	Port     string   `json:"port"`
	Database Database `json:"database"`
}

type Database struct {
	Path   string `json:"path"`
	Driver string `json:"driver"`
}

func Load(filePath string) (config Config, err error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	if err = json.Unmarshal(file, &config); err != nil {
		return
	}

	return
}
