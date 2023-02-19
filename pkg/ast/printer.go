package ast

import (
	"fmt"
	"strings"
)

type PV struct {
}

func PrinterVisitor(expression Expression) string {
	switch e := expression.(type) {
	case Binary:
		return parenthesize(e.Operator().Lexeme(), e.Left(), e.Right())
	case Unary:
		return parenthesize(e.Operator().Lexeme(), e.Right())
	case Grouping:
		return parenthesize("group", e.Expression())
	case Literal:
		if e.Value() == nil {
			return "nil"
		}
		return fmt.Sprint(e.Value())
	}

	return ""
}

func parenthesize(name string, expressions ...Expression) string {
	builder := strings.Builder{}
	builder.WriteString("(")
	builder.WriteString(name)
	for _, expression := range expressions {
		builder.WriteString(" ")
		builder.WriteString(PrinterVisitor(expression))
	}
	builder.WriteString(")")

	return builder.String()
}
