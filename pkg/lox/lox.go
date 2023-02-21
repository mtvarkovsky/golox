package lox

import (
	"bufio"
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/ast"
	"github.com/mtvarkovsky/golox/pkg/parser"
	"github.com/mtvarkovsky/golox/pkg/scanner"
	"math"
	"os"
)

type (
	Interpreter interface {
		RunFile(path string)
		RunPrompt()
		Error(line int, message string)
	}

	TreeWalkInterpreter struct {
		hadError bool
	}
)

func (lox *TreeWalkInterpreter) RunFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lox.Run(string(bytes))

	if lox.hadError {
		os.Exit(65)
	}
}

func (lox *TreeWalkInterpreter) RunPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lox.Run(line)
		lox.hadError = false
	}
}

func (lox *TreeWalkInterpreter) Run(source string) {
	scnr := scanner.NewScanner(source)
	tokens, scannerErrs := scnr.ScanTokens()
	for _, err := range scannerErrs {
		lox.ScannerError(err)
		return
	}

	prsr := parser.NewParser(tokens)
	expression, parserErr := prsr.Parse()
	if parserErr != nil {
		lox.ParserError(parserErr)
		return
	}

	res, runtimeErr := ast.InterpreterVisitor(expression)
	if runtimeErr != nil {
		lox.RuntimeError(runtimeErr)
	} else {
		fmt.Println(lox.StringifyResult(res))
	}
}

func (lox *TreeWalkInterpreter) StringifyResult(res any) string {
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

func (lox *TreeWalkInterpreter) ScannerError(err *scanner.Error) {
	lox.Report(err.Line, err.Pos, "", err.Error())
}

func (lox *TreeWalkInterpreter) ParserError(err *parser.Error) {
	if err.Token.Type() == scanner.EOF {
		lox.Report(err.Token.Line(), err.Token.Position(), " at end", err.Error())
	} else {
		lox.Report(err.Token.Line(), err.Token.Position(), " at", err.Error())
	}
}

func (lox *TreeWalkInterpreter) RuntimeError(err error) {
	e, ok := err.(*ast.RuntimeError)
	if ok {
		if e.Token != nil {
			lox.Report(e.Token.Line(), e.Token.Position(), "", err.Error())
		}
		lox.Report(-1, -1, "", err.Error())
	} else {
		lox.Report(-1, -1, "", "unknown error")
	}
}

func (lox *TreeWalkInterpreter) Report(line int, pos int, where string, message string) {
	_, _ = fmt.Fprintln(
		os.Stderr,
		fmt.Sprintf("[Line %d][%d] Error %s: %s", line, pos, where, message),
	)
	lox.hadError = true
}
