package scanner

import (
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/tokens"
	"strconv"
	"strings"
	"unicode"
)

type (
	Scanner interface {
		ScanTokens() ([]tokens.Token, []*Error)
		IsAtEnd() bool
	}

	scanner struct {
		input  string
		tokens []tokens.Token

		lexemeStartPos int
		currentPos     int
		currentLine    int
		currentLinePos int
	}

	Error struct {
		Line int
		Pos  int
		Err  error
	}
)

var (
	Keywords = map[string]tokens.TokenType{
		"and":    tokens.And,
		"class":  tokens.Class,
		"else":   tokens.Else,
		"false":  tokens.False,
		"for":    tokens.For,
		"fun":    tokens.Fun,
		"if":     tokens.If,
		"nil":    tokens.Nil,
		"or":     tokens.Or,
		"print":  tokens.Print,
		"return": tokens.Return,
		"super":  tokens.Super,
		"this":   tokens.This,
		"true":   tokens.True,
		"var":    tokens.Var,
		"while":  tokens.While,
	}

	SingleCharLexemeToToken = map[rune]tokens.TokenType{
		'(': tokens.LeftParen,
		')': tokens.RightParen,
		'{': tokens.LeftBrace,
		'}': tokens.RightBrace,
		',': tokens.Comma,
		'.': tokens.Dot,
		'-': tokens.Minus,
		'+': tokens.Plus,
		';': tokens.Semicolon,
		'*': tokens.Star,
		'!': tokens.Bang,
		'=': tokens.Equal,
		'<': tokens.Less,
		'>': tokens.Greater,
		'/': tokens.Slash,
	}

	TwoCharLexemeToToken = map[string]tokens.TokenType{
		"!=": tokens.BangEqual,
		"==": tokens.EqualEqual,
		"<=": tokens.LessEqual,
		">=": tokens.GreaterEqual,
	}

	StringCharLexemeToToken = map[rune]tokens.TokenType{
		'"': tokens.String,
	}

	IgnoredCharsSet = map[rune]bool{
		' ':  true,
		'\r': true,
		'\t': true,
	}

	NewLineCharsSet = map[rune]bool{
		'\n': true,
	}

	CommentCharsSet = map[string]bool{
		"//": true,
	}
)

// TODO: Scanner implementation is straight from the book and needs improvements
func NewScanner(input string) Scanner {
	return &scanner{
		input:          input,
		lexemeStartPos: 0,
		currentPos:     0,
		currentLine:    1,
		currentLinePos: 0,
	}
}

func (s *scanner) ScanTokens() ([]tokens.Token, []*Error) {
	var errs []*Error

	for !s.IsAtEnd() {
		s.lexemeStartPos = s.currentPos
		if err := s.scanToken(); err != nil {
			errs = append(errs, err)
		}
	}

	s.appendToken(tokens.NewToken(tokens.EOF, "", nil, s.currentLine, s.currentLinePos))

	return s.tokens, errs
}

func (s *scanner) appendToken(t tokens.Token) {
	s.tokens = append(s.tokens, t)
}

func (s *scanner) addToken(tType tokens.TokenType, literal any) {
	text := s.input[s.lexemeStartPos:s.currentPos]
	s.appendToken(tokens.NewToken(tType, text, literal, s.currentLine, s.currentLinePos-len(text)+1))
}

func (s *scanner) scanToken() *Error {
	c := s.next()

	if singleCharToken, foundSingleChar := SingleCharLexemeToToken[c]; foundSingleChar {
		nextC := s.peek()

		if nextC != 0 {
			twoCharBuffer := strings.Builder{}
			twoCharBuffer.WriteRune(c)
			twoCharBuffer.WriteRune(nextC)

			if twoCharToken, foundTwoChar := TwoCharLexemeToToken[twoCharBuffer.String()]; foundTwoChar {
				_ = s.next()
				s.addToken(twoCharToken, nil)
				return nil
			}

			if _, foundComment := CommentCharsSet[twoCharBuffer.String()]; foundComment {
				for s.peek() != '\n' && !s.IsAtEnd() {
					_ = s.next()
				}
				return nil
			}
		}

		s.addToken(singleCharToken, nil)
		return nil
	}

	if _, foundIgnored := IgnoredCharsSet[c]; foundIgnored {
		return nil
	}

	if _, foundNewLine := NewLineCharsSet[c]; foundNewLine {
		s.currentLine++
		s.currentLinePos = 0
		return nil
	}

	if _, foundString := StringCharLexemeToToken[c]; foundString {
		return s.string()
	}

	if s.isDigit(c) {
		return s.number()
	}

	if s.isAlpha(c) {
		s.identifier()
		return nil
	}

	return &Error{
		Err:  fmt.Errorf("unexpected character %c", c),
		Pos:  s.currentPos,
		Line: s.currentLinePos,
	}
}

func (s *scanner) next() rune {
	if s.IsAtEnd() {
		return 0
	}
	r := rune(s.input[s.currentPos])
	s.currentPos++
	s.currentLinePos++
	return r
}

func (s *scanner) peek() rune {
	if s.IsAtEnd() {
		return 0
	}

	return rune(s.input[s.currentPos])
}

func (s *scanner) peekNext() rune {
	if s.currentPos+1 >= len(s.input)-1 {
		return 0
	}

	return rune(s.input[s.currentPos+1])
}

func (s *scanner) string() *Error {
	for s.peek() != '"' && !s.IsAtEnd() {
		if s.peek() == '\n' {
			s.currentLine++
			s.currentLinePos++
			s.currentLinePos = 0
		}

		_ = s.next()
	}

	if s.IsAtEnd() {
		return &Error{
			Line: s.currentLine,
			Pos:  s.currentLinePos,
			Err:  fmt.Errorf("unterminated string"),
		}
	}

	_ = s.next()

	val := s.input[s.lexemeStartPos+1 : s.currentPos-1]
	s.addToken(tokens.String, val)

	return nil
}

func (s *scanner) isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (s *scanner) number() *Error {
	for s.isDigit(s.peek()) {
		_ = s.next()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		_ = s.next()

		for s.isDigit(s.peek()) {
			_ = s.next()
		}
	}

	val, err := strconv.ParseFloat(s.input[s.lexemeStartPos:s.currentPos], 54)
	if err != nil {
		return &Error{
			Line: s.currentLine,
			Pos:  s.currentLinePos,
			Err:  fmt.Errorf("unsopported number format"),
		}
	}
	s.addToken(tokens.Number, val)

	return nil
}

func (s *scanner) isAlpha(c rune) bool {
	return unicode.IsLetter(c) || c == '_'
}

func (s *scanner) isAlphaNumeric(c rune) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		_ = s.next()
	}

	val := s.input[s.lexemeStartPos:s.currentPos]

	tokenType, found := Keywords[val]
	if !found {
		tokenType = tokens.Identifier
	}

	s.addToken(tokenType, nil)
}

func (s *scanner) IsAtEnd() bool {
	return s.currentPos >= len(s.input)-1
}

func (e *Error) Error() string {
	return e.Err.Error()
}
