package main

import (
	"fmt"

	"github.com/guilherme-dell/zpl-down/internal/config"
)

func main() {
	//var barcode zpl.Barcode

	//barcode.Dpmm = "8dpmm"
	//barcode.Width = "0.98"
	//barcode.Height = "0.39"
	//barcode.Prefix = "CJ"
	//barcode.GenerateAmount = 1

	//barcode.Download()

	cfg := config.LoadConfig("config_1.json")
	
	fmt.Println(cfg)
}
