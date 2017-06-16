package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLineByLine(target string, requiredNumHits int64, pathToFile string) {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numHits int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, target) {
			numHits++
		}
	}

	if numHits >= requiredNumHits {
		fmt.Printf("%v\n", pathToFile)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	root := flag.String("root", ".", "where to search from")
	fileType := flag.String("type", "xml", "file type to search for")
	searchFor := flag.String("word", "foo", "word to search for")
	requiredNumHits := flag.Int64("hits", 1, "required number of hits")

	flag.Parse()

	err := filepath.Walk(*root, func(path string, f os.FileInfo, err error) error {
		if f != nil && f.Mode().IsRegular() && strings.HasSuffix(f.Name(), *fileType) {
			readLineByLine(*searchFor, *requiredNumHits, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
