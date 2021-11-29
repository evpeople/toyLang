//Evaluate是实际解释执行的代码所在的包
package evaluator

import (
	"evpeople/toyLang/ast"
	"evpeople/toyLang/object"
	"evpeople/toyLang/parser"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var Eval_Conn map[int]net.Conn

//Eval函数是实际执行的函数，为不同节点调用不同的执行方法，main函数调用其来执行Step节点的集合，evalProgram函数调用其来遍历Step节点自身所拥有的语句们
func Eval(node ast.Node, env *object.Environment) object.Object {

	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.SpeakStatement:
		return evalSpeak(node, env)
	case *ast.ListenStatement:
		return evalListen(node, env)
	case *ast.ExitStatement:
		return evalExit(node, env)
	case *ast.StepStatement:
		temp := evalProgram(&ast.Program{Statements: node.ALLStatement}, env)
		_, ok := temp.(*object.String)
		if ok {
			temp.(*object.String).Value += "Step"
			return temp
		} else {
			break
		}
	}
	var result object.Boolean
	return &result
}
func evalListen(p *ast.ListenStatement, env *object.Environment) object.Object {

	id, ok := env.Get("ID")
	if !ok {
		log.Fatal("can't find right ID")
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("can't find convert ID")
	}
	conn := Eval_Conn[idNum]

	var result object.String
	s := p.Expression.TokenLiteral()
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
	time.Sleep(time.Duration(b) * time.Second)
	fmt.Println("请输入答案")
	conn.Write([]byte("\n请输入答案\n"))
	temp := make([]byte, 20)
	length, _ := conn.Read(temp)
	ans := string(temp[:length])
	// scan := bufio.NewScanner(os.Stdin)
	// scan.Scan()
	// ans := scan.Text()
	// ans = strings.TrimSpace(ans)
	// println(ans)
	// ans = "tousu"
	fmt.Println(strings.HasPrefix(ans, "silence"))
	if strings.HasPrefix(ans, "silence") {
		ans = "ListenSilence"
		time.Sleep(time.Duration(e) * time.Second)
	} else {
		ans = "Listen" + ans
		ans = strings.ReplaceAll(ans, "\r\n", "")
		fmt.Print(ans)
	}
	result.Value = ans
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

	id, ok := env.Get("ID")
	if !ok {
		log.Fatal("can't find right ID")
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("can't find convert ID")
	}
	conn := Eval_Conn[idNum]
	result.Value = strings.ReplaceAll(result.Value, "\n", "")
	conn.Write([]byte(result.Value + "\n"))
	fmt.Println(result.Value)
	return &result
}
func evalExit(program *ast.ExitStatement, env *object.Environment) object.Object {
	var result object.String
	result.Value = "Exit"
	time.Sleep(time.Duration(4) * time.Second)
	return &result
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for i := 0; i < len(program.Statements); i++ {
		statement := program.Statements[i]
		result = Eval(statement, env)
		temp := result.Inspect()
		// var temp2 string

		if index := strings.Index(temp, "Exit"); index != -1 {
			return result
		}
		if strings.HasPrefix(temp, "Listen") {
			result.(*object.String).Value = temp[6:]
			return result
		}

		if index := strings.Index(temp, "Step"); index != -1 {
			// result.(*object.String).Value = temp[:index]
			temp := statement.(*ast.StepStatement).GetBranch(temp[:index])
			index := parser.STEP_INDEX[temp]
			i = index - 1
		}
	}
	return result
}
