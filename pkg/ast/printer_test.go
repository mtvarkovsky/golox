package ast

import (
	"github.com/mtvarkovsky/golox/pkg/scanner"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrinter(t *testing.T) {
	expression := NewBinary(
		NewUnary(
			scanner.NewToken(scanner.Minus, "-", nil, 1, 1),
			NewLiteral(123),
		),
		scanner.NewToken(scanner.Star, "*", nil, 1, 1),
		NewGrouping(
			NewLiteral(45.67),
		),
	)
	res := PrinterVisitor(expression)
	assert.Equal(t, "(* (- 123) (group 45.67))", res)
}
