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
	Token        token.Token
	Name         *Identifier
	ALLStatement []Statement
	CaseBranch   map[string]string
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
	Token     token.Token //理论上是string
	Value     string      //通过拼接 + 和$ 得出的结果
	DollarMap map[string]string
}

func (st *SentenceStatement) expressionNode() {
}
func (st *SentenceStatement) TokenLiteral() string {
	return st.Token.Literal
}
func (st *SentenceStatement) RealTokenLiteral() string {
	// var s string
	for i := 0; i < len(st.Value); i++ {
		if st.Value[i] == '$' {
			s := st.Value[0:i]
			t, index := st.readmap(i + 1)
			st.Value = s + " " + t + st.Value[index:len(st.Value)]
			i = index
		}
	}
	return st.Value
}
func (st *SentenceStatement) readmap(index int) (string, int) {
	var s string
	i := index
	for ; i < len(st.Value) && st.Value[i] != ' '; i++ {
		s += string(st.Value[i])
	}
	trueVar, ok := st.DollarMap[s]
	if ok {
		return trueVar, i
	} else {
		println("readmap is wrong, and the key is " + s)
		return s, i
	}
}

type ListenStatement struct {
	Token      token.Token
	Expression Expression
}

func (LS *ListenStatement) statementNode()       {}
func (LS *ListenStatement) TokenLiteral() string { return LS.Token.Literal }

type ListenTime struct {
	Start string
	Last  string
}

func (lt *ListenTime) expressionNode() {
	return
}
func (lt *ListenTime) TokenLiteral() string {
	return "Start is " + lt.Start + "\nEnd is " + lt.Last
}

// func (lt *ListenTime) New(start, last string) {

// }

type BranchStatement struct {
	Token      token.Token
	Expression Expression
}

func (LS *BranchStatement) statementNode()       {}
func (LS *BranchStatement) TokenLiteral() string { return LS.Token.Literal }

type BranchCase struct {
	Case   string
	Branch string
}

func (lt *BranchCase) expressionNode() {
}
func (lt *BranchCase) TokenLiteral() string {
	return "Case is " + lt.Case + "\nBranch is " + lt.Branch
}

type SilenceStatement struct {
	Token      token.Token
	Expression Expression
}

func (LS *SilenceStatement) statementNode()       {}
func (LS *SilenceStatement) TokenLiteral() string { return LS.Token.Literal }

type SilenceBranch struct {
	Branch string
}

func (lt *SilenceBranch) expressionNode() {
}
func (lt *SilenceBranch) TokenLiteral() string {
	return lt.Branch
}

type DefaultStatement struct {
	Token      token.Token
	Expression Expression
}

func (LS *DefaultStatement) statementNode()       {}
func (LS *DefaultStatement) TokenLiteral() string { return LS.Token.Literal }

type ExitStatement struct {
	Token token.Token
}

func (LS *ExitStatement) statementNode()       {}
func (LS *ExitStatement) TokenLiteral() string { return LS.Token.Literal }
