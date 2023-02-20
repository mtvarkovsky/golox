package parser

import (
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/ast"
	"github.com/mtvarkovsky/golox/pkg/scanner"
)

type (
	Parser interface {
		Parse() (ast.Expression, *Error)
	}

	parser struct {
		input      []scanner.Token
		currentPos int
	}

	Error struct {
		Token scanner.Token
		err   error
	}
)

var StopSyncTokensSet = map[scanner.TokenType]bool{
	scanner.Class:  true,
	scanner.Fun:    true,
	scanner.Var:    true,
	scanner.For:    true,
	scanner.If:     true,
	scanner.While:  true,
	scanner.Print:  true,
	scanner.Return: true,
}

func NewParser(input []scanner.Token) Parser {
	return &parser{
		input:      input,
		currentPos: 0,
	}
}

func (p *parser) Parse() (ast.Expression, *Error) {
	expression, err := p.expression()
	if err != nil {
		return nil, err
	}
	return expression, nil
}

func (p *parser) expression() (ast.Expression, *Error) {
	return p.equality()
}

func (p *parser) equality() (ast.Expression, *Error) {
	expression, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(scanner.EqualEqual, scanner.BangEqual) {
		operator := p.previous()
		right, e := p.comparison()
		if e != nil {
			return nil, e
		}
		expression = ast.NewBinary(expression, operator, right)
	}

	return expression, nil
}

func (p *parser) comparison() (ast.Expression, *Error) {
	expression, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(scanner.Greater, scanner.GreaterEqual, scanner.Less, scanner.LessEqual) {
		operator := p.previous()
		right, e := p.term()
		if e != nil {
			return nil, e
		}
		expression = ast.NewBinary(expression, operator, right)
	}

	return expression, nil
}

func (p *parser) term() (ast.Expression, *Error) {
	expression, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(scanner.Minus, scanner.Plus) {
		operator := p.previous()
		right, e := p.factor()
		if e != nil {
			return nil, e
		}
		expression = ast.NewBinary(expression, operator, right)
	}

	return expression, nil
}

func (p *parser) factor() (ast.Expression, *Error) {
	expression, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(scanner.Star, scanner.Slash) {
		operator := p.previous()
		right, e := p.unary()
		if e != nil {
			return nil, e
		}
		expression = ast.NewBinary(expression, operator, right)
	}

	return expression, nil
}

func (p *parser) unary() (ast.Expression, *Error) {
	if p.match(scanner.Bang, scanner.Minus) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return ast.NewUnary(operator, right), nil
	}

	return p.primary()
}

func (p *parser) primary() (ast.Expression, *Error) {
	if p.match(scanner.False) {
		return ast.NewLiteral(false), nil
	}
	if p.match(scanner.True) {
		return ast.NewLiteral(true), nil
	}
	if p.match(scanner.Nil) {
		return ast.NewLiteral(nil), nil
	}
	if p.match(scanner.Number, scanner.String) {
		return ast.NewLiteral(p.previous().Literal()), nil
	}
	if p.match(scanner.LeftParen) {
		expression, err := p.expression()
		if err != nil {
			return nil, err
		}
		if _, e := p.consume(scanner.RightParen, "Expect ')' after expression."); e != nil {
			return nil, e
		}
		return ast.NewGrouping(expression), nil
	}

	if !p.isAtEnd() {
		return nil, &Error{
			Token: p.peek(),
			err:   fmt.Errorf("expect exception"),
		}
	}

	return nil, nil
}

func (p *parser) match(tokenTypes ...scanner.TokenType) bool {
	for _, tt := range tokenTypes {
		if p.check(tt) {
			_ = p.advance()
			return true
		}
	}
	return false
}

func (p *parser) consume(tokenType scanner.TokenType, message string) (scanner.Token, *Error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}

	return nil, &Error{
		Token: p.peek(),
		err:   fmt.Errorf(message),
	}
}

func (p *parser) synchronize() {
	_ = p.advance()

	for !p.isAtEnd() {
		if p.previous().Type() == scanner.Semicolon {
			return
		}

		if _, found := StopSyncTokensSet[p.peek().Type()]; found {
			return
		}
	}

	_ = p.advance()
}

func (p *parser) check(tokenType scanner.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type() == tokenType
}

func (p *parser) isAtEnd() bool {
	return p.peek().Type() == scanner.EOF
}

func (p *parser) peek() scanner.Token {
	return p.input[p.currentPos]
}

func (p *parser) advance() scanner.Token {
	if !p.isAtEnd() {
		p.currentPos++
	}

	return p.previous()
}

func (p *parser) previous() scanner.Token {
	return p.input[p.currentPos-1]
}

func (e *Error) Error() string {
	return e.err.Error()
}
