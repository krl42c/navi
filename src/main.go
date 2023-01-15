package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	/*
		db := database{name: "localdb"}
		group_tbl := create_table("groups")

		insert := db.insert_table(group_tbl)

		if insert != nil {
			fmt.Println("Error appending table %s", group_tbl.name)
		}

		name_col := column[string]{name: "name"}
		password_col := column[string]{name: "category"}

		db.tables[0].insert_col_str(name_col)
		db.tables[0].insert_col_str(password_col)

		r1 := row[string]{index: 1, value: "admin"}
		r2 := row[string]{index: 2, value: "guest"}

		db.tables[0].cols_str[0].rows = append(db.tables[0].cols_str[0].rows, r1)
		db.tables[0].cols_str[0].rows = append(db.tables[0].cols_str[0].rows, r2)

		lines := lex_line("create table groups (name value)")
		tokens := construct_tokens(lines)

		for _, tok := range tokens {
			fmt.Println(tok)
		}*/

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

		for file_scanner.Scan() {
			execute_line(&db, file_scanner.Text())
		}

		read_file.Close()
	}

	for _, tbl := range db.tables {
		fmt.Println(tbl)
	}
}
