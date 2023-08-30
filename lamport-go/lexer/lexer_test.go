package lexer

import (
	"fmt"
	"lamort-go/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ADD, "+"},
		{token.SUB, "-"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, testToken := range tests {
		tok := l.NextToken()
		fmt.Printf("%d %d %d \n", i, tok.Type, testToken.expectedType)
		fmt.Printf("%d %s %s \n", i, tok.Literal, testToken.expectedLiteral)
		// if tok.Type != tokenTest.expectedType {
		// 	t.Fatalf("tests[%d]- tokentype wrong. expected=%q got=%q", i, tokenTest.expectedType, tok.Type)
		// }
		//
		// if tok.Literal != tokenTest.expectedLiteral {
		// 	t.Fatalf("tests[%d]- Literal Type wrong. expected=%s got=%s", i, tokenTest.expectedLiteral, tok.Literal)
		// }
	}
}
