package lexer

import (
	"errors"

	"github.com/jibrankalia/monkeylang/token"
)

type Lexer struct {
	input        string
	position     int  // position of current char
	readPosition int  // position after current char
	ch           byte // the actual current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) backTrackIndex() error {
	if l.position < 1 {
		return errors.New("Cannot backtrack")
	}
	l.readPosition = l.position
	l.position -= 1
	return nil
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			literal := l.readTwoChar()
			tok = newTokenStr(token.EQ, literal)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '\n':
		tok = newToken(token.NEWLINE, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			literal := l.readTwoChar()
			tok = newTokenStr(token.NOT_EQ, literal)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func newTokenStr(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

func (l *Lexer) readTwoChar() string {
	ch := l.ch
	l.readChar()
	return string(ch) + string(l.ch)
}

// Get the whole word
func (l *Lexer) readIdentifier() string {
	initial_position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	identifier := l.input[initial_position:l.position]
	l.backTrackIndex()
	return identifier
}

func (l *Lexer) readNumber() string {
	initial_position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	number := l.input[initial_position:l.position]
	l.backTrackIndex()
	return number
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
