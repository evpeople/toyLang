package ast

import "evpeople/toyLang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface { //表达式 比如 Branch "3.4"
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type StepStatement struct {
	Token token.Token
	Name  *Identifier
}

func (SS *StepStatement) statementNode()       {}
func (SS *StepStatement) TokenLiteral() string { return SS.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type SpeakStatement struct {
	Token      token.Token
	Expression Expression
}

func (SS *SpeakStatement) statementNode()       {}
func (SS *SpeakStatement) TokenLiteral() string { return SS.Token.Literal }

type SentenceStatement struct {
	Token token.Token //理论上是string
	Value string      //通过拼接 + 和$ 得出的结果
}

func (st *SentenceStatement) expressionNode() {
}
func (st *SentenceStatement) TokenLiteral() string {
	return st.Token.Literal
}

type ListenStatement struct {
}
type BranchStatement struct {
}
type SilenceStatement struct {
}
type DefaultStatement struct {
}
type ExitStatement struct {
}
