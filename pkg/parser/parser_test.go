package parser

//func TestParser(t *testing.T) {
//	code := "(5 * (2 + 3)) - 25 == 0"
//	scnr := scanner.NewScanner(code)
//	tokens, errs := scnr.ScanTokens()
//	assert.Empty(t, errs)
//	prsr := NewParser(tokens)
//	statements, err := prsr.Parse()
//	for _, statement := range statements {
//		stringRepr, _ := ast.PrinterVisitor(statement)
//		assert.Nil(t, err)
//		assert.NotNil(t, statements)
//		assert.Equal(t, "(== (- (group (* 5 (group (+ 2 3)))) 25) )", stringRepr)
//	}
//}
