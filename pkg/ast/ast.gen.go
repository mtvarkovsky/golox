// This code is automatically generated. DO NOT EDIT.

package ast

import "github.com/mtvarkovsky/golox/pkg/scanner"

type Expression interface {
	Accept(visitor ExpressionVisitor) (any, error)
}

type ExpressionVisitor = func(Expression) (any, error)

type Binary interface {
	Expression
	Left() Expression
	Operator() scanner.Token
	Right() Expression
}

type binary struct {
	left Expression
	operator scanner.Token
	right Expression
}

var _ Binary = (*binary)(nil)

func NewBinary(left Expression, operator scanner.Token, right Expression) Binary {
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

func (e *binary) Operator() scanner.Token {
	return e.operator
}

func (e *binary) Right() Expression {
	return e.right
}


type Unary interface {
	Expression
	Operator() scanner.Token
	Right() Expression
}

type unary struct {
	operator scanner.Token
	right Expression
}

var _ Unary = (*unary)(nil)

func NewUnary(operator scanner.Token, right Expression) Unary {
	return &unary{
		operator: operator,
		right: right,
	}
}

func (e *unary) Accept(visitor ExpressionVisitor) (any, error) {
	return visitor(e)
}
func (e *unary) Operator() scanner.Token {
	return e.operator
}

func (e *unary) Right() Expression {
	return e.right
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

