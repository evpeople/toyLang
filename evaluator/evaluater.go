package evaluator

import (
	"evpeople/toyLang/ast"
	"evpeople/toyLang/object"
	"strconv"
	"time"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.SpeakStatement:
		return evalSpeak(node, env)
	case *ast.ListenStatement:
		return evalListen(node, env)
	case *ast.BranchStatement:
	case *ast.SilenceStatement:
	case *ast.ExitStatement:
		return evalExit(node, env)
	}
	var result object.Boolean
	return &result
}
func evalListen(p *ast.ListenStatement, env *object.Environment) object.Object {
	var result object.String
	s := p.Expression.TokenLiteral()
	// a := len(s)
	begin := ""
	for i := 9; s[i] != '\n'; i++ {
		begin += string(s[i])
	}

	end := ""
	for i := 18; i < len(s); i++ {
		end += string(s[i])
	}
	b, _ := strconv.Atoi(begin)
	e, _ := strconv.Atoi(end)
	time.Sleep(time.Duration(b+e) * time.Second)
	return &result
}
func evalSpeak(program *ast.SpeakStatement, env *object.Environment) object.Object {
	var result object.String
	if program.Expression.(*ast.SentenceStatement).DollarMap == nil {
		result.Value = program.Expression.TokenLiteral()
	} else {
		for k, _ := range program.Expression.(*ast.SentenceStatement).DollarMap {
			if realVar, ok := env.Get(k); ok {
				program.Expression.(*ast.SentenceStatement).DollarMap[k] = realVar
			}
		}
		result.Value = program.Expression.(*ast.SentenceStatement).RealTokenLiteral()
	}
	return &result
}
func evalExit(program *ast.ExitStatement, env *object.Environment) object.Object {
	var result object.Boolean
	return &result
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
		// switch result {
		// case true:
		// 	return true
		// case *object.Error:
		// 	return result //包装的值，处理错误的传递
		// }
		// result, ok := result.(*object.Boolean)
		// if !ok {
		// 	return result
		// }
		// if !result.Value {
		// 	return result
		// }
	}
	return result
}
