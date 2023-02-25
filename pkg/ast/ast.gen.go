// This code is automatically generated. DO NOT EDIT.

package ast

import "github.com/mtvarkovsky/golox/pkg/tokens"

type Expression interface {
	Accept(visitor ExpressionVisitor) (any, error)
}

type ExpressionVisitor = func(Expression) (any, error)

type Assignment interface {
	Expression
	Name() tokens.Token
	Value() Expression
}

type assignment struct {
	name tokens.Token
	value Expression
}

var _ Assignment = (*assignment)(nil)

func NewAssignment(name tokens.Token, value Expression) Assignment {
	return &assignment{
		name: name,
		value: value,
	}
}

func (e *assignment) Accept(visitor ExpressionVisitor) (any, error) {
	return visitor(e)
}
func (e *assignment) Name() tokens.Token {
	return e.name
}

func (e *assignment) Value() Expression {
	return e.value
}


type Binary interface {
	Expression
	Left() Expression
	Operator() tokens.Token
	Right() Expression
}

type binary struct {
	left Expression
	operator tokens.Token
	right Expression
}

var _ Binary = (*binary)(nil)

func NewBinary(left Expression, operator tokens.Token, right Expression) Binary {
	return &binary{
		left: left,
		operator: operator,
		right: right,
	}
}

func (e *binary) Accept(visitor ExpressionVisitor) (any, error) {
	return visitor(e)
}
func (e *binary) Left() Expression {
	return e.left
}

func (e *binary) Operator() tokens.Token {
	return e.operator
}

func (e *binary) Right() Expression {
	return e.right
}


type Unary interface {
	Expression
	Operator() tokens.Token
	Right() Expression
}

type unary struct {
	operator tokens.Token
	right Expression
}

var _ Unary = (*unary)(nil)

func NewUnary(operator tokens.Token, right Expression) Unary {
	return &unary{
		operator: operator,
		right: right,
	}
}

func (e *unary) Accept(visitor ExpressionVisitor) (any, error) {
	return visitor(e)
}
func (e *unary) Operator() tokens.Token {
	return e.operator
}

func (e *unary) Right() Expression {
	return e.right
}


type Variable interface {
	Expression
	Name() tokens.Token
}

type variable struct {
	name tokens.Token
}

var _ Variable = (*variable)(nil)

func NewVariable(name tokens.Token) Variable {
	return &variable{
		name: name,
	}
}

func (e *variable) Accept(visitor ExpressionVisitor) (any, error) {
	return visitor(e)
}
func (e *variable) Name() tokens.Token {
	return e.name
}


type Literal interface {
	Expression
	Value() any
}

type literal struct {
	value any
}

var _ Literal = (*literal)(nil)

func NewLiteral(value any) Literal {
	return &literal{
		value: value,
	}
}

func (e *literal) Accept(visitor ExpressionVisitor) (any, error) {
	return visitor(e)
}
func (e *literal) Value() any {
	return e.value
}


type Grouping interface {
	Expression
	Expression() Expression
}

type grouping struct {
	expression Expression
}

var _ Grouping = (*grouping)(nil)

func NewGrouping(expression Expression) Grouping {
	return &grouping{
		expression: expression,
	}
}

func (e *grouping) Accept(visitor ExpressionVisitor) (any, error) {
	return visitor(e)
}
func (e *grouping) Expression() Expression {
	return e.expression
}

type Statement interface {
	Accept(visitor StatementVisitor) (any, error)
}

type StatementVisitor = func(Statement) (any, error)

type BlockStatement interface {
	Statement
	Statements() []Statement
}

type blockstatement struct {
	statements []Statement
}

var _ BlockStatement = (*blockstatement)(nil)

func NewBlockStatement(statements []Statement) BlockStatement {
	return &blockstatement{
		statements: statements,
	}
}

func (e *blockstatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *blockstatement) Statements() []Statement {
	return e.statements
}


type ExpressionStatement interface {
	Statement
	Expression() Expression
}

type expressionstatement struct {
	expression Expression
}

var _ ExpressionStatement = (*expressionstatement)(nil)

func NewExpressionStatement(expression Expression) ExpressionStatement {
	return &expressionstatement{
		expression: expression,
	}
}

func (e *expressionstatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *expressionstatement) Expression() Expression {
	return e.expression
}


type PrintStatement interface {
	Statement
	Expression() Expression
}

type printstatement struct {
	expression Expression
}

var _ PrintStatement = (*printstatement)(nil)

func NewPrintStatement(expression Expression) PrintStatement {
	return &printstatement{
		expression: expression,
	}
}

func (e *printstatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *printstatement) Expression() Expression {
	return e.expression
}


type VarStatement interface {
	Statement
	Name() tokens.Token
	Initializer() Expression
}

type varstatement struct {
	name tokens.Token
	initializer Expression
}

var _ VarStatement = (*varstatement)(nil)

func NewVarStatement(name tokens.Token, initializer Expression) VarStatement {
	return &varstatement{
		name: name,
		initializer: initializer,
	}
}

func (e *varstatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *varstatement) Name() tokens.Token {
	return e.name
}

func (e *varstatement) Initializer() Expression {
	return e.initializer
}

