package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA   = ","
	LPAREN  = "("
	RPAREN  = ")"
	NEWLINE = "\n"

	// Keywords
	DEF    = "DEF"
	DO     = "DO"
	END    = "END"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"do":  DO,
	"end": END,
	"def": DEF,
}

func LookupIdent(identifier string) TokenType {
	if tok, found := keywords[identifier]; found {
		return tok
	}
	return IDENT
}
