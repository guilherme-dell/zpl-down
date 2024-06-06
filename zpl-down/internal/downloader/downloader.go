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

		fileName := createFileName(index, cfg)
		filePath := fmt.Sprintf("%s/%s.png", dir, fileName)

		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("Fail to create file:%v", err)
		}

		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatalf("Fail to save file:%v", err)
		}

		fmt.Printf("BARCODE_%s\t\t\033[0;32m[OK]\033[0m\n", fileName)

		index++
		time.Sleep(200 * time.Millisecond)
	}
}

func UrlMount(barcodeConfig config.BarcodeConfig, indice uint) string {
	format := fmt.Sprintf("http://api.labelary.com/v1/printers/%%s/labels/%%sx%%s/0/ --data-urlencode \"%%s%%0%dd%%s^fs^xz\"", barcodeConfig.PadWidth)
	return fmt.Sprintf(format, barcodeConfig.Dpmm, barcodeConfig.Width, barcodeConfig.Height, barcodeConfig.ZplConfig, indice, barcodeConfig.Prefix)
}

func createFileName(index uint, cfg config.Config) string {
	format := fmt.Sprintf("%%0%dd%%s", cfg.BarcodeConfig.PadWidth)
	return fmt.Sprintf(format, index, cfg.BarcodeConfig.Prefix)
}
