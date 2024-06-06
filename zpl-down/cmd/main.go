package main

import (
	"github.com/guilherme-dell/zpl-down/internal/config"
	"github.com/guilherme-dell/zpl-down/internal/downloader"
	"github.com/guilherme-dell/zpl-down/pkg/utils"
)

func main() {

	cfg := config.LoadConfig("./layouts/layout_1.json")

	utils.CreateDir(cfg)

	downloader.DownloadBarCodes(cfg.BarcodeConfig, cfg)
}
