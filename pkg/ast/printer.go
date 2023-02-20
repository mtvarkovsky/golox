package ast

import (
	"fmt"
	"strings"
)

func PrinterVisitor(expression Expression) (string, error) {
	switch e := expression.(type) {
	case Binary:
		return parenthesize(e.Operator().Lexeme(), e.Left(), e.Right()), nil
	case Unary:
		return parenthesize(e.Operator().Lexeme(), e.Right()), nil
	case Grouping:
		return parenthesize("group", e.Expression()), nil
	case Literal:
		if e.Value() == nil {
			return "nil", nil
		}
		return fmt.Sprint(e.Value()), nil
	}

	return "", nil
}

func parenthesize(name string, expressions ...Expression) string {
	builder := strings.Builder{}
	builder.WriteString("(")
	builder.WriteString(name)
	for _, expression := range expressions {
		builder.WriteString(" ")
		s, _ := PrinterVisitor(expression)
		builder.WriteString(s)
	}
	builder.WriteString(")")

	return builder.String()
}
