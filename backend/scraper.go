package main

import (
	"code.sajari.com/docconv"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func readSubFiles(path string, stringFiles []string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	for _, file := range files {
		fmt.Println(file.Name(), " ", file.IsDir())
		if file.IsDir() {
			stringFiles = readSubFiles(AddSubFileToPath(path, file.Name()), stringFiles)
		} else {
			if err != nil {
				log.Fatal(err)
			}
			if strings.HasSuffix(file.Name(), "pdf") {
				c, err := ReadPDF(AddSubFileToPath(path, file.Name()))
				if err != nil {
					log.Fatal(err)
				}
				stringFiles = append(stringFiles, c)
			}
		}
		i++
	}
	return stringFiles
}

func ReadSubfiles(path string) []string {
	stringFiles := make([]string, 50)
	return readSubFiles(path, stringFiles)
}

func ReadPDF(file string) (string, error) {
	content, err := docconv.ConvertPath(file)
	if err != nil {
		return "", err
	}
	return content.Body, nil
}

func AddSubFileToPath(path string, file string) string {
	return path + "/" + file
}
