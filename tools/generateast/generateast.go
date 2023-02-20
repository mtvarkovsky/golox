package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
)

type (
	generator struct {
		builder     strings.Builder
		pkg         string
		destination string
	}
)

var (
	expressionRules = []string{
		"Binary   : left Expression, operator scanner.Token, right Expression",
		"Unary    : operator scanner.Token, right Expression",
		"Literal  : value any",
		"Grouping : expression Expression",
	}
)

func main() {
	if len(os.Args) != 3 {
		_, _ = fmt.Fprintln(os.Stderr, "usage: generatelatest {package name} {destination file name}")
		os.Exit(65)
		return
	}
	pkg := os.Args[1]
	destination := os.Args[2]
	f, err := os.Create(destination)

	gen := newGenerator(pkg, destination)
	if err = gen.defineAST("Expression", expressionRules); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(65)
		return
	}

	_, _ = fmt.Fprint(f, gen.builder.String())
}

func newGenerator(pkg string, destination string) *generator {
	return &generator{
		builder:     strings.Builder{},
		pkg:         pkg,
		destination: destination,
	}
}

func (g *generator) defineAST(baseName string, rules []string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can't generate ast boilerplate: %v", r)
		}
	}()

	// insert header
	g.builder.WriteString("// This code is automatically generated. DO NOT EDIT.\n\n")

	// define package
	g.builder.WriteString("package ")
	g.builder.WriteString(g.pkg)
	g.builder.WriteString("\n\n")

	// import scanner
	g.builder.WriteString("import \"")
	g.builder.WriteString("github.com/mtvarkovsky/golox/pkg/scanner")
	g.builder.WriteString("\"\n\n")

	// define expression interface
	g.builder.WriteString("type ")
	g.builder.WriteString(fmt.Sprintf("%s interface {\n", baseName))
	g.builder.WriteString(fmt.Sprintf("\tAccept(visitor %sVisitor) (any, error)\n", baseName))
	g.builder.WriteString("}\n\n")

	// define expression visitor signature
	g.builder.WriteString(fmt.Sprintf("type %sVisitor = func(%s) (any, error)\n", baseName, baseName))

	// go through all rules
	for _, rule := range rules {
		parts := strings.Split(rule, ":")
		name := strings.Trim(parts[0], " ")
		args := strings.Trim(parts[1], " ")
		fields := strings.Split(args, ",")

		// define interface
		g.builder.WriteString("\n")
		g.builder.WriteString("type ")
		g.builder.WriteString(name)
		g.builder.WriteString(" interface {\n")
		g.builder.WriteString(fmt.Sprintf("\t%s\n", baseName))
		// for each field add getter
		for _, field := range fields {
			fieldParts := strings.Split(strings.Trim(field, " "), " ")
			methodName := cases.Title(language.Und).String(fieldParts[0])
			methodRType := fieldParts[1]
			g.builder.WriteString("\t")
			g.builder.WriteString(methodName)
			g.builder.WriteString("() ")
			g.builder.WriteString(methodRType)
			g.builder.WriteString("\n")
		}
		g.builder.WriteString("}\n\n")

		// define struct
		g.builder.WriteString("type ")
		g.builder.WriteString(strings.ToLower(name))
		g.builder.WriteString(" struct {\n")
		// for each field add it to struct definition
		for _, field := range fields {
			g.builder.WriteString("\t")
			g.builder.WriteString(strings.Trim(field, " "))
			g.builder.WriteString("\n")
		}
		g.builder.WriteString("}\n\n")

		// add compile-time check for interface implementation
		g.builder.WriteString("var _ ")
		g.builder.WriteString(name)
		g.builder.WriteString(" = (*")
		g.builder.WriteString(strings.ToLower(name))
		g.builder.WriteString(")(nil)\n\n")

		// define constructor
		g.builder.WriteString("func New")
		g.builder.WriteString(name)
		g.builder.WriteString("")
		g.builder.WriteString("(")
		g.builder.WriteString(args)
		g.builder.WriteString(") ")
		g.builder.WriteString(name)
		g.builder.WriteString("")
		g.builder.WriteString(" {\n")
		g.builder.WriteString("\treturn &")
		g.builder.WriteString(strings.ToLower(name))
		g.builder.WriteString("")
		g.builder.WriteString("{\n")
		// for each field initialize it
		for _, field := range fields {
			fieldParts := strings.Split(strings.Trim(field, " "), " ")
			fieldName := fieldParts[0]
			g.builder.WriteString("\t\t")
			g.builder.WriteString(fieldName)
			g.builder.WriteString(": ")
			g.builder.WriteString(fieldName)
			g.builder.WriteString(",\n")
		}
		g.builder.WriteString("\t}\n")
		g.builder.WriteString("}\n\n")

		// define Accept method
		g.builder.WriteString("func (")
		g.builder.WriteString("e *")
		g.builder.WriteString(strings.ToLower(name))
		g.builder.WriteString("")
		g.builder.WriteString(") ")
		g.builder.WriteString(fmt.Sprintf("Accept(visitor %sVisitor) (any, error) {\n", baseName))
		g.builder.WriteString("\treturn visitor(e)")
		g.builder.WriteString("\n}\n")

		// for each field define getter
		for _, field := range fields {
			fieldParts := strings.Split(strings.Trim(field, " "), " ")
			fieldName := fieldParts[0]
			methodName := cases.Title(language.Und).String(fieldParts[0])
			methodRType := fieldParts[1]

			g.builder.WriteString("func (")
			g.builder.WriteString("e *")
			g.builder.WriteString(strings.ToLower(name))
			g.builder.WriteString("")
			g.builder.WriteString(") ")
			g.builder.WriteString(methodName)
			g.builder.WriteString("() ")
			g.builder.WriteString(methodRType)
			g.builder.WriteString(" {\n")
			g.builder.WriteString("\treturn e.")
			g.builder.WriteString(fieldName)
			g.builder.WriteString("\n}\n\n")
		}
	}

	return nil
}
