package main

import (
	"evpeople/toyLang/ast"
	"evpeople/toyLang/db"
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
	// flag.ContinueOnError
	flag.StringVar(&port, "port", "20000", "默认采用端口是20000")
	flag.StringVar(&file, "file", "test.Toy", "默认使用的文件时test.Toy")
}
func main() {
	flag.Parse()
	if flag.NFlag() != 2 {
		fmt.Println("程序将以test.Toy作为输入，在20000端口启动")
	}

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
}

func process(conn net.Conn, program *ast.Program, id int) {
	defer conn.Close()
	env := object.NewEnvironment()
	env.Set("ID", strconv.Itoa(id))
	if id > len(db.Peoples)-1 {
		conn.Write([]byte("初始化，姓名："))
		b := make([]byte, 20)
		length, err := conn.Read(b) //每次新建一个环境
		if err != nil {
			log.Fatal(err)
		}
		env.Set("name", strings.TrimSpace(string(b[:length])))
		conn.Write([]byte("账单余额"))
		length, err = conn.Read(b) //每次新建一个环境
		if err != nil {
			log.Fatal(err)
		}
		env.Set("amount", strings.TrimSpace(string(b[:length])))
	} else {
		env.Set("name", db.Peoples[id].Name)
		env.Set("amount", db.Peoples[id].Amount)
	}

	evaluator.Eval(program, env)
}
