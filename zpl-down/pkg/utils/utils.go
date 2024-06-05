package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/guilherme-dell/zpl-down/internal/config"
)

var GREEN = "\033[0;32m"
var RED = "\033[0;31m"
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
	fmt.Printf("%s[CONFIG]%s\n", GREEN,RESET)
	fmt.Printf(" %sDIR:%s➜%s\n",RED, RESET, cfg.BarcodeDir)
	fmt.Printf("%s[BARCODE-CONFIG]%s\n", GREEN,RESET)
	fmt.Printf(" %sDPMM%s➜%s\n",RED, RESET, cfg.BarcodeConfig.Dpmm)
	fmt.Printf(" %sWIDTH%s➜%s\n",RED, RESET, cfg.BarcodeConfig.Width)
	fmt.Printf(" %sHEIGHT%s➜%s\n",RED, RESET, cfg.BarcodeConfig.Height)
	fmt.Printf(" %sZPL_CONFIG%s➜%s\n",RED, RESET, cfg.BarcodeConfig.ZplConfig)
	fmt.Printf(" %sPREFIX%s➜%s\n",RED, RESET, cfg.BarcodeConfig.Prefix)
	fmt.Printf(" %sINDEX%s➜%d\n",RED, RESET, cfg.BarcodeConfig.Index)
	fmt.Printf(" %sAMOUNT%s➜%d\n",RED, RESET, cfg.BarcodeConfig.Amount)
}
