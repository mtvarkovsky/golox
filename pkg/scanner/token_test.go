package scanner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenTypeToString(t *testing.T) {
	cases := []struct {
		tType TokenType
		str   string
	}{
		{
			tType: LeftParen,
			str:   "LEFT_PAREN",
		},
		{
			tType: RightParen,
			str:   "RIGHT_PAREN",
		},
		{
			tType: LeftBrace,
			str:   "LEFT_BRACE",
		},
		{
			tType: RightBrace,
			str:   "RIGHT_BRACE",
		},
		{
			tType: Comma,
			str:   "COMMA",
		},
		{
			tType: Dot,
			str:   "DOT",
		},
		{
			tType: Minus,
			str:   "MINUS",
		},
		{
			tType: Plus,
			str:   "PLUS",
		},
		{
			tType: Semicolon,
			str:   "SEMICOLON",
		},
		{
			tType: Slash,
			str:   "SLASH",
		},
		{
			tType: Star,
			str:   "STAR",
		},
		{
			tType: Bang,
			str:   "BANG",
		},
		{
			tType: BangEqual,
			str:   "BANG_EQUAL",
		},
		{
			tType: Equal,
			str:   "EQUAL",
		},
		{
			tType: EqualEqual,
			str:   "EQUAL_EQUAL",
		},
		{
			tType: Greater,
			str:   "GREATER",
		},
		{
			tType: GreaterEqual,
			str:   "GREATER_EQUAL",
		},
		{
			tType: Less,
			str:   "LESS",
		},
		{
			tType: LessEqual,
			str:   "LESS_EQUAL",
		},
		{
			tType: SlashSlash,
			str:   "SLASH_SLASH",
		},
		{
			tType: Identifier,
			str:   "IDENTIFIER",
		},
		{
			tType: String,
			str:   "STRING",
		},
		{
			tType: Number,
			str:   "NUMBER",
		},
		{
			tType: And,
			str:   "AND",
		},
		{
			tType: Class,
			str:   "CLASS",
		},
		{
			tType: Else,
			str:   "ELSE",
		},
		{
			tType: False,
			str:   "FALSE",
		},
		{
			tType: Fun,
			str:   "FUN",
		},
		{
			tType: For,
			str:   "FOR",
		},
		{
			tType: If,
			str:   "IF",
		},
		{
			tType: Nil,
			str:   "NIL",
		},
		{
			tType: Or,
			str:   "OR",
		},
		{
			tType: Print,
			str:   "PRINT",
		},
		{
			tType: Return,
			str:   "RETURN",
		},
		{
			tType: Super,
			str:   "SUPER",
		},
		{
			tType: This,
			str:   "THIS",
		},
		{
			tType: True,
			str:   "TRUE",
		},
		{
			tType: Var,
			str:   "VAR",
		},
		{
			tType: While,
			str:   "WHILE",
		},
		{
			tType: EOF,
			str:   "EOF",
		},
	}

	for _, tc := range cases {
		t.Run(tc.str, func(t *testing.T) {
			result := tc.tType.String() == tc.str
			assert.True(t, result)
		})
	}
}

var testToken = NewToken(Number, "3.14", 3.14, 12, 13)

func TestToken_String(t *testing.T) {
	assert.Equal(t, `NUMBER 3.14 3.14`, testToken.String())
}

func TestToken_Lexeme(t *testing.T) {
	assert.Equal(t, "3.14", testToken.Lexeme())
}

func TestToken_Literal(t *testing.T) {
	assert.Equal(t, 3.14, testToken.Literal())
}

func TestToken_Type(t *testing.T) {
	assert.Equal(t, Number, testToken.Type())
}

func TestToken_Line(t *testing.T) {
	assert.Equal(t, 12, testToken.Line())
}

func TestToken_Position(t *testing.T) {
	assert.Equal(t, 12, testToken.Position())
}
