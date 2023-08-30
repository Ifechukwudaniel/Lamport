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
	INT    // 12345
	FLOAT  // 123.45
	IMAG   // 123.45i
	CHAR   // 'a'
	STRING // "abc"
	literal_end

	operator_beg

	ADD
	SUB
	MUL
	DIV

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

	// Delimiters
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
	PACKAGE
	RETURN

	STRUCT
	SWITCH
	VAR

	PUBLIC
	PRIVATE

	IMPLEMETAION
	keyword_end
)

var tokens = [...]string{
	ILLEGER:  "ILLEGER",
	EOF:      "EOF",
	COMMENTS: "COMMENTS",

	INT:    "INT",
	FLOAT:  "FLOAT",
	IMAG:   "IMAG",
	CHAR:   "CHAR",
	STRING: "STRING",

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

	STRUCT:       "struct",
	IMPLEMETAION: "impl",

	SWITCH: "switch",
	VAR:    "var",

	PUBLIC:  "public",
	PRIVATE: "private",
}
