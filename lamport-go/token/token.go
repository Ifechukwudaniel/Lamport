package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (

	// SPACIAL TOKENS
	ILLEGER = iota
	EOF
	COMMENTS

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT
	INT    // 12345
	FLOAT  // 123.45
	IMAG   // 123.45i
	CHAR   // 'a'
	STRING // "abc"
	literal_end

	operator_beg

	ADD // +
	SUB // -
	MUL // *
	DIV // /
	REM // %

	AND // &
	OR  // |
	XOR // ^
	SHL // <<
	SHR // >>

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=

	LAND       // &&
	ARROW_FUNC // =>
	INC        // ++
	DEC        // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =

	NOT // !
	LOR // ||
	NEQ // !=
	LEQ // <=
	GEQ // >=

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :

	operator_end

	keyword_beg
	// Keywords
	BREAK
	CASE
	CONST
	CONTINUE

	ELSE
	FOR

	FUNC
	GO
	IF
	IMPORT

	INTERFACE
	EXPORT
	RETURN

	STRUCT
	SWITCH
	VAR
	LET

	PUBLIC
	PRIVATE

	IMPLEMENTATION
	keyword_end
)

var tokens = [...]string{
	ILLEGER:  "ILLEGER",
	EOF:      "EOF",
	COMMENTS: "COMMENTS",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	IMAG:   "IMAG",
	CHAR:   "CHAR",
	STRING: "STRING",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",
	REM: "%",

	AND: "&",
	OR:  "|",
	XOR: "^",
	SHL: "<<",
	SHR: ">>",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",

	LAND:       "&&",
	LOR:        "||",
	ARROW_FUNC: "=>",
	INC:        "++",
	DEC:        "--",

	NEQ: "!=",
	LEQ: "<=",
	GEQ: ">=",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",

	BREAK:    "break",
	CASE:     "case",
	CONST:    "const",
	CONTINUE: "continue",

	ELSE: "else",
	FOR:  "for",

	FUNC:   "fn",
	GO:     "go",
	IF:     "if",
	IMPORT: "import",

	INTERFACE: "interface",
	RETURN:    "return",

	STRUCT:         "struct",
	IMPLEMENTATION: "impl",
	EXPORT:         "export",

	SWITCH: "switch",
	VAR:    "var",
	LET:    "let",

	PUBLIC:  "public",
	PRIVATE: "private",
}

var keywords map[string]TokenType

func init() {
	keywords = make(map[string]TokenType, keyword_end-(keyword_beg+1))
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = TokenType(i)
	}

}

// Return for true if tokens is a keyword
func isKeyword(name string) bool {
	_, ok := keywords[name]
	return ok
}

// Return for tokens for corresponding literal
func (tok TokenType) isLiteral() bool { return literal_beg < tok && tok < literal_end }

// Return  true if for  operator tokens
func (tok TokenType) isOperator() bool { return operator_beg < tok && tok < operator_end }

// Lookup tokens by there strings
func Lookup(ident string) TokenType {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return IDENT
}

// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
func (op TokenType) Precedence() int {
	switch op {
	case LOR:
		return 1
	case LAND:
		return 2
	case EQL, NEQ, LSS, LEQ, GTR, GEQ:
		return 3
	case ADD, SUB, OR, XOR:
		return 4
	case MUL, DIV, REM, SHL, SHR, AND:
		return 5
	}
	return LowestPrec
}
