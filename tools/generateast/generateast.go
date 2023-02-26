package main

import (
	"fmt"
	"github.com/gobeam/stringy"
	"os"
	"strings"
)

// Lox grammar:
// -----------------------------------------------------------------
//
// expression     -> literal
//                 | unary
//                 | binary
//                 | grouping ;
//
// literal        -> number | string | "true" | "false" | "nil" ;
// grouping       -> "(" expression ")" ;
// unary          -> ( "-" | "!" ) expression ;
// binary         -> expression operator expression ;
// operator       -> "==" | "!=" | "<" | "<=" | ">" | ">="
//                 | "+"  | "-"  | "*" | "/" ;
//
// -----------------------------------------------------------------

type (
	generator struct {
		builder     strings.Builder
		pkg         string
		destination string
	}
)

var (
	expressionRules = []string{
		"Assignment          : name tokens.Token, value Expression",
		"Binary              : left Expression, operator tokens.Token, right Expression",
		"Unary               : operator tokens.Token, right Expression",
		"Variable            : name tokens.Token",
		"Logical             : left Expression, operator tokens.Token, right Expression",
		"Literal             : value any",
		"Grouping            : expression Expression",
	}
	statementRules = []string{
		"BlockStatement      : statements []Statement",
		"ExpressionStatement : expression Expression",
		"IfStatement         : condition Expression, thenStatement Statement, elseStatement Statement",
		"PrintStatement      : expression Expression",
		"VarStatement        : name tokens.Token, initializer Expression",
		"WhileStatement      : condition Expression, body Statement",
	}
)

func main() {
	if len(os.Args) != 3 {
		_, _ = fmt.Fprintln(os.Stderr, "usage: generateast {package name} {destination folder} {file name}")
		os.Exit(65)
		return
	}
	pkg := os.Args[1]
	destination := os.Args[2]
	f, err := os.Create(destination)
	if err != nil {
		fmt.Println(err)
		os.Exit(65)
		return
	}
	defer f.Close()

	gen := newGenerator(pkg, destination)
	gen.defineFileHeader()

	if err = gen.defineAST("Expression", expressionRules); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(65)
		return
	}
	if err = gen.defineAST("Statement", statementRules); err != nil {
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

func (g *generator) defineFileHeader() {
	// insert header
	g.builder.WriteString("// This code is automatically generated. DO NOT EDIT.\n\n")

	// define package
	g.builder.WriteString("package ")
	g.builder.WriteString(g.pkg)
	g.builder.WriteString("\n\n")

	// import tokens
	g.builder.WriteString("import \"")
	g.builder.WriteString("github.com/mtvarkovsky/golox/pkg/tokens")
	g.builder.WriteString("\"\n\n")
}

func (g *generator) defineAST(baseName string, rules []string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can't generate ast boilerplate: %v", r)
		}
	}()

	// define expression interface
	g.builder.WriteString("type ")
	g.builder.WriteString(fmt.Sprintf("%s interface {\n", baseName))
	g.builder.WriteString(fmt.Sprintf("\tAccept(visitor %sVisitor) (any, error)\n", baseName))
	g.builder.WriteString(fmt.Sprintf("Type() %sType\n", baseName))
	g.builder.WriteString("}\n\n")

	// define expression visitor signature
	g.builder.WriteString(fmt.Sprintf("type %sVisitor = func(%s) (any, error)\n", baseName, baseName))

	// define ast type type
	g.builder.WriteString("type ")
	g.builder.WriteString(baseName + "Type int")
	g.builder.WriteString("\n\n")

	// define ast type values
	g.builder.WriteString("const (\n")
	for i, rule := range rules {
		parts := strings.Split(rule, ":")
		name := strings.Trim(parts[0], " ")
		if i == 0 {
			g.builder.WriteString(fmt.Sprintf("\t%s%sType %sType = iota\n", name, baseName, baseName))
		} else {
			g.builder.WriteString(fmt.Sprintf("\t%s%sType\n", name, baseName))
		}
	}
	g.builder.WriteString(")\n")

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
			methodName := stringy.New(fieldParts[0]).CamelCase()
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
			methodName := stringy.New(fieldParts[0]).CamelCase()
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

		// define type method
		g.builder.WriteString("func (")
		g.builder.WriteString("e *")
		g.builder.WriteString(strings.ToLower(name))
		g.builder.WriteString("")
		g.builder.WriteString(") ")
		g.builder.WriteString("Type")
		g.builder.WriteString("() ")
		g.builder.WriteString(fmt.Sprintf("%sType {\n", baseName))
		g.builder.WriteString(fmt.Sprintf("\treturn %s%sType\n", name, baseName))
		g.builder.WriteString("}\n\n")
	}

	return nil
}
