package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

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
