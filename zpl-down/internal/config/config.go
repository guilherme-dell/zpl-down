package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type BarcodeConfig struct {
	Dpmm      string `json:"Dpmm"`
	Width     string `json:"Width"`
	Height    string `json:"Height"`
	ZplConfig string `json:"ZplConfig"`
	Prefix    string `json:"Prefix"`
	Index	  int32  `json:"Index"` 
	Amount    int32  `json:"Amount"`
}

type Config struct {
	APIBaseURL    string        `json:"APIBaseURL"`
	BarcodeDir    string        `json:"BarcodeDir"`
	BarcodeConfig BarcodeConfig `json:"BarcodeConfig"`
}

func LoadConfig(configFile string) Config {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}

	return config
}