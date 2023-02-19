package scanner

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
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
	}

	Error struct {
		Line int
		Pos  int
		Err  error
	}
)

var (
	Keywords = map[string]TokenType{
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
	}

	SingleCharLexemeToToken = map[rune]TokenType{
		'(': LeftParen,
		')': RightParen,
		'{': LeftBrace,
		'}': RightBrace,
		',': Comma,
		'.': Dot,
		'-': Minus,
		'+': Plus,
		';': Semicolon,
		'*': Star,
		'!': Bang,
		'=': Equal,
		'<': Less,
		'>': Greater,
		'/': Slash,
	}

	TwoCharLexemeToToken = map[string]TokenType{
		"!=": BangEqual,
		"==": EqualEqual,
		"<=": LessEqual,
		">=": GreaterEqual,
	}

	StringCharLexemeToToken = map[rune]TokenType{
		'"': String,
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

	s.appendToken(NewToken(EOF, "", nil, s.currentLine))

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
				for s.peek() != '\n' && !s.isAtEnd {
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
		Line: s.currentLine,
	}
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
			Pos:  s.currentPos,
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
			Pos:  s.currentPos,
			Err:  fmt.Errorf("unsopported number format"),
		}
	}
	s.addToken(Number, val)

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
