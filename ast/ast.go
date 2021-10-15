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
