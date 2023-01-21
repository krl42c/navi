package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	run_script := flag.String("s", "nil", "script path")
	flag.Parse()

	db := database{name: "local"}
	if *run_script != "nil" {
		read_file, err := os.Open(*run_script)
		if err != nil {
			fmt.Println("Error reading file: ", err)
		}

		file_scanner := bufio.NewScanner(read_file)
		file_scanner.Split(bufio.ScanLines)

		line_number := 0
		for file_scanner.Scan() {
			toks, err := execute_line(&db, file_scanner.Text(), line_number)
			if err != nil {
				fmt.Printf(err.Error(), toks)
			}
			line_number++
		}

		read_file.Close()
	}

	for _, tbl := range db.tables {
		fmt.Println(tbl)
	}
}
