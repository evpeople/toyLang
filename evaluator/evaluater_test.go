package evaluator

import (
	"evpeople/toyLang/lexer"
	"evpeople/toyLang/object"
	"evpeople/toyLang/parser"
	"testing"
)

func TestEvalExit(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Exit", false},
	}
	for _, tt := range tests {
		evaluated := testEvalExit(tt.input)
		// testIntegerObject(t, evaluated, tt.expected)
		if evaluated {
			t.Errorf("exit is not false. got=%T (%+v)", evaluated, evaluated)
		}
	}
}
func testEvalExit(input string) bool {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParserProgram()
	env := object.NewEnvironment() //每次新建一个环境
	return Eval(program, env).(*object.Boolean).Value
}

func TestEvalSpeak(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"	Speak 'hello wo pao de zui kuai'", "hello wo pao de zui kuai"},
		{"Speak 'hello'+' world'", "hello world"},
	}
	for _, tt := range tests {
		evaluated := testEvalSpeak(tt.input)
		// testIntegerObject(t, evaluated, tt.expected)
		if evaluated != tt.expected {
			t.Errorf("speak is false. got=%T (%+v)", evaluated, evaluated)
		}
	}
}
func testEvalSpeak(input string) string {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParserProgram()
	env := object.NewEnvironment() //每次新建一个环境
	// Eval(program, env).(*object.String).Value

	return Eval(program, env).Inspect()
	// return Eval(program, env).(*object.String).Value
}
