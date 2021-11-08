package evaluator

import (
	"evpeople/toyLang/ast"
	"evpeople/toyLang/object"
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
func evalSpeak(program *ast.SpeakStatement, env *object.Environment) object.Object {
	var result object.String
	result.Value = program.Expression.TokenLiteral()
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
