package main

import (
	"evpeople/toyLang/ast"
	"evpeople/toyLang/evaluator"
	"evpeople/toyLang/lexer"
	"evpeople/toyLang/object"
	"evpeople/toyLang/parser"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

var port, file string

func init() {
	flag.StringVar(&port, "port", "20000", "默认采用端口是20000")
	flag.StringVar(&file, "file", "test.Toy", "默认使用的文件时test.Toy")
}
func main() {
	// os.Open("./test.Toy")
	flag.Parse()
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("open file failed,err:", err)
	}
	input := string(content)
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParserProgram()
	id := 0
	evaluator.Eval_Conn = make(map[int]net.Conn)
	listen, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		evaluator.Eval_Conn[id] = conn
		go process(conn, program, id)
		id++
	}
	// evEval2(program, env)
}

func process(conn net.Conn, program *ast.Program, id int) {
	defer conn.Close()
	// var program *ast.Program
	// program.Statements = make([]ast.Statement, len(p.Statements))
	// copy(program.Statements, p.Statements)
	//TODO:考虑深拷贝
	conn.Write([]byte("name "))
	env := object.NewEnvironment()
	env.Set("ID", strconv.Itoa(id))
	b := make([]byte, 20)
	length, err := conn.Read(b) //每次新建一个环境
	if err != nil {
		log.Fatal(err)
	}
	env.Set("name", strings.TrimSpace(string(b[:length])))
	conn.Write([]byte("amount "))
	length, err = conn.Read(b) //每次新建一个环境
	if err != nil {
		log.Fatal(err)
	}
	env.Set("amount", strings.TrimSpace(string(b[:length])))
	// Eval(program, env).(*object.String).Value
	// q, _ := env.Get("name")

	// fmt.Printf("%s", []byte(q)[:len(q)-1])
	evaluator.Eval(program, env)
}
