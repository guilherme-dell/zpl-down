package zpl

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Barcode struct {
	url       string
	ZplConfig string
	Prefix    string
	Amount    int
}

func (b Barcode) Download() error {
	b.url = "http://api.labelary.com/v1/printers/8dpmm/labels/0.98x0.39/0/ --data-urlencode"
	for i := 1; i <= b.Amount; i++ {
		URL := fmt.Sprintf("%s\"^xa^fo4,15^be1^bcn,40,,,,a^fd%04d%s^fs^xz\"", b.url, i, b.Prefix)
		fmt.Println(URL)
		response, err := http.Get(URL)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return errors.New("STATUS CODE ERROR")
		}
		fileName := fmt.Sprintf("../bars/%04d%s.png", i, b.Prefix)
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}
	}

	return nil
}
