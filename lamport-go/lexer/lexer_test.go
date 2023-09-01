package lexer

import (
	"fmt"
	"lamort-go/token"
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `
		let five = 5;
		let ten = 10;
		fn add(x, y) {
			x + y;
		};
		let result = add(five, ten);
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.FUNC, "fn"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.ADD, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, testToken := range tests {
		tok := l.NextToken()
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
