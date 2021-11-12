package main

import (
	"evpeople/toyLang/evaluator"
	"evpeople/toyLang/lexer"
	"evpeople/toyLang/object"
	"evpeople/toyLang/parser"
)

func main() {

	input := `
	Step welcome
	Speak $name + ' happy'+'world'
	Listen 2,3
	Branch 'tousu',complainProc
	Branch 'zhangdan',billProc
	Silence silence
	Default defaultProc
	Step complainProc
	Speak 'ni de yi jian shi wo men de'
	Listen 2,4
	Default thanks
	Step thanks
	Speak 'thank you'
	Exit
	Step billProc
	Speak 'your zhangdan'+$amount
	Exit
	Step silenceProc
	Speak 'I can't Listen'
	Branch 'tousu',complainProc
	Branch 'zhangdan',billProc
	Silence silenceProc
	Default defaultProc
	Step defaultProc
	Speak 'I am defautProc'
	Exit
	`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParserProgram()
	env := object.NewEnvironment() //每次新建一个环境
	env.Set("name", "evpeople")
	gdf := parser.STEPINDEX
	println(gdf)
	// Eval(program, env).(*object.String).Value
	evaluator.Eval(program, env)
	// evEval2(program, env)
}
