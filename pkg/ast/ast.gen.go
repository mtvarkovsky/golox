// This code is automatically generated. DO NOT EDIT.

package ast

import "github.com/mtvarkovsky/golox/pkg/tokens"

type Expression interface {
	Accept(visitor ExpressionVisitor) (any, error)
	Type() ExpressionType
}

type ExpressionVisitor = func(Expression) (any, error)
type ExpressionType int

const (
	AssignmentExpressionType ExpressionType = iota
	BinaryExpressionType
	UnaryExpressionType
	VariableExpressionType
	LogicalExpressionType
	LiteralExpressionType
	GroupingExpressionType
)

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

func (e *assignment) Type() ExpressionType {
	return AssignmentExpressionType
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

func (e *binary) Type() ExpressionType {
	return BinaryExpressionType
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

func (e *unary) Type() ExpressionType {
	return UnaryExpressionType
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

func (e *variable) Type() ExpressionType {
	return VariableExpressionType
}


type Logical interface {
	Expression
	Left() Expression
	Operator() tokens.Token
	Right() Expression
}

type logical struct {
	left Expression
	operator tokens.Token
	right Expression
}

var _ Logical = (*logical)(nil)

func NewLogical(left Expression, operator tokens.Token, right Expression) Logical {
	return &logical{
		left: left,
		operator: operator,
		right: right,
	}
}

func (e *logical) Accept(visitor ExpressionVisitor) (any, error) {
	return visitor(e)
}
func (e *logical) Left() Expression {
	return e.left
}

func (e *logical) Operator() tokens.Token {
	return e.operator
}

func (e *logical) Right() Expression {
	return e.right
}

func (e *logical) Type() ExpressionType {
	return LogicalExpressionType
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

func (e *literal) Type() ExpressionType {
	return LiteralExpressionType
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

func (e *grouping) Type() ExpressionType {
	return GroupingExpressionType
}

type Statement interface {
	Accept(visitor StatementVisitor) (any, error)
	Type() StatementType
}

type StatementVisitor = func(Statement) (any, error)
type StatementType int

const (
	BlockStatementStatementType StatementType = iota
	ExpressionStatementStatementType
	IfStatementStatementType
	PrintStatementStatementType
	VarStatementStatementType
	WhileStatementStatementType
)

type BlockStatement interface {
	Statement
	Statements() []Statement
}

type blockStatement struct {
	statements []Statement
}

var _ BlockStatement = (*blockStatement)(nil)

func NewBlockStatement(statements []Statement) BlockStatement {
	return &blockStatement{
		statements: statements,
	}
}

func (e *blockStatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *blockStatement) Statements() []Statement {
	return e.statements
}

func (e *blockStatement) Type() StatementType {
	return BlockStatementStatementType
}


type ExpressionStatement interface {
	Statement
	Expression() Expression
}

type expressionStatement struct {
	expression Expression
}

var _ ExpressionStatement = (*expressionStatement)(nil)

func NewExpressionStatement(expression Expression) ExpressionStatement {
	return &expressionStatement{
		expression: expression,
	}
}

func (e *expressionStatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *expressionStatement) Expression() Expression {
	return e.expression
}

func (e *expressionStatement) Type() StatementType {
	return ExpressionStatementStatementType
}


type IfStatement interface {
	Statement
	Condition() Expression
	ThenStatement() Statement
	ElseStatement() Statement
}

type ifStatement struct {
	condition Expression
	thenStatement Statement
	elseStatement Statement
}

var _ IfStatement = (*ifStatement)(nil)

func NewIfStatement(condition Expression, thenStatement Statement, elseStatement Statement) IfStatement {
	return &ifStatement{
		condition: condition,
		thenStatement: thenStatement,
		elseStatement: elseStatement,
	}
}

func (e *ifStatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *ifStatement) Condition() Expression {
	return e.condition
}

func (e *ifStatement) ThenStatement() Statement {
	return e.thenStatement
}

func (e *ifStatement) ElseStatement() Statement {
	return e.elseStatement
}

func (e *ifStatement) Type() StatementType {
	return IfStatementStatementType
}


type PrintStatement interface {
	Statement
	Expression() Expression
}

type printStatement struct {
	expression Expression
}

var _ PrintStatement = (*printStatement)(nil)

func NewPrintStatement(expression Expression) PrintStatement {
	return &printStatement{
		expression: expression,
	}
}

func (e *printStatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *printStatement) Expression() Expression {
	return e.expression
}

func (e *printStatement) Type() StatementType {
	return PrintStatementStatementType
}


type VarStatement interface {
	Statement
	Name() tokens.Token
	Initializer() Expression
}

type varStatement struct {
	name tokens.Token
	initializer Expression
}

var _ VarStatement = (*varStatement)(nil)

func NewVarStatement(name tokens.Token, initializer Expression) VarStatement {
	return &varStatement{
		name: name,
		initializer: initializer,
	}
}

func (e *varStatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *varStatement) Name() tokens.Token {
	return e.name
}

func (e *varStatement) Initializer() Expression {
	return e.initializer
}

func (e *varStatement) Type() StatementType {
	return VarStatementStatementType
}


type WhileStatement interface {
	Statement
	Condition() Expression
	Body() Statement
}

type whileStatement struct {
	condition Expression
	body Statement
}

var _ WhileStatement = (*whileStatement)(nil)

func NewWhileStatement(condition Expression, body Statement) WhileStatement {
	return &whileStatement{
		condition: condition,
		body: body,
	}
}

func (e *whileStatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *whileStatement) Condition() Expression {
	return e.condition
}

func (e *whileStatement) Body() Statement {
	return e.body
}

func (e *whileStatement) Type() StatementType {
	return WhileStatementStatementType
}

