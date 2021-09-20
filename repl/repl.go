package repl

import (
	"bufio"
	"evpeople/toyLang/lexer"
	"evpeople/toyLang/token"
	"fmt"
	"io"
)

const PROMPT = ">>"

//Start 函数的out似乎是个伏笔
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf("%s", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
