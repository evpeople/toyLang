package lexer

import (
	"evpeople/toyLang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	Step welcome
	Speak $name + 'w哈哈能藏wang哇e'
	Listen 5,20
	Branch 'not good', cProc
	Exit
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.STEP, "Step"},
		{token.IDENT, "welcome"},
		{token.SPEAK, "Speak"},
		{token.DOLLAR, "$"},
		{token.IDENT, "name"},
		{token.PLUS, "+"},
		{token.STRING, "w哈哈能藏wang哇e"},
		{token.LISTEN, "Listen"},
		{token.NUM, "5"},
		{token.COMMA, ","},
		{token.NUM, "20"},
		{token.BRANCH, "Branch"},
		{token.STRING, "not good"},
		{token.COMMA, ","},
		{token.IDENT, "cProc"},
		{token.EXIT, "Exit"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
