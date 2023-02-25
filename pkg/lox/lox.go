package lox

import (
	"bufio"
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/interpreter"
	"github.com/mtvarkovsky/golox/pkg/parser"
	"github.com/mtvarkovsky/golox/pkg/scanner"
	"github.com/mtvarkovsky/golox/pkg/tokens"
	"os"
)

type (
	Interpreter interface {
		RunFile(path string)
		RunPrompt()
		Error(line int, message string)
	}

	TreeWalkInterpreter struct {
		hadError        bool
		hadRuntimeError bool
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
	if lox.hadRuntimeError {
		os.Exit(70)
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
		lox.hadRuntimeError = false
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
	statements, parserErrs := prsr.Parse()
	for _, parseErr := range parserErrs {
		lox.ParserError(parseErr)
		return
	}

	_, runtimeErr := interpreter.Interpret(statements)
	if runtimeErr != nil {
		lox.RuntimeError(runtimeErr)
	}
	//res, runtimeErr := interpreter.ExpressionVisitor(expression)
	//if runtimeErr != nil {
	//	lox.RuntimeError(runtimeErr)
	//} else {
	//	fmt.Println(interpreter.StringifyResult(res))
	//}
}

func (lox *TreeWalkInterpreter) ScannerError(err *scanner.Error) {
	lox.Report(err.Line, err.Pos, "", err.Error())
	lox.hadError = true
}

func (lox *TreeWalkInterpreter) ParserError(err *parser.Error) {
	if err.Token.Type() == tokens.EOF {
		lox.Report(err.Token.Line(), err.Token.Position(), " at end", err.Error())
	} else {
		lox.Report(err.Token.Line(), err.Token.Position(), fmt.Sprintf(" at line %d", err.Token.Line()), err.Error())
	}
	lox.hadError = true
}

func (lox *TreeWalkInterpreter) RuntimeError(err error) {
	e, ok := err.(*interpreter.RuntimeError)
	if ok {
		if e.Token != nil {
			lox.Report(e.Token.Line(), e.Token.Position(), fmt.Sprintf(" at line %d", e.Token.Line()), err.Error())
		}
		lox.Report(-1, -1, "", err.Error())
	} else {
		lox.Report(-1, -1, "", "unknown error")
	}
	lox.hadError = true
	lox.hadRuntimeError = true
}

func (lox *TreeWalkInterpreter) Report(line int, pos int, where string, message string) {
	_, _ = fmt.Fprintln(
		os.Stderr,
		fmt.Sprintf("[Line %d][%d] Error %s: %s", line, pos, where, message),
	)
}
