package lox

import (
	"bufio"
	"fmt"
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
		lox.Error(err.Line, err.Pos, err.Error())
	}

	for _, t := range tokens {
		fmt.Println(t)
	}
}

func (lox *TreeWalkInterpreter) Error(line int, pos int, message string) {
	lox.Report(line, pos, "", message)
}

func (lox *TreeWalkInterpreter) Report(line int, pos int, where string, message string) {
	_, _ = fmt.Fprintln(
		os.Stderr,
		fmt.Sprintf("[Line %d][%d] Error %s: %s", line, pos, where, message),
	)
	lox.hadError = true
}
