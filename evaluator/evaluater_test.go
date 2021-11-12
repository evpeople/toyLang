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
		{"Speak 'hello wo pao de zui kuai'", "hello wo pao de zui kuai"},
		{"Speak 'hello'+' world'", "hello world"},
		{"Speak 'hello'+' world'+$name+' happy'", "hello world evpeople happy"},
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
	env.Set("name", "evpeople")
	// Eval(program, env).(*object.String).Value

	return Eval(program, env).Inspect()
	// return Eval(program, env).(*object.String).Value
}
func TestEvalListen(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"Listen 5,20"},
	}
	for _, tt := range tests {
		testEvalListen(tt.input)
		// testIntegerObject(t, evaluated, tt.expected)
	}
}
func testEvalListen(input string) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParserProgram()
	env := object.NewEnvironment() //每次新建一个环境
	env.Set("name", "evpeople")
	// Eval(program, env).(*object.String).Value

	Eval(program, env)
}

func TestAll(t *testing.T) {
	input := `
	Step welcome
	Speak $name + ' happy'+'world'
	Listen 2,3
	Branch "tousu",complainProc
	Branch "zhangdan",billProc
	Silence silence
	Default defaultProc
	Step complainProc
	Speak 'ni de yi jian shi wo men de'
	Listen 2,4
	Default thanks
	`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParserProgram()
	env := object.NewEnvironment() //每次新建一个环境
	env.Set("name", "evpeople")
	// Eval(program, env).(*object.String).Value

	Eval(program, env)
}
