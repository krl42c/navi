package main

import "fmt"

type syntax_error struct {
	line     int
	actual   string
	excepted string
}

func (r *syntax_error) Error() string {
	return fmt.Sprintf("Syntax error on line %d: Excepted %s , Found %s ", r.line, r.excepted, r.actual)
}

func syntax_err(line int, actual string, excepted string) error {
	return &syntax_error{
		line:     line,
		actual:   actual,
		excepted: excepted,
	}
}

func execute_line(db *database, line string, line_number int) (err error) {
	tokens_str := lex_line(line)
	tokens := construct_tokens(tokens_str)
	fmt.Println(tokens)
	switch tokens[0].ttype {
	case CREATE:
		return create(db, tokens, line_number)
	case INSERT:
		insert(db, tokens)
	case DROP:
		drop(db, tokens)
	}
	return fmt.Errorf("Couldnt execute line")
}

func create(db *database, toks []token, line_number int) (err error) {
	tbl_name := toks[1].tvalue
	tbl := create_table(tbl_name)
	db.insert_table(tbl)

	if toks[2].ttype != OPEN_PAREN {
		return syntax_err(line_number, toks[2].tvalue, "(")
	}

	if toks[3].ttype != IDENTIFIER {
		return syntax_err(line_number, toks[2].tvalue, "column name")
	}

	if toks[2].ttype == OPEN_PAREN {
		params := toks[3:]
		if params == nil {
			return syntax_err(line_number, "nil", "parameter_list")
		}
		for _, p := range params {
			if p.ttype != CLOSED_PAREN && p.ttype != ENDLINE {
				col := column[string]{name: p.tvalue}
				db.tables[len(db.tables)-1].cols_str = append(db.tables[len(db.tables)-1].cols_str, col)
			}
		}
	}
	return nil
}

func insert(db *database, toks []token) {

}

func drop(db *database, toks []token) {

}
