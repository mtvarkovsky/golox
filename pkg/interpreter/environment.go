package interpreter

import (
	"fmt"
	"github.com/mtvarkovsky/golox/pkg/tokens"
)

type (
	Environment interface {
		Define(name string, value any)
		Get(name tokens.Token) (any, error)
		Assign(name tokens.Token, value any) error
		GetValues() map[string]any
		SetEnclosing(enclosing Environment)
		GetEnclosing() Environment
	}

	environment struct {
		enclosing Environment
		values    map[string]any
	}
)

func NewEnvironment(enclosing Environment) Environment {
	return &environment{
		enclosing: enclosing,
		values:    make(map[string]any),
	}
}

func (e *environment) Define(name string, value any) {
	e.values[name] = value
}

func (e *environment) Get(name tokens.Token) (any, error) {
	if value, found := e.values[name.Lexeme()]; found {
		return value, nil
	}

	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}

	return nil, &RuntimeError{
		Token: name,
		err:   fmt.Errorf("undefined variable '%s'", name.Lexeme()),
	}
}

func (e *environment) Assign(name tokens.Token, value any) error {
	if _, found := e.values[name.Lexeme()]; found {
		e.values[name.Lexeme()] = value
		return nil
	}
	if e.enclosing != nil {
		err := e.enclosing.Assign(name, value)
		if err == nil {
			return nil
		}
	}
	return &RuntimeError{
		Token: name,
		err:   fmt.Errorf("undefined variable '%s'", name.Lexeme()),
	}
}

func (e *environment) GetValues() map[string]any {
	return e.values
}

func (e *environment) SetEnclosing(enclosing Environment) {
	e.enclosing = enclosing
}

func (e *environment) GetEnclosing() Environment {
	return e.enclosing
}
