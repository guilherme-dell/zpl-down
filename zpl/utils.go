package zpl

import (
	"fmt"
	"os"
)

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
