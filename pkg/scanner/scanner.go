package scanner

import (
	"fmt"
	"strconv"
)

type (
	Scanner interface {
		ScanTokens() ([]Token, []*Error)
		IsAtEnd() bool
	}

	scanner struct {
		input   string
		tokens  []Token
		isAtEnd bool

		lexemeStartPos int
		currentPos     int
		currentLine    int

		keywords map[string]TokenType
	}

	Error struct {
		Line int
		Err  error
	}
)

// TODO: Scanner implementation is straight from the book and needs improvements
func NewScanner(input string) Scanner {
	return &scanner{
		input:          input,
		lexemeStartPos: 0,
		currentPos:     0,
		currentLine:    1,

		keywords: map[string]TokenType{
			"and":    And,
			"class":  Class,
			"else":   Else,
			"false":  False,
			"for":    For,
			"fun":    Fun,
			"if":     If,
			"nil":    Nil,
			"or":     Or,
			"print":  Print,
			"return": Return,
			"super":  Super,
			"this":   This,
			"true":   True,
			"var":    Var,
			"while":  While,
		},
	}
}

func (s *scanner) ScanTokens() ([]Token, []*Error) {
	var errs []*Error

	for !s.IsAtEnd() {
		s.lexemeStartPos = s.currentPos
		if err := s.scanToken(); err != nil {
			errs = append(errs, err)
		}
	}

	s.appendToken(NewToken(Eof, "", nil, s.currentLine))

	return s.tokens, errs
}

func (s *scanner) appendToken(t Token) {
	s.tokens = append(s.tokens, t)
}

func (s *scanner) addToken(tType TokenType, literal any) {
	text := s.input[s.lexemeStartPos:s.currentPos]
	s.appendToken(NewToken(tType, text, literal, s.currentLine))
}

func (s *scanner) scanToken() *Error {
	c := s.next()

	switch c {

	// Single char tokens
	case '(':
		s.addToken(LeftParen, nil)
	case ')':
		s.addToken(RightParen, nil)
	case '{':
		s.addToken(LeftBrace, nil)
	case '}':
		s.addToken(RightBrace, nil)
	case ',':
		s.addToken(Comma, nil)
	case '.':
		s.addToken(Dot, nil)
	case '-':
		s.addToken(Minus, nil)
	case '+':
		s.addToken(Plus, nil)
	case ';':
		s.addToken(Semicolon, nil)
	case '*':
		s.addToken(Star, nil)

	// Single or two character tokens
	case '!':
		if s.match('=') {
			s.addToken(BangEqual, nil)
		} else {
			s.addToken(Bang, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(EqualEqual, nil)
		} else {
			s.addToken(Equal, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(LessEqual, nil)
		} else {
			s.addToken(Less, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(GreaterEqual, nil)
		} else {
			s.addToken(Greater, nil)
		}

	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd {
				_ = s.next()
			}
		} else {
			s.addToken(Slash, nil)
		}

	// Ignored whitespace chars
	case ' ':
	case '\r':
	case '\t':

	// Advance line counter
	case '\n':
		s.currentLine++

	// Strings
	case '"':
		return s.string()

	// Unrecognized token
	default:
		if s.isDigit(c) {
			return s.number()
		}

		if s.isAlpha(c) {
			s.identifier()
			return nil
		}

		return &Error{
			Err:  fmt.Errorf("unexpected character %c", c),
			Line: s.currentLine,
		}
	}

	return nil
}

func (s *scanner) match(expected rune) bool {
	if s.isAtEnd {
		return false
	}

	if rune(s.input[s.currentPos]) != expected {
		return false
	}

	s.currentPos++
	return true
}

func (s *scanner) next() rune {
	r := rune(s.input[s.currentPos])
	s.currentPos++
	return r
}

func (s *scanner) peek() rune {
	if s.isAtEnd {
		return 0
	}

	return rune(s.input[s.currentPos])
}

func (s *scanner) peekNext() rune {
	if s.currentPos+1 >= len(s.input) {
		return 0
	}

	return rune(s.input[s.currentPos+1])
}

func (s *scanner) string() *Error {
	for s.peek() != '"' && !s.isAtEnd {
		if s.peek() == '\n' {
			s.currentLine++
		}

		_ = s.next()
	}

	if s.isAtEnd {
		return &Error{
			Line: s.currentLine,
			Err:  fmt.Errorf("unterminated string"),
		}
	}

	_ = s.next()

	val := s.input[s.lexemeStartPos+1 : s.currentPos-1]
	s.addToken(String, val)

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
			Err:  fmt.Errorf("unsopported number format"),
		}
	}
	s.addToken(Number, val)

	return nil
}

func (s *scanner) isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (s *scanner) isAlphaNumeric(c rune) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		_ = s.next()
	}

	val := s.input[s.lexemeStartPos:s.currentPos]

	tokenType, found := s.keywords[val]
	if !found {
		tokenType = Identifier
	}

	s.addToken(tokenType, nil)
}

func (s *scanner) IsAtEnd() bool {
	return s.currentPos >= len(s.input)
}

func (e *Error) Error() string {
	return e.Err.Error()
}
