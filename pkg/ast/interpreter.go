//go:generate go run ../../tools/generateast/ ast ast.gen.go
package ast

import (
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/scanner"
)

type (
	RuntimeError struct {
		err   error
		Token scanner.Token
	}
)

func InterpreterVisitor(expression Expression) (any, error) {
	switch e := expression.(type) {
	case Binary:
		return visitBinaryExpression(e)
	case Unary:
		return visitUnaryExpression(e)
	case Grouping:
		return visitGrouping(e)
	case Literal:
		return visitLiteral(e)
	}

	return nil, &RuntimeError{err: fmt.Errorf("unknow expression type")}
}

func visitLiteral(expression Literal) (any, error) {
	return expression.Value(), nil
}

func visitGrouping(expression Grouping) (any, error) {
	return evaluate(expression.Expression())
}

func evaluate(expression Expression) (any, error) {
	v, err := expression.Accept(InterpreterVisitor)
	if err != nil {
		return nil, &RuntimeError{err: err}
	}
	return v, nil
}

func visitUnaryExpression(expression Unary) (v any, err error) {
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
	case scanner.Bang:
		v, err = toBoolean(right)
		return v, err
	case scanner.Minus:
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

func visitBinaryExpression(expression Binary) (v any, err error) {
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
	case scanner.Greater:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) > right.(float64), nil
	case scanner.GreaterEqual:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) >= right.(float64), nil
	case scanner.Less:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) < right.(float64), nil
	case scanner.LessEqual:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) <= right.(float64), nil
	case scanner.Minus:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) - right.(float64), nil
	case scanner.Slash:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) / right.(float64), nil
	case scanner.Star:
		e := checkNumberOperands(expression.Operator(), left, right)
		if e != nil {
			return nil, e
		}
		return left.(float64) * right.(float64), nil
	case scanner.Plus:
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
	case scanner.Equal:
		return isEqual(left, right)
	case scanner.BangEqual:
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

func checkNumberOperands(operator scanner.Token, operands ...any) error {
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
