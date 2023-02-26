package interpreter

import (
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/ast"
	"github.com/mtvarkovsky/golox/pkg/tokens"
	"math"
	"os"
)

type (
	RuntimeError struct {
		err   error
		Token tokens.Token
	}
)

var Env = NewEnvironment(nil)

// TODO: refactor this
func Interpret(statements []ast.Statement) (any, error) {
	for _, statement := range statements {
		err := execute(statement)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func execute(statement ast.Statement) error {
	_, err := statement.Accept(StatementVisitor)
	return err
}

func StatementVisitor(statement ast.Statement) (any, error) {
	switch statement.Type() {
	case ast.VarStatementStatementType:
		return visitVarStatement(statement.(ast.VarStatement))
	case ast.BlockStatementStatementType:
		return visitBlockStatement(statement.(ast.BlockStatement))
	case ast.WhileStatementStatementType:
		return visitWhileStatement(statement.(ast.WhileStatement))
	case ast.ExpressionStatementStatementType:
		return visitExpressionStatement(statement.(ast.ExpressionStatement))
	case ast.PrintStatementStatementType:
		return visitPrintStatement(statement.(ast.PrintStatement))
	case ast.IfStatementStatementType:
		return visitIfStatement(statement.(ast.IfStatement))
	}

	return nil, &RuntimeError{err: fmt.Errorf("unknow statement type")}
}

func visitIfStatement(statement ast.IfStatement) (any, error) {
	res, err := evaluate(statement.Condition())
	if err != nil {
		return nil, err
	}
	b, err := toBoolean(res)
	if b {
		err = execute(statement.ThenStatement())
		if err != nil {
			return nil, err
		}
	} else if statement.ElseStatement() != nil {
		err = execute(statement.ElseStatement())
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func visitWhileStatement(statement ast.WhileStatement) (any, error) {
	cond, err := evaluate(statement.Condition())
	if err != nil {
		return nil, err
	}
	b, err := toBoolean(cond)
	if err != nil {
		return nil, err
	}
	for b {
		err = execute(statement.Body())
		if err != nil {
			return nil, err
		}
		cond, err = evaluate(statement.Condition())
		if err != nil {
			return nil, err
		}
		b, err = toBoolean(cond)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func visitPrintStatement(statement ast.PrintStatement) (any, error) {
	val, err := evaluate(statement.Expression())
	if err != nil {
		return nil, err
	}
	_, err = fmt.Fprintln(os.Stdout, StringifyResult(val))
	return nil, err
}

func visitExpressionStatement(statement ast.ExpressionStatement) (any, error) {
	_, err := evaluate(statement.Expression())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func visitVarStatement(statement ast.VarStatement) (any, error) {
	var value any
	var err error
	if statement.Initializer() != nil {
		value, err = evaluate(statement.Initializer())
		if err != nil {
			return nil, err
		}
	}
	Env.Define(statement.Name().Lexeme(), value)
	return nil, nil
}

func visitBlockStatement(statement ast.BlockStatement) (any, error) {
	err := executeBlock(statement.Statements(), NewEnvironment(Env))
	return nil, err
}

func executeBlock(statements []ast.Statement, env Environment) error {
	outerEnv := Env
	for k, v := range Env.GetValues() {
		outerEnv.Define(k, v)
	}
	outerEnv.SetEnclosing(Env.GetEnclosing())

	var err error
	Env = env
	for _, statement := range statements {
		err = execute(statement)
		if err != nil {
			return err
		}
	}
	Env = outerEnv
	return nil
}

func StringifyResult(res any) string {
	if res == nil {
		return "nil"
	}
	if _, ok := res.(float64); ok {
		if res.(float64) == math.Trunc(res.(float64)) {
			return fmt.Sprintf("%.0f", res)
		}

		return fmt.Sprintf("%f", res)
	}
	return fmt.Sprint(res)
}

func ExpressionVisitor(expression ast.Expression) (any, error) {
	switch expression.Type() {
	case ast.AssignmentExpressionType:
		return visitAssignmentExpression(expression.(ast.Assignment))
	case ast.BinaryExpressionType:
		return visitBinaryExpression(expression.(ast.Binary))
	case ast.LogicalExpressionType:
		return visitLogical(expression.(ast.Logical))
	case ast.UnaryExpressionType:
		return visitUnaryExpression(expression.(ast.Unary))
	case ast.LiteralExpressionType:
		return visitLiteral(expression.(ast.Literal))
	case ast.GroupingExpressionType:
		return visitGrouping(expression.(ast.Grouping))
	case ast.VariableExpressionType:
		return visitVariable(expression.(ast.Variable))
	}

	return nil, &RuntimeError{err: fmt.Errorf("unknow expression type")}
}

func visitLogical(expression ast.Logical) (any, error) {
	left, err := evaluate(expression.Left())
	if err != nil {
		return nil, err
	}
	b, err := toBoolean(left)
	if err != nil {
		return nil, err
	}
	if expression.Operator().Type() == tokens.Or {
		if b {
			return left, nil
		}
	} else {
		if !b {
			return left, nil
		}
	}

	return evaluate(expression.Right())
}

func visitVariable(expression ast.Variable) (any, error) {
	return Env.Get(expression.Name())
}

func visitAssignmentExpression(expression ast.Assignment) (any, error) {
	value, err := evaluate(expression.Value())
	if err != nil {
		return nil, err
	}
	err = Env.Assign(expression.Name(), value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func visitLiteral(expression ast.Literal) (any, error) {
	return expression.Value(), nil
}

func visitGrouping(expression ast.Grouping) (any, error) {
	return evaluate(expression.Expression())
}

func evaluate(expression ast.Expression) (any, error) {
	v, err := expression.Accept(ExpressionVisitor)
	if err != nil {
		return nil, &RuntimeError{err: err}
	}
	return v, nil
}

func visitUnaryExpression(expression ast.Unary) (v any, err error) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
			err = &RuntimeError{err: fmt.Errorf("can't interpret unary expression: %v", r)}
		}
	}()

	right, err := evaluate(expression.Right())
	if err != nil {
		return nil, err
	}

	switch expression.Operator().Type() {
	case tokens.Bang:
		v, err = toBoolean(right)
		return v, err
	case tokens.Minus:
		e := checkNumberOperands(expression.Operator(), right)
		if e != nil {
			return nil, e
		}
		return -right.(float64), nil
	}

	return nil, nil
}

func toBoolean(o any) (v bool, err error) {
	if o == nil {
		return false, nil
	}

	switch o.(type) {
	case bool:
		return o.(bool), nil
	}

	return true, nil
}

func visitBinaryExpression(expression ast.Binary) (v any, err error) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
			err = &RuntimeError{err: fmt.Errorf("can't interpret binary expression: %v", r)}
		}
	}()

	left, err := evaluate(expression.Left())
	if err != nil {
		return nil, err
	}
	right, err := evaluate(expression.Right())
	if err != nil {
		return nil, err
	}

	switch expression.Operator().Type() {
	case tokens.Greater:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) > right.(float64), nil
	case tokens.GreaterEqual:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) >= right.(float64), nil
	case tokens.Less:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) < right.(float64), nil
	case tokens.LessEqual:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) <= right.(float64), nil
	case tokens.Minus:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) - right.(float64), nil
	case tokens.Slash:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) / right.(float64), nil
	case tokens.Star:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) * right.(float64), nil
	case tokens.Plus:
		if _, lIsNumber := left.(float64); lIsNumber {
			if _, rIsNumber := right.(float64); rIsNumber {
				return left.(float64) + right.(float64), nil
			}
		}
		if _, lIsString := left.(string); lIsString {
			if _, rIsString := right.(string); rIsString {
				return left.(string) + right.(string), nil
			}
		}
		return nil, &RuntimeError{err: fmt.Errorf("operands must be both numbers or both strings"), Token: expression.Operator()}
	case tokens.EqualEqual:
		return isEqual(left, right)
	case tokens.BangEqual:
		val, e := isEqual(left, right)
		return !val, e
	}

	return nil, nil
}

func isEqual(left any, right any) (bool, error) {
	if left == nil && right == nil {
		return true, nil
	}
	if left == nil {
		return false, nil
	}
	return fmt.Sprint(left) == fmt.Sprint(right), nil
}

func checkNumberOperands(operator tokens.Token, operands ...any) error {
	for _, operand := range operands {
		if _, ok := operand.(float64); !ok {
			return &RuntimeError{err: fmt.Errorf("operand must be a number"), Token: operator}
		}
	}

	return nil
}

func (re *RuntimeError) Error() string {
	return re.err.Error()
}
