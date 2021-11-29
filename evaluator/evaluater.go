package evaluator

import (
	"bufio"
	"evpeople/toyLang/ast"
	"evpeople/toyLang/object"
	"evpeople/toyLang/parser"
	"fmt"
	"os"
	"strconv"
	"strings"

	// "strings"
	"time"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		// println("Program")
		return evalProgram(node, env)
	case *ast.SpeakStatement:
		// println("Speak")
		return evalSpeak(node, env)
	case *ast.ListenStatement:
		// println("Listen")
		return evalListen(node, env)
		// case *ast.BranchStatement:
		// case *ast.SilenceStatement:
		// return evalBranch(node, env) //接受这个channel作为参数
	case *ast.ExitStatement:
		// println("Exit")
		return evalExit(node, env)
	case *ast.StepStatement:
		temp := evalProgram(&ast.Program{Statements: node.ALLStatement}, env)
		_, ok := temp.(*object.String)
		if ok {
			temp.(*object.String).Value += "Step" //TODO:假如不能转换成功，说明是Exit//运行完了返回的
			return temp
		} else {
			break
		}
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
	time.Sleep(time.Duration(b) * 1 * time.Second)
	c := make(chan string)

	go func() {
		bf := time.Now()
		fmt.Println("请输入答案")
		a := bufio.NewScanner(os.Stdin)

		var ans string
		for a.Scan() {
			if time.Since(bf) > time.Duration(5)*time.Second {
				// fmt.Println("chao shi")
				break
			}
			ans += a.Text()
			fmt.Println(ans, time.Since(bf))
			c <- ans
		}
		// fmt.Scanln(&ans)
	}()

	select {
	case m := <-c:
		result.Value = "Listen" + m
		handle(m)
	case <-time.After(5 * time.Second * time.Duration(e)):
		result.Value = "Listen" + "Silence"
		handle("time out")
	}
	close(c)

	time.Sleep(time.Duration(b) * 2 * time.Second)
	return &result
}
func handle(m string) {
	fmt.Println(m)
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
	fmt.Println(result.Value)
	return &result
}
func evalExit(program *ast.ExitStatement, env *object.Environment) object.Object {
	var result object.Boolean
	return &result
}

// func evalProgram(program *ast.Program, env *object.Environment) object.Object {
// 	var result object.Object
// 	for _, statement := range program.Statements {
// 		result = Eval(statement, env)
// 		temp := result.Inspect()
// 		var temp2 string
// 		if index := strings.Index(temp, "Listen"); index != -1 {
// 			temp2 = temp[6:]
// 			result.(*object.String).Value = temp2
// 			return result
// 		}

// 		if index := strings.Index(temp, "Step"); index != -1 {
// 			temp2 = temp[:index]
// 			result.(*object.String).Value = temp2
// 			// temp3 := statement.(*ast.StepStatement).CaseBranch[temp2]
// 			return result
// 		}

// 		//TODO:result 的类型，如果为Listen的ans的话，添加到env中。
// 		//TODO:在CaseBranch中，首先假设是Silence，然后在CaseBranch里读取对应的ident名字
// 		//TODO:将ident包裹在Object中，直接在此return
// 		//TODO: return

// 		// switch result {
// 		// case true:
// 		// 	return true
// 		// case *object.Error:
// 		// 	return result //包装的值，处理错误的传递
// 		// }
// 		// result, ok := result.(*object.Boolean)
// 		// if !ok {
// 		// 	return result
// 		// }
// 		// if !result.Value {
// 		// 	return result
// 		// }
// 	}
// 	return result
// }

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for i := 0; i < len(program.Statements); i++ {
		statement := program.Statements[i]
		result = Eval(statement, env)
		temp := result.Inspect()
		var temp2 string
		if index := strings.Index(temp, "Listen"); index != -1 {
			temp2 = temp[6:]
			result.(*object.String).Value = temp2
			return result
		}

		if index := strings.Index(temp, "Step"); index != -1 {
			temp2 = temp[:index]
			result.(*object.String).Value = temp2
			temp2 = "tousu"
			temp3 := statement.(*ast.StepStatement).CaseBranch[temp2]
			temp4 := parser.STEPINDEX[temp3]
			i = temp4 - 1
			// i = 4
			// return result
		}
	}
	// for _, statement := range program.Statements {

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
	// }
	return result
}
