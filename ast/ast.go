package ast

import "github.com/jibrankalia/monkeylang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type AssignStatement struct {
	Token token.Token // token .ASSIGN
	Name  *Identifier
	Value Expression
}

// statetment newToken
func (ls *AssignStatement) statementNode() {}

func (ls *AssignStatement) TokenLiteral() string { return ls.Token.Literal }

// Expression interface
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
