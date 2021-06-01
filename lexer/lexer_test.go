package lexer

import (
	"testing"

	"github.com/jibrankalia/monkeylang/token"
)

func TestNextToken(t *testing.T) {
	input := `five = 5
ten = 10

def add(x, y) do
	x + y
end

result = add(five, ten)

!-/*5
5 < 10 > 5

==
!=

10 == 10
10 != 9
`
	// if (5 < 10) {
	// 	return true;
	// } else {
	// 	return false;
	// }

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.DEF, "def"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.DO, "do"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.NEWLINE, "\n"},
		{token.END, "end"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.NEWLINE, "\n"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.EQ, "=="},
		{token.NEWLINE, "\n"},
		{token.NOT_EQ, "!="},
		// {token.IF, "if"},
		// {token.LPAREN, "("},
		// {token.INT, "5"},
		// {token.LT, "<"},
		// {token.INT, "10"},
		// {token.RPAREN, ")"},
		// {token.LBRACE, "{"},
		// {token.RETURN, "return"},
		// {token.TRUE, "true"},
		// {token.SEMICOLON, ";"},
		// {token.RBRACE, "}"},
		// {token.ELSE, "else"},
		// {token.LBRACE, "{"},
		// {token.RETURN, "return"},
		// {token.FALSE, "false"},
		// {token.SEMICOLON, ";"},
		// {token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.NEWLINE, "\n"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.NEWLINE, "\n"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, test := range tests {
		token := l.NextToken()
		if token.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, token.Type)

		}
		if token.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedLiteral, token.Literal)
		}
	}
}
