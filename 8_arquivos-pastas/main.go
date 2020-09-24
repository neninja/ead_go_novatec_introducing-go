package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// showContentOfFile1()
	showContentOfFile2()
	// criaArquivo()
	showContentOfDir()
	showContentOfDirRecursivamente()
}

func showContentOfFile1() {
	file, err := os.Open("../README.md")
	if err != nil {
		return
	}
	defer file.Close()

	// status do arquivo
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// cria array de bytes do mesmo tamanho que o arquivo
	// e passa seu conteudo para a variavel posteriormente
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

// forma mais simples de leitura de arquivo
func showContentOfFile2() {
	bs, err := ioutil.ReadFile("../README.md")
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func criaArquivo() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}

	defer file.Close()
	file.WriteString("exemplo")
}

func showContentOfDir() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}

	defer dir.Close()

	// -1 retorna todo o conteudo do diretório,
	// se for passado n valor será retornado n arquivos/pastas
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func showContentOfDirRecursivamente() {
	// função anonima é um callback executado a cada arquivo e pasta
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
