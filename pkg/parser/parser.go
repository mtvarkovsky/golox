package parser

import (
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/ast"
	"github.com/mtvarkovsky/golox/pkg/tokens"
)

// Lox expression grammar with precedence:
// -----------------------------------------------------------------
//
// expression     -> equality ;
// equality       -> comparison ( ( "!=" | "==") comparison  )* ;
// comparison     -> term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
// term           -> factor ( ( "-" | "+" ) factor )* ;
// factor         -> unary ( ( "/" | "*" ) unary )* ;
// unary          -> ( "!" | "-" ) unary )
//                 | primary ;
//
// primary        -> number | string | "true" | "false" | "nil"
//                 | "(" expression ")" ;
//
// -----------------------------------------------------------------

type (
	Parser interface {
		Parse() ([]ast.Statement, []*Error)
	}

	parser struct {
		input      []tokens.Token
		currentPos int
	}

	Error struct {
		Token tokens.Token
		err   error
	}
)

var (
	StopSyncTokensSet = map[tokens.TokenType]bool{
		tokens.Class:  true,
		tokens.Fun:    true,
		tokens.Var:    true,
		tokens.For:    true,
		tokens.If:     true,
		tokens.While:  true,
		tokens.Print:  true,
		tokens.Return: true,
	}
)

func NewParser(input []tokens.Token) Parser {
	return &parser{
		input:      input,
		currentPos: 0,
	}
}

