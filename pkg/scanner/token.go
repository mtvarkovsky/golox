package scanner

import "fmt"

type TokenType int

const (
	// Single character tokens
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	// One or two character tokens
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual
	SlashSlash

	// Literals
	Identifier
	String
	Number

	// Keywords
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	EOF
)

func (tt TokenType) String() string {
	return [...]string{
		// Single character tokens
		"LEFT_PAREN",
		"RIGHT_PAREN",
		"LEFT_BRACE",
		"RIGHT_BRACE",
		"COMMA",
		"DOT",
		"MINUS",
		"PLUS",
		"SEMICOLON",
		"SLASH",
		"STAR",

		// One or two character tokens
		"BANG",
		"BANG_EQUAL",
		"EQUAL",
		"EQUAL_EQUAL",
		"GREATER",
		"GREATER_EQUAL",
		"LESS",
		"LESS_EQUAL",
		"SLASH_SLASH",

		// Literals
		"IDENTIFIER",
		"STRING",
		"NUMBER",

		// Keywords
		"AND",
		"CLASS",
		"ELSE",
		"FALSE",
		"FUN",
		"FOR",
		"IF",
		"NIL",
		"OR",
		"PRINT",
		"RETURN",
		"SUPER",
		"THIS",
		"TRUE",
		"VAR",
		"WHILE",

		"EOF",
	}[tt]
}

type (
	Token interface {
		String() string
		Lexeme() string
		Literal() any
		Line() int
		Type() TokenType
	}

	token struct {
		lexeme    string
		tokenType TokenType
		line      int
		literal   any
	}
)

func NewToken(tType TokenType, lexeme string, literal any, line int) Token {
	return &token{
		tokenType: tType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

func (t *token) String() string {
	return fmt.Sprintf("%s %s %v", t.tokenType, t.lexeme, t.literal)
}

func (t *token) Lexeme() string {
	return t.lexeme
}

func (t *token) Type() TokenType {
	return t.tokenType
}

func (t *token) Literal() any {
	return t.literal
}

func (t *token) Line() int {
	return t.line
}
