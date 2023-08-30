package lexer

import "lamort-go/token"

type Lexer struct {
	input        string
	postion      int
	readPosition int
	ch           string
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {

	if l.readPosition >= len(l.input) {
		l.ch = ""
	} else {
		l.ch = string(l.input[l.readPosition])
	}

	l.postion = l.readPosition
	l.readPosition += 1

}

func newToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	switch l.ch {
	case "=":
		tok = newToken(token.ASSIGN, l.ch)
	case ";":
		tok = newToken(token.SEMICOLON, l.ch)
	case "(":
		tok = newToken(token.LBRACK, l.ch)
	case ")":
		tok = newToken(token.RBRACK, l.ch)
	case ",":
		tok = newToken(token.COMMA, l.ch)
	case "+":
		tok = newToken(token.ADD, l.ch)
	case "{":
		tok = newToken(token.RBRACE, l.ch)
	case "}":
		tok = newToken(token.LBRACE, l.ch)
	case "":
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}
