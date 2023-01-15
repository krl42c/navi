package main

type token_type int16

const (
	CREATE token_type = iota
	TABLE
	IDENTIFIER
	DROP
	INSERT
	TYPE
	OPEN_PAREN
	CLOSED_PAREN
	ENDLINE
)

type token struct {
	ttype  token_type
	tvalue string
}

func lex_line(line string) []string {
	var single string
	var toks []string

	for _, r := range line {
		if string(r) == "(" {
			toks = append(toks, string(r))
			continue
		}

		if string(r) == ")" {
			toks = append(toks, single)
			toks = append(toks, string(r))
			single = ""
			continue
		}

		if string(r) == " " {
			if single != "" {
				toks = append(toks, single)
			}
			single = ""
			continue
		}
		single += string(r)
	}
	toks = append(toks, single)
	return toks
}

func construct_tokens(tokens []string) []token {
	ret_toks := []token{}
	for _, t := range tokens {
		var tok token
		switch t {
		case "create":
			tok = token{ttype: CREATE}
		case "table":
			tok = token{ttype: TABLE}
		case "insert":
			tok = token{ttype: INSERT}
		case "(":
			tok = token{ttype: OPEN_PAREN}
		case ")":
			tok = token{ttype: CLOSED_PAREN}
		default:
			tok = token{ttype: IDENTIFIER, tvalue: t}
		}
		if t != "" {
			ret_toks = append(ret_toks, tok)
		}
	}

	ret_toks = append(ret_toks, token{ttype: ENDLINE})
	return ret_toks
}
