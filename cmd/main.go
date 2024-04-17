package main

import (
	"github.com/guilherme-dell/zpl-down/zpl"
)

func main() {
	var barcode zpl.Barcode

	barcode.Dpmm = "8dpmm"
	barcode.Width = "0.98"
	barcode.Height = "0.39"
	barcode.Prefix = "CJ"
	barcode.GenerateAmount = 50

	barcode.Download()
}
