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
		// println("Program")
		return evalProgram(node, env)
		//TODO:从ast.StepStatement返回到这里，然后解包装，获取ident，然后从parser.STEPINDEX获取对应的数字。
		//TODO：然后直接Eval(获取的下标对应的Statement数组的元素)
	case *ast.SpeakStatement:
		// println("Speak")
		return evalSpeak(node, env)
	case *ast.ListenStatement:
		// println("Listen")
		return evalListen(node, env) //TODO:返回一个字符串，进行一个时间延迟的channel，先sleep5s，然后wait最长时间，另一个goroutine里有键盘读取的代码，最后返回，读取到的信息。
		// case *ast.BranchStatement:
		// case *ast.SilenceStatement:
		// return evalBranch(node, env) //接受这个channel作为参数
	case *ast.ExitStatement:
		// println("Exit")
		return evalExit(node, env)
	case *ast.StepStatement:
		return evalProgram(&ast.Program{Statements: node.ALLStatement}, env)
		//TODO:从evalProgram 返回到这里，读取数据，然后再加一层包装，直接返回。
	}
	// println(node)
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
	time.Sleep(time.Duration(b+e) * 1 * time.Second)
	return &result
}
func evalSpeak(program *ast.SpeakStatement, env *object.Environment) object.Object {
	var result object.String
	if program.Expression.(*ast.SentenceStatement).DollarMap == nil {
		result.Value = program.Expression.TokenLiteral()
	} else {
		for k := range program.Expression.(*ast.SentenceStatement).DollarMap {
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

		//TODO:result 的类型，如果为Listen的ans的话，添加到env中。
		//TODO:在CaseBranch中，首先假设是Silence，然后在CaseBranch里读取对应的ident名字
		//TODO:将ident包裹在Object中，直接在此return
		//TODO: return

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
