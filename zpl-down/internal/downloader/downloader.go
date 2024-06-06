package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/guilherme-dell/zpl-down/internal/config"
	
)



func DownloadBarCodes(barcodeConfig config.BarcodeConfig, cfg config.Config) {
	index := barcodeConfig.Index
	dir := cfg.BarcodeDir

	for index <= barcodeConfig.Amount {

		URL := UrlMount(barcodeConfig, index)

		response, err := http.Get(URL)
		if err != nil {
			log.Fatalf("Failed to try request barcode: %v", err)
		}

		if response.StatusCode != 200 {
			log.Fatalf("Request failed - status code::%v", response.StatusCode)
		}

		fileName := fmt.Sprintf("%s/%04d%s.png", dir, index, barcodeConfig.Prefix)

		file, err := os.Create(fileName)
		if err != nil {
			log.Fatalf("Fail to create file:%v", err)
		}

		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatalf("Fail to save file:%v", err)
		}

		printPretty(index, barcodeConfig.PadWidth, barcodeConfig.Prefix)
		
		index++
		time.Sleep(200 * time.Millisecond)
	}
}

func UrlMount(barcodeConfig config.BarcodeConfig, indice uint) string {
	format := fmt.Sprintf("http://api.labelary.com/v1/printers/%%s/labels/%%sx%%s/0/ --data-urlencode \"%%s%%0%dd%%s^fs^xz\"", barcodeConfig.PadWidth)

	return fmt.Sprintf(format, barcodeConfig.Dpmm, barcodeConfig.Width, barcodeConfig.Height, barcodeConfig.ZplConfig, indice, barcodeConfig.Prefix)
}

func printPretty (index uint, padWidth int, prefix string) {
	format := fmt.Sprintf("BARCODE_%%0%dd%%s\t\t\033[0;32m[OK]\033[0m\n", padWidth)
	fmt.Printf(format, index, prefix)
}