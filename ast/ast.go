//Ast提供了抽象语法树节点的鸭子类型的定义
package ast

import (
	"evpeople/toyLang/token"
)

//Node是所有的节点共有的类型
type Node interface {
	TokenLiteral() string
}

//Statement是可执行的语句，如Speak和Listen
type Statement interface {
	Node
	statementNode()
}

//Expression是表意的语句，不能执行，但是说明了其他语句怎么执行
type Expression interface { //表达式 比如 Branch "3.4"
	Node
	expressionNode()
}

//Program是DSL生成的程序实际构造出的对象
type Program struct {
	Statements []Statement
}

//TokenLiteral用于实现 Node类型，同时返回首语句的类型
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//StepStatement是每一个Step所构造出的对象，在AllStatement里存放所有的语句，在CaseBranch里存储不同的输入所对应跳转对象
type StepStatement struct {
	Token        token.Token
	Name         *Identifier
	ALLStatement []Statement
	CaseBranch   map[string]string
}

//statementNode，用于实现 statement
func (SS *StepStatement) statementNode() {}

//用于实现Node，并返回当前节点内部的属性
func (SS *StepStatement) TokenLiteral() string { return SS.Token.Literal }

//用于获取输入对应的Branch，并在不能匹配成功的时候返回Default
func (SS *StepStatement) GetBranch(a string) string {
	temp, ok := SS.CaseBranch[a]
	if ok {
		return temp
	} else {
		return SS.CaseBranch["Default"]
	}
}

//Identifier 标识符表达式对应的语法节点
type Identifier struct {
	Token token.Token
	Value string
}

//用于实现expression
func (i *Identifier) expressionNode() {}

//用于实现Node，并返回当前节点内部的属性
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

//Speak 语句所对应的语法节点
type SpeakStatement struct {
	Token      token.Token
	Expression Expression
}

//用于实现statement
func (SS *SpeakStatement) statementNode() {}

//用于实现Node，并返回当前节点内部的属性
func (SS *SpeakStatement) TokenLiteral() string { return SS.Token.Literal }

//Sentence在生成的过程，保存了一个DollarMap，用于在不同的用户执行的时候，从不同的environment中取值，具体流程是，发现一个语句有不为空的DollarMap时，遍历DollarMap，并在Environment中取相应的键，填充DollarMap的值，最后通过RealLiteral返回结果
type Sentence struct {
	Token     token.Token
	Value     string            //通过拼接 + 和$ 得出的结果
	DollarMap map[string]string //用于存储语法节点中，$出现的位置和对应的表达式，并在执行的时候给对应的表达式通过环境变量赋值。
}

//用于实现express
func (st *Sentence) expressionNode() {
}

//用于实现Node，并返回当前节点内部的属性
func (st *Sentence) TokenLiteral() string {
	return st.Token.Literal
}

//由于在运行时返回从 环境变量中已经提取了信息的最终构成的Sentence
func (st Sentence) RealTokenLiteral() string {
	// var s string
	for i := 0; i < len(st.Value); i++ {
		if st.Value[i] == '$' {
			s := st.Value[0:i]
			// s = strings.TrimSpace(s)
			t, index := st.readMap(i + 1)
			st.Value = s + " " + t + st.Value[index:len(st.Value)]
			i = index
		}
	}
	return st.Value
}

//用于从DollarMap中读值，当值不存在的时候报错
func (st Sentence) readMap(index int) (string, int) {
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

//Listen构造的语法节点
type ListenStatement struct {
	Token      token.Token
	Expression Expression
}

//用于实现statement
func (LS *ListenStatement) statementNode() {}

//用于实现Node，并返回当前节点内部的属性
func (LS *ListenStatement) TokenLiteral() string { return LS.Token.Literal }

//用于记录什么时候开始听，听多长时间的表达式
type ListenTime struct {
	Start string
	Last  string
}

//用于实现expression
func (lt *ListenTime) expressionNode() {}

//用于实现Node，并返回当前节点内部的属性
func (lt *ListenTime) TokenLiteral() string {
	return "Start is " + lt.Start + "\nEnd is " + lt.Last
}

// func (lt *ListenTime) New(start, last string) {

// }

//Branch语句构造语法节点
type BranchStatement struct {
	Token      token.Token
	Expression Expression
}

//用于实现statement
func (LS *BranchStatement) statementNode() {}

//用于实现Node，并返回当前节点内部的属性
func (LS *BranchStatement) TokenLiteral() string { return LS.Token.Literal }

//用于记录Branch语句的跳转条件已经对应的跳转结果的expression
type BranchCase struct {
	Case   string
	Branch string
}

//用于实现expression
func (lt *BranchCase) expressionNode() {
}

//用于实现Node，并返回当前节点内部的属性
func (lt *BranchCase) TokenLiteral() string {
	return "Case is " + lt.Case + "\nBranch is " + lt.Branch
}

//Silence构造的语句的语法节点
type SilenceStatement struct {
	Token      token.Token
	Expression Expression
}

//用于实现statement
func (LS *SilenceStatement) statementNode() {}

//用于实现Node，并返回当前节点内部的属性
func (LS *SilenceStatement) TokenLiteral() string { return LS.Token.Literal }

//SilenceBranch所对应的表达式的语法节点
type SilenceBranch struct {
	Branch string
}

//用于实现表达式
func (lt *SilenceBranch) expressionNode() {
}

//用于实现Node，并返回当前节点内部的属性
func (lt *SilenceBranch) TokenLiteral() string {
	return lt.Branch
}

//Default语句对应的语法节点
type DefaultStatement struct {
	Token      token.Token
	Expression Expression
}

//用于实现statement
func (LS *DefaultStatement) statementNode() {}

//用于实现表达式
func (LS *DefaultStatement) TokenLiteral() string { return LS.Token.Literal }

//Exit语句对应的语法节点
type ExitStatement struct {
	Token token.Token
}

//用于实现statement
func (LS *ExitStatement) statementNode() {}

//用于实现表达式
func (LS *ExitStatement) TokenLiteral() string { return LS.Token.Literal }