// TODO: refactor this
func (p *parser) Parse() ([]ast.Statement, []*Error) {
	var statements []ast.Statement
	var errs []*Error
	for !p.isAtEnd() {
		statement, err := p.declaration()
		statements = append(statements, statement)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return statements, errs
}

func (p *parser) declaration() (ast.Statement, *Error) {
	var statement ast.Statement
	var err *Error
	if p.match(tokens.Var) {
		statement, err = p.varDeclaration()
		if err != nil {
			p.synchronize()
			return nil, err
		}
		return statement, err
	}
	statement, err = p.statement()
	if err != nil {
		p.synchronize()
		return nil, err
	}
	return statement, nil
}

func (p *parser) varDeclaration() (ast.Statement, *Error) {
	name, err := p.consume(tokens.Identifier, "Expect variable name.")
	if err != nil {
		return nil, err
	}

	var initializer ast.Expression
	if p.match(tokens.Equal) {
		initializer, err = p.expression()
		if err != nil {
			return nil, err
		}
	}

	_, err = p.consume(tokens.Semicolon, "Expect ';' after variable declaration.")
	if err != nil {
		return nil, err
	}
	return ast.NewVarStatement(name, initializer), nil
}

func (p *parser) statement() (ast.Statement, *Error) {
	if p.match(tokens.If) {
		return p.ifStatement()
	}
	if p.match(tokens.Print) {
		return p.printStatement()
	}
	if p.match(tokens.While) {
		return p.whileStatement()
	}
	if p.match(tokens.LeftBrace) {
		block, err := p.blockStatement()
		if err != nil {
			return nil, err
		}
		return ast.NewBlockStatement(block), nil
	}

	return p.expressionStatement()
}

func (p *parser) whileStatement() (ast.Statement, *Error) {
	_, err := p.consume(tokens.LeftParen, "Expect '(' after 'while'.")
	if err != nil {
		return nil, err
	}
	condition, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(tokens.RightParen, "Expect ')' after condition.")
	if err != nil {
		return nil, err
	}
	body, err := p.statement()
	if err != nil {
		return nil, err
	}
	return ast.NewWhileStatement(condition, body), nil
}

func (p *parser) ifStatement() (ast.Statement, *Error) {
	_, err := p.consume(tokens.LeftParen, "Expect '(' after 'if'.")
	if err != nil {
		return nil, err
	}
	condition, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(tokens.RightParen, "Expect ')' after 'if' condition.")
	if err != nil {
		return nil, err
	}
	thenStatement, err := p.statement()
	if err != nil {
		return nil, err
	}
	var elseStatement ast.Statement
	if p.match(tokens.Else) {
		elseStatement, err = p.statement()
	}
	if err != nil {
		return nil, err
	}
	return ast.NewIfStatement(condition, thenStatement, elseStatement), nil
}

func (p *parser) blockStatement() ([]ast.Statement, *Error) {
	var statements []ast.Statement
	for !p.check(tokens.RightBrace) && !p.isAtEnd() {
		statement, err := p.declaration()
		if err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}

	_, err := p.consume(tokens.RightBrace, "Expect '}' after block.")
	if err != nil {
		return nil, err
	}
	return statements, nil
}

func (p *parser) printStatement() (ast.Statement, *Error) {
	val, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(tokens.Semicolon, "Expect ';' after value.")
	if err != nil {
		return nil, err
	}
	return ast.NewPrintStatement(val), nil
}

func (p *parser) expressionStatement() (ast.Statement, *Error) {
	expression, err := p.expression()
	if err != nil {
		return nil, err
	}
	_, err = p.consume(tokens.Semicolon, "Expect ';' after expression.")
	if err != nil {
		return nil, err
	}
	return ast.NewExpressionStatement(expression), nil
}

func (p *parser) expression() (ast.Expression, *Error) {
	return p.assignment()
}

func (p *parser) assignment() (ast.Expression, *Error) {
	expression, err := p.or()
	if err != nil {
		return nil, err
	}

	if p.match(tokens.Equal) {
		equals := p.previous()
		value, e := p.assignment()
		if e != nil {
			return nil, e
		}

		switch expr := expression.(type) {
		case ast.Variable:
			name := expr.Name()
			return ast.NewAssignment(name, value), nil
		}

		return nil, &Error{
			Token: equals,
			err:   fmt.Errorf("invalid assignment target"),
		}
	}

	return expression, nil
}

func (p *parser) or() (ast.Expression, *Error) {
	expression, err := p.and()
	if err != nil {
		return nil, err
	}
	for p.match(tokens.Or) {
		operator := p.previous()
		right, e := p.and()
		if e != nil {
			return nil, e
		}
		expression = ast.NewLogical(expression, operator, right)
	}
	return expression, nil
}

func (p *parser) and() (ast.Expression, *Error) {
	expression, err := p.equality()
	if err != nil {
		return nil, err
	}
	for p.match(tokens.And) {
		operator := p.previous()
		right, e := p.equality()
		if e != nil {
			return nil, e
		}
		expression = ast.NewLogical(expression, operator, right)
	}
	return expression, nil
}

func (p *parser) equality() (ast.Expression, *Error) {
	expression, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(tokens.EqualEqual, tokens.BangEqual) {
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

	for p.match(tokens.Greater, tokens.GreaterEqual, tokens.Less, tokens.LessEqual) {
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

	for p.match(tokens.Minus, tokens.Plus) {
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

	for p.match(tokens.Star, tokens.Slash) {
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
	if p.match(tokens.Bang, tokens.Minus) {
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
	if p.match(tokens.False) {
		return ast.NewLiteral(false), nil
	}
	if p.match(tokens.True) {
		return ast.NewLiteral(true), nil
	}
	if p.match(tokens.Nil) {
		return ast.NewLiteral(nil), nil
	}
	if p.match(tokens.Number, tokens.String) {
		return ast.NewLiteral(p.previous().Literal()), nil
	}
	if p.match(tokens.Identifier) {
		return ast.NewVariable(p.previous()), nil
	}
	if p.match(tokens.LeftParen) {
		expression, err := p.expression()
		if err != nil {
			return nil, err
		}
		if _, e := p.consume(tokens.RightParen, "Expect ')' after expression."); e != nil {
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

func (p *parser) match(tokenTypes ...tokens.TokenType) bool {
	for _, tt := range tokenTypes {
		if p.check(tt) {
			_ = p.advance()
			return true
		}
	}
	return false
}

func (p *parser) consume(tokenType tokens.TokenType, message string) (tokens.Token, *Error) {
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
		if p.previous().Type() == tokens.Semicolon {
			return
		}

		if _, found := StopSyncTokensSet[p.peek().Type()]; found {
			return
		}
	}

	_ = p.advance()
}

func (p *parser) check(tokenType tokens.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type() == tokenType
}

func (p *parser) isAtEnd() bool {
	return p.peek().Type() == tokens.EOF
}

func (p *parser) peek() tokens.Token {
	return p.input[p.currentPos]
}

func (p *parser) advance() tokens.Token {
	if !p.isAtEnd() {
		p.currentPos++
	}

	return p.previous()
}

func (p *parser) previous() tokens.Token {
	return p.input[p.currentPos-1]
}

func (e *Error) Error() string {
	return e.err.Error()
}
