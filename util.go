package aoc

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func GetFileContents() string {

	path, _ := filepath.Abs("input.txt")

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	values, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	return string(values)
}
