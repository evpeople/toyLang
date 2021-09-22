package ast

import (
	"bytes"
	"evpeople/toyLang/token"
)

type Node interface {
	TokenLiteral() string //只用于debugging and testing
	String() string
}

//Statement 的 statementNode方法主要用于帮助golang的编译器抛出Statement和Expression误用的错误
type Statement interface {
	Node
	statementNode()
}

//Expression 的 statementNode方法主要用于帮助golang的编译器抛出Statement和Expression误用的错误
type Expression interface {
	Node
	expressionNode()
}

//Program 是根节点
type Program struct {
	Statements []Statement
}

//Y 用于实现Node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//LetStatement 表示let 语句，实现state接口，说明自己是个state
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier //用于保存标识符的值
	Value Expression  //interface 类型，用于保存表达式，是这个表达式产生了Name里的value的值
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

//Identifier 表示标识符？，实现express接口，说明自己是个express，虽然在let x=10 里没有实现产生一个值，但是其他会产生，比如 let x=5*5
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

//ReturnStatement 表示Return语句
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

//Expression表达式语句
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (i *Identifier) String() string { return i.Value }

//IntegerLiteral 的Value是实际的值，但是token中存的仍然是字符串
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
