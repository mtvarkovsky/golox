package main

import (
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/lox"
	"os"
)

func main() {
	inst := lox.TreeWalkInterpreter{}
	if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		inst.RunFile(os.Args[1])
	} else {
		inst.RunPrompt()
	}
}
