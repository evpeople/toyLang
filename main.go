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
	Speak $name + ' hello'+' ,world do you want to tousu or zhangdan Listen'
	Listen 2,3
	Branch 'tousu',complainProc
	Branch 'zhangdan',billProc
	Silence silenceProc
	Default defaultProc
	
	Step complainProc
	Speak 'I am tousu complainProc'
	Listen 2,4
	Default thanks
	
	Step thanks
	Speak 'I am thanks thank you'
	Exit
	
	Step billProc
	Speak 'I am zhangdan billProc your zhangdan'+ $amount
	Exit
	
	Step silenceProc
	Speak 'I am silence  I can not listen'
	Listen 2,4
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
	env.Set("amount", "1000")
	// Eval(program, env).(*object.String).Value
	evaluator.Eval(program, env)
	// evEval2(program, env)
}
