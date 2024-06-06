package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type BarcodeConfig struct {
	Dpmm      string `json:"Dpmm" validate:"required"`
	Width     string `json:"Width" validate:"required"`
	Height    string `json:"Height" validate:"required"`
	ZplConfig string `json:"ZplConfig" validate:"required"`
	Prefix    string `json:"Prefix"`
	Index     uint  `json:"Index" validate:"required"`
	Amount    uint  `json:"Amount" validate:"required"`
	PadWidth  int  `json:"PadWidth" validate:"required"`
}

type Config struct {
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
