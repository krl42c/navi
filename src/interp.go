package main

import "fmt"

type syntax_error struct {
	line     int
	actual   string
	excepted string
}

type statement_type int16

const (
	SCREATE statement_type = iota
	SINSERT
)

type statement struct {
	stype  statement_type
	tokens []token
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

func execute_line(db *database, line string, line_number int) (toks []token, err error) {
	tokens_str := lex_line(line)
	tokens := construct_tokens(tokens_str)
	switch tokens[0].ttype {
	case CREATE:
		st := statement{stype: SCREATE, tokens: tokens}
		nvv_create_table(db, st.tokens[1].tvalue, 0)
		return tokens, nil
	case INSERT:
		//data, err := insert(db, tokens, line_number)
		st := statement{stype: SINSERT, tokens: tokens}
		nvv_insert(db, st)
		return tokens, nil
	case DROP:
		drop(db, tokens)
	}
	return nil, fmt.Errorf("Couldnt execute line")
}

func create(db *database, toks []token, line_number int) (tokens []token, err error) {
	tbl_name := toks[1].tvalue
	tbl := create_table(tbl_name)
	db.insert_table(tbl)

	if toks[2].ttype != OPEN_PAREN {
		return nil, syntax_err(line_number, toks[2].tvalue, "(")
	}

	if toks[3].ttype != IDENTIFIER {
		return nil, syntax_err(line_number, toks[2].tvalue, "column name")
	}

	if toks[2].ttype == OPEN_PAREN {
		params := toks[3:]
		if params == nil {
			return nil, syntax_err(line_number, "nil", "parameter_list")
		}
		for _, p := range params {
			if p.ttype != CLOSED_PAREN && p.ttype != ENDLINE {
				col := column[string]{name: p.tvalue}
				db.tables[len(db.tables)-1].cols_str = append(db.tables[len(db.tables)-1].cols_str, col) // Mock creation
			}
		}
	}
	return toks, nil
}

func insert(db *database, toks []token, line_number int) (tokens []token, err error) {
	tbl, err := get_table(db, toks[1].tvalue)
	if err != nil {
		return nil, err
	}
	insert_values := []string{}

	if toks[2].ttype == OPEN_PAREN {
		params := toks[3:]
		for _, p := range params {
			if p.ttype != CLOSED_PAREN && p.ttype != ENDLINE {
				insert_values = append(insert_values, p.tvalue)
			}
		}
	}

	for i, ins := range insert_values {
		tbl.cols_str[i].rows = append(tbl.cols_str[i].rows, row[string]{value: ins})
	}
	return tokens, nil
}

func drop(db *database, toks []token) {

}
