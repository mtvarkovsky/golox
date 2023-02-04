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
			tType: Eof,
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
