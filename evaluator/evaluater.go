//Evaluate是实际解释执行的代码所在的包
package evaluator

import (
	"errors"
	"evpeople/toyLang/ast"
	"evpeople/toyLang/object"
	"evpeople/toyLang/parser"
	"fmt"
	"log"
	"net"
	"os"
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
			//包装返回值，用于给evalProgram判别
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

	var result object.String
	//解析什么时候开始听，听多久
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

	ans := sendMessageWithTimeOut("\n请输入答案\n", env, e)
	//对听到的结果进行包装，用于evalProgram
	if strings.HasPrefix(ans, "silence") {
		ans = "ListenSilence"
	} else {
		ans = "Listen" + ans
		ans = strings.ReplaceAll(ans, "\r\n", "")
	}
	result.Value = ans
	return &result
}

func evalSpeak(program *ast.SpeakStatement, env *object.Environment) object.Object {
	var result object.String
	//判断在语句中是否有$
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
	result.Value = sendMessage(result.Value, env)
	return &result
}
func evalExit(program *ast.ExitStatement, env *object.Environment) object.Object {
	var result object.String
	result.Value = "Exit"
	time.Sleep(time.Duration(1) * time.Second)
	sendMessage("欢迎下次使用", env)
	return &result
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object
	for i := 0; i < len(program.Statements); i++ {
		//遍历一个Step的句子们
		statement := program.Statements[i]
		result = Eval(statement, env)
		temp := result.Inspect()

		if index := strings.Index(temp, "Exit"); index != -1 {
			return result
		}
		if strings.HasPrefix(temp, "Listen") {
			result.(*object.String).Value = temp[6:]
			fmt.Println(result.Inspect())
			return result
		}

		if index := strings.Index(temp, "Step"); index != -1 {
			//根据Step解除包装后的结果，判断将要跳转到的语句
			temp := statement.(*ast.StepStatement).GetBranch(temp[:index])
			index := parser.STEP_INDEX[temp]
			//重置i，进行跳转
			i = index - 1
		}
	}
	return result
}

//发送消息
func sendMessage(s string, env *object.Environment) string {
	id, ok := env.Get("ID")
	if !ok {
		log.Fatal("can't find right ID")
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("can't find convert ID")
	}
	conn := Eval_Conn[idNum]
	s = strings.ReplaceAll(s, "\n", "")
	conn.Write([]byte(s + "\n"))
	fmt.Println(s)
	return s
}

//发送消息并设置等待时长
func sendMessageWithTimeOut(s string, env *object.Environment, e int) string {
	id, ok := env.Get("ID")
	if !ok {
		log.Fatal("can't find right ID")
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("can't find convert ID")
	}
	conn := Eval_Conn[idNum]
	conn.SetReadDeadline(time.Now().Add(time.Duration(e) * time.Second))
	temp := make([]byte, 20)
	length, err := conn.Read(temp)
	var ans string
	if errors.Is(err, os.ErrDeadlineExceeded) {
		ans = "silenceS"
	} else {
		ans = string(temp[:length])
	}
	conn.SetReadDeadline(time.Time{})
	return ans
}
