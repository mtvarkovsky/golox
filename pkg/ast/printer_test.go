package ast

import (
	"github.com/mtvarkovsky/golox/pkg/tokens"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrinter(t *testing.T) {
	expression := NewBinary(
		NewUnary(
			tokens.NewToken(tokens.Minus, "-", nil, 1, 1),
			NewLiteral(123),
		),
		tokens.NewToken(tokens.Star, "*", nil, 1, 1),
		NewGrouping(
			NewLiteral(45.67),
		),
	)
	res, _ := PrinterVisitor(expression)
	assert.Equal(t, "(* (- 123) (group 45.67))", res)
}
