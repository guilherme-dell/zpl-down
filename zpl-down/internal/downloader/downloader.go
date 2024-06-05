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

func urlMount(barcodeConfig config.BarcodeConfig, indice int32) string {
	return fmt.Sprintf("http://api.labelary.com/v1/printers/%s/labels/%sx%s/0/ --data-urlencode \"^xa^fo4,15^be1^bcn,40,,,,a^fd%04d%s^fs^xz\"",
		barcodeConfig.Dpmm, barcodeConfig.Width, barcodeConfig.Height, indice, barcodeConfig.Prefix)
}

func DownloadBarCodes(barcodeConfig config.BarcodeConfig, cfg config.Config) {
	index := barcodeConfig.Index
	dir := cfg.BarcodeDir

	for index <= barcodeConfig.Amount {

		URL := urlMount(barcodeConfig, index)

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

		fmt.Printf("BARCODE_%04d%s\t\t\033[0;32m[OK]\033[0m\n", index, barcodeConfig.Prefix)
		index++
		time.Sleep(200 * time.Millisecond)
	}
}
