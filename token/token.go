package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

//TODO: 完成变量在运行前的读取，一个env的包
//TODO: Speak 'Speak'的测试
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

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
