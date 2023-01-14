package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type file_ident struct {
	filename  string
	full_path string
	uptodate  bool
}

type file_record struct {
	file_list []file_ident
}

func clear(record *file_record) {
}

func main() {
	record := file_record{file_list: nil}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		full_path := wd + file.Name()
		record.file_list = append(record.file_list, file_ident{filename: file.Name(), full_path: full_path, uptodate: true})
	}

	for _, f := range record.file_list {
		fmt.Println(f)
	}
}
