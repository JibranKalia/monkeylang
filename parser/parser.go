package parser

import (
	"github.com/jibrankalia/monkeylang/ast"
	"github.com/jibrankalia/monkeylang/lexer"
	"github.com/jibrankalia/monkeylang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {
	if !p.peekTokenIs(token.ASSIGN) {
		// todo raise error
		return nil
	}

	stmt := &ast.AssignStatement{Token: p.peekToken}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// TODO: We're skipping the expressions until we
	// encounter a newline

	for !p.curTokenIs(token.NEWLINE) {
		p.nextToken()
	}

	return stmt

}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.IDENT:
		if p.peekToken.Type == token.ASSIGN {
			return p.parseAssignStatement()

		} else {
			// TODO: handle else case
			return nil
		}

	default:
		return nil
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// populate currToken and peekToken
	p.nextToken()
	p.nextToken()

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
