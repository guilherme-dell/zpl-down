package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/guilherme-dell/zpl-down/internal/config"
	"github.com/guilherme-dell/zpl-down/internal/downloader"
)

var GREEN = "\033[0;32m"
var RED = "\033[0;31m"
var YELLOW = "\033[0;33m"
var RESET = "\033[0m"

func CreateDir(configFile config.Config) {

	dir := configFile.BarcodeDir

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf("Failed to creater dir: %v", err)
		}
	}
}

func PrinterFILE(cfg config.Config) {
	fmt.Printf("%s[CONFIG]%s\n", GREEN, RESET)
	fmt.Printf(" %s-DIR:%s➜%s\n", RED, RESET, cfg.BarcodeDir)
	fmt.Printf("%s[BARCODE-CONFIG]%s\n", GREEN, RESET)
	fmt.Printf(" %s-DPMM%s➜%s\n", RED, RESET, cfg.BarcodeConfig.Dpmm)
	fmt.Printf(" %s-WIDTH%s➜%s\n", RED, RESET, cfg.BarcodeConfig.Width)
	fmt.Printf(" %s-HEIGHT%s➜%s\n", RED, RESET, cfg.BarcodeConfig.Height)
	fmt.Printf(" %s-ZPL_CONFIG%s➜%s\n", RED, RESET, cfg.BarcodeConfig.ZplConfig)
	fmt.Printf(" %s-PREFIX%s➜%s\n", RED, RESET, cfg.BarcodeConfig.Prefix)
	fmt.Printf(" %s-INDEX%s➜%d\n", RED, RESET, cfg.BarcodeConfig.Index)
	fmt.Printf(" %s-AMOUNT%s➜%d\n", RED, RESET, cfg.BarcodeConfig.Amount)
	fmt.Printf(" %s-PADWIDTH%s➜%d\n", RED, RESET, cfg.BarcodeConfig.PadWidth)
}

func PrinterURL(cfg config.Config) {

	URL := downloader.UrlMount(cfg.BarcodeConfig, 1)
	fmt.Printf("%s[URL VARs]%s\n", GREEN, RESET)
	fmt.Printf(" %s-PATH_URL%s➜%s%s%s\n", RED, RESET, YELLOW, URL, RESET)
	fmt.Printf(" %s-ZPL_CONFIG%s➜%s%s%s\n", RED, RESET, YELLOW, cfg.BarcodeConfig.ZplConfig, RESET)
}

func PrintSucces(index uint, cfg config.Config) {
	format := fmt.Sprintf("BARCODE_%%0%dd%%s\t\t\033[0;32m[OK]\033[0m\n", cfg.BarcodeConfig.PadWidth)
	fmt.Printf(format, index, cfg.BarcodeConfig.Prefix)
}
