package main

func execute_line(db *database, line string) {
	tokens_str := lex_line(line)
	tokens := construct_tokens(tokens_str)

	switch tokens[0].ttype {
	case CREATE:
		create(db, tokens)
	case INSERT:
		insert(db, tokens)
	case DROP:
		drop(db, tokens)
	}
}

func create(db *database, toks []token) {
	tbl_name := toks[1].tvalue
	tbl := create_table(tbl_name)
	db.insert_table(tbl)

	if toks[2].ttype == OPEN_PAREN {
		params := toks[3:]
		for _, p := range params {
			if p.ttype != CLOSED_PAREN && p.ttype != ENDLINE {
				col := column[string]{name: p.tvalue}
				db.tables[len(db.tables)-1].cols_str = append(db.tables[len(db.tables)-1].cols_str, col)
			}
		}
	}
}

func insert(db *database, toks []token) {

}

func drop(db *database, toks []token) {

}
