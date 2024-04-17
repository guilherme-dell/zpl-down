package zpl

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Barcode struct {
	Dpmm           string
	Width          string
	Height         string
	Prefix         string
	GenerateAmount int
}

func (b Barcode) urlMount(indice int) string {
	return fmt.Sprintf("http://api.labelary.com/v1/printers/%s/labels/%sx%s/0/ --data-urlencode \"^xa^fo4,15^be1^bcn,40,,,,a^fd%04d%s^fs^xz\"",
		b.Dpmm, b.Width, b.Height, indice, b.Prefix)
}

func (b Barcode) Download() error {
	createDir()
	for i := 1; i <= b.GenerateAmount; i++ {
		URL := b.urlMount(i)

		response, err := http.Get(URL)
		if err != nil {
			return fmt.Errorf("Falha ao obter o código de barras da URL.\n%v", err)
		}

		defer response.Body.Close()

		if response.StatusCode != 200 {
			return errors.New("StatusCode ERROR")
		}

		fileName := fmt.Sprintf("../codigos_gerados/%04d%s.png", i, b.Prefix)
		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("Falha ao criar o arquivo.\n%v", err)
		}

		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			return fmt.Errorf("Falha ao salvar arquivo.\n%v", err)
		}

		fmt.Printf("BARCODE_%04d%s\t\t\033[0;32m[OK]\033[0m\n", i, b.Prefix)
	}

	return nil
}

func createDir() {
	dir := "../codigos_gerados"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("Erro ao criar diretorio.")
			panic(err)
		}
	}
}
