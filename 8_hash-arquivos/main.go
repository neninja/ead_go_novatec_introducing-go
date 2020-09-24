package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}

	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}

	// se os hashes forem iguais a chance dos arquivos serem iguais é alta
	fmt.Println(h1, h2, h1 == h2)
}

func getHash(filename string) (uint32, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer f.Close()

	// hasher
	h := crc32.NewIEEE()

	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}

	// retorna codificado
	return h.Sum32(), nil
}
