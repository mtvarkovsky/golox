package lox

import (
	"bufio"
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/ast"
	"github.com/mtvarkovsky/golox/pkg/parser"
	"github.com/mtvarkovsky/golox/pkg/scanner"
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
	tokens, errs := scnr.ScanTokens()
	for _, err := range errs {
		lox.ScannerError(err)
	}

	prsr := parser.NewParser(tokens)
	expression, err := prsr.Parse()
	if err != nil {
		lox.ParserError(err)
	}

	fmt.Println(ast.PrinterVisitor(expression))
}

func (lox *TreeWalkInterpreter) ScannerError(err *scanner.Error) {
	lox.Report(err.Line, err.Pos, "", err.Error())
}

func (lox *TreeWalkInterpreter) ParserError(err *parser.Error) {
	if err.Token.Type() == scanner.EOF {
		lox.Report(err.Token.Line(), err.Token.Position(), " at end", err.Error())
	} else {
		lox.Report(err.Token.Line(), err.Token.Position(), " at '", err.Error())
	}
}

func (lox *TreeWalkInterpreter) Report(line int, pos int, where string, message string) {
	_, _ = fmt.Fprintln(
		os.Stderr,
		fmt.Sprintf("[Line %d][%d] Error %s: %s", line, pos, where, message),
	)
	lox.hadError = true
}
