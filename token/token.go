//Package token 包含了文法中所用的所有token类型
package token

type TokenType string

//Token 的Type用于存储类型，Literal是类似于String的Token中实际存储的内容
type Token struct {
	Type    TokenType
	Literal string
}

//所有的Token的类型
const (
	STRING  = "STRING"
	EOF     = "EOF"
	STEP    = "STEP"
	SPEAK   = "SPEAK"
	BRANCH  = "BRAHCH"
	IDENT   = "IDENT"
	EXIT    = "EXIT"
	DEFAULT = "DEFAULT"
	LISTEN  = "LISTEN"
	NUM     = "NUM"
	SILENCE = "SILENCE"
	PLUS    = "PLUS"
	ILLEGAL = "ILLEGAL"
	COMMA   = ","
	DOLLAR  = "$"
)

var keywords = map[string]TokenType{
	"Step":    STEP,
	"Speak":   SPEAK,
	"Listen":  LISTEN,
	"Branch":  BRANCH,
	"Silence": SILENCE,
	"Default": DEFAULT,
	"Exit":    EXIT,
	"+":       PLUS,
}

//LookupIdent 函数以一个文法分析读取到的字符串作为参数，然后在Token包内置的一个keyword字典中进行检索，若能检索到，则此字符串是一个关键字，不能检索到，则字符串是一个标识符
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
