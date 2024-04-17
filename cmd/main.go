package main

import (
	"github.com/guilherme-dell/zpl-down/zpl"
)

func main() {
	var barcode zpl.Barcode

	barcode.ZplConfig = "\"^xa^fo5,15^be1^bcn,40,,,,a^fd0001CJ^fs^xz\""
	barcode.Prefix = "CJ"
	barcode.Amount = 1

	barcode.Download()
}
