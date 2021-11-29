//Lexer 用于词法分析
package lexer

import "evpeople/toyLang/token"

//Lexer因为需要提供给Parser使用，所以是可导出的类型，但内部值都是不可导出的
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char) 当前位置
	readPosition int  // current reading position in input (after current char)  下一个位置
	ch           byte // current char under examination
}

//通过输入的DSL程序，形成一个Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() //初始化l 的其他变量
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //ascii null,读到了终点，使用byte是因为都是英文
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

//提供给Parser获取下一个Token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace() //只是读到之后就丢弃，不是预先除去所有的whitespace
	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '$':
		tok = newToken(token.DOLLAR, l.ch)
	case '\'':
		l.readChar()
		tok.Literal = l.readString()
		tok.Type = token.STRING
		//单目符号如上
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	//EOF如此
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) //这里循环读取，得到单词
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.NUM
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch) //非法字符
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func (l *Lexer) readString() string {
	position := l.position
	for isLetter(l.ch) || l.ch == ' ' || l.ch == ',' {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

//与readchar相同，但是不移动当前指针的位置
// func (l *Lexer) peekChar() byte {
// 	if l.readPosition >= len(l.input) {
// 		return 0
// 	} else {
// 		return l.input[l.readPosition]
// 	}
// }
