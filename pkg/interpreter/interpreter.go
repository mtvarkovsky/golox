package interpreter

import (
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/ast"
	"github.com/mtvarkovsky/golox/pkg/tokens"
)

type (
	RuntimeError struct {
		err   error
		Token tokens.Token
	}
)

func Visitor(expression ast.Expression) (any, error) {
	switch e := expression.(type) {
	case ast.Binary:
		return visitBinaryExpression(e)
	case ast.Unary:
		return visitUnaryExpression(e)
	case ast.Grouping:
		return visitGrouping(e)
	case ast.Literal:
		return visitLiteral(e)
	}

	return nil, &RuntimeError{err: fmt.Errorf("unknow expression type")}
}

func visitLiteral(expression ast.Literal) (any, error) {
	return expression.Value(), nil
}

func visitGrouping(expression ast.Grouping) (any, error) {
	return evaluate(expression.Expression())
}

func evaluate(expression ast.Expression) (any, error) {
	v, err := expression.Accept(Visitor)
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
