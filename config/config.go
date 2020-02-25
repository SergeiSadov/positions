package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	DbPath   string  `json:"dbPath"`
	DbDriver string  `json:"dbDriver"`
	Port     string  `json:"port"`
	Elastic  Elastic `json:"elastic"`
}

type Elastic struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Config config

func Load(filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err == nil {
		return json.Unmarshal(file, &Config)
	}

	return fmt.Errorf("config.Load() error #1: \t\n %w \n", err)
}
