package parser

import (
	"github.com/mtvarkovsky/golox/pkg/ast"
	"github.com/mtvarkovsky/golox/pkg/scanner"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	code := "(5 * (2 + 3)) - 25 == 0"
	scnr := scanner.NewScanner(code)
	tokens, errs := scnr.ScanTokens()
	assert.Empty(t, errs)
	prsr := NewParser(tokens)
	expression, err := prsr.Parse()
	stringRepr := ast.PrinterVisitor(expression)
	assert.Nil(t, err)
	assert.NotNil(t, expression)
	assert.Equal(t, "(== (- (group (* 5 (group (+ 2 3)))) 25) )", stringRepr)
}
