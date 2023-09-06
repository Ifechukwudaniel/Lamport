package ast

import (
	"lamort-go/token"
)

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

type Identifier struct {
	Token token.Token
	Value string
}

type Program struct {
	Statement []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}
