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

func (e *blockstatement) Type() StatementType {
	return BlockStatementStatementType
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

func (e *expressionstatement) Type() StatementType {
	return ExpressionStatementStatementType
}


type IfStatement interface {
	Statement
	Condition() Expression
	ThenStatement() Statement
	ElseStatement() Statement
}

type ifstatement struct {
	condition Expression
	thenStatement Statement
	elseStatement Statement
}

var _ IfStatement = (*ifstatement)(nil)

func NewIfStatement(condition Expression, thenStatement Statement, elseStatement Statement) IfStatement {
	return &ifstatement{
		condition: condition,
		thenStatement: thenStatement,
		elseStatement: elseStatement,
	}
}

func (e *ifstatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *ifstatement) Condition() Expression {
	return e.condition
}

func (e *ifstatement) ThenStatement() Statement {
	return e.thenStatement
}

func (e *ifstatement) ElseStatement() Statement {
	return e.elseStatement
}

func (e *ifstatement) Type() StatementType {
	return IfStatementStatementType
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

func (e *printstatement) Type() StatementType {
	return PrintStatementStatementType
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

func (e *varstatement) Type() StatementType {
	return VarStatementStatementType
}


type WhileStatement interface {
	Statement
	Condition() Expression
	Body() Statement
}

type whilestatement struct {
	condition Expression
	body Statement
}

var _ WhileStatement = (*whilestatement)(nil)

func NewWhileStatement(condition Expression, body Statement) WhileStatement {
	return &whilestatement{
		condition: condition,
		body: body,
	}
}

func (e *whilestatement) Accept(visitor StatementVisitor) (any, error) {
	return visitor(e)
}
func (e *whilestatement) Condition() Expression {
	return e.condition
}

func (e *whilestatement) Body() Statement {
	return e.body
}

func (e *whilestatement) Type() StatementType {
	return WhileStatementStatementType
}

