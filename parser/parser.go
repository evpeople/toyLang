package parser

import (
	"evpeople/toyLang/ast"
	"evpeople/toyLang/lexer"
	"evpeople/toyLang/token"
	"fmt"
)

var STEPINDEX map[string]int

// var DEFAULT string

type Parser struct {
	l         *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	i := 0
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement(i)
		i++
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		// p.nextToken()
	}
	return program
}
func (p *Parser) parseStatement(index int) ast.Statement {
	switch p.curToken.Type {
	case token.STEP:
		a := p.parseStepStatement()
		if STEPINDEX == nil {
			STEPINDEX = make(map[string]int)
		}
		STEPINDEX[a.Name.TokenLiteral()] = index
		return a
	case token.SPEAK:
		return p.parseSpeakStatement()
	case token.LISTEN:
		return p.parseListenStatement()
	case token.BRANCH:
		return p.parseBranchStatement()
	case token.SILENCE:
		return p.parseSilenceStatement()
	case token.DEFAULT:
		return p.parseSilenceStatement()
	case token.EXIT:
		return p.parseExitStatement()
	default:
		return nil
	}
}
func (p *Parser) parseExitStatement() *ast.ExitStatement {
	stmt := &ast.ExitStatement{Token: p.curToken}
	// p.nextToken()
	return stmt
}
func (p *Parser) parseSilenceStatement() *ast.SilenceStatement {

	stmt := &ast.SilenceStatement{Token: p.curToken}
	stmt.Expression = p.parseSilence()
	return stmt
}
func (p *Parser) parseDefaultStatement() *ast.DefaultStatement {

	stmt := &ast.DefaultStatement{Token: p.curToken}
	stmt.Expression = p.parseSilence()
	return stmt
}
func (p *Parser) parseSilence() *ast.SilenceBranch {
	stmt := &ast.SilenceBranch{}
	if p.expectPeek(token.IDENT) {
		stmt.Branch = p.peekToken.Literal
		p.nextToken()
	}
	return stmt
}
func (p *Parser) parseBranchStatement() *ast.BranchStatement {
	stmt := &ast.BranchStatement{Token: p.curToken}
	stmt.Expression = p.parseBranchCase()
	return stmt
}

func (p *Parser) parseBranchCase() *ast.BranchCase {
	stmt := &ast.BranchCase{}
	if p.expectPeek(token.STRING) {
		stmt.Case = p.peekToken.Literal
		p.nextToken()
	}
	if p.expectPeek(token.COMMA) {
		p.nextToken()
	}
	if p.expectPeek(token.IDENT) {
		stmt.Branch = p.peekToken.Literal
		p.nextToken()
	}
	return stmt
}
func (p *Parser) parseListenStatement() *ast.ListenStatement {
	stmt := &ast.ListenStatement{Token: p.curToken}
	stmt.Expression = p.parseListenTime()
	p.errors = nil
	//TODO:处理error重复报出的问题
	return stmt
}
func (p *Parser) parseSpeakStatement() *ast.SpeakStatement {
	stmt := &ast.SpeakStatement{Token: p.curToken}
	stmt.Expression = p.parseSentence()
	p.errors = nil
	//TODO:处理error重复报出的问题
	return stmt
}
func (p *Parser) parseListenTime() *ast.ListenTime {
	time := make([]string, 0)
	p.nextToken()
	for p.peekToken.Type == token.NUM || p.peekToken.Type == token.COMMA {
		if p.curToken.Type == token.COMMA {
			p.nextToken()
			continue
		}
		time = append(time, p.curToken.Literal)
		p.nextToken()
	}
	time = append(time, p.curToken.Literal)
	stmt := &ast.ListenTime{Start: time[0], Last: time[1]}
	return stmt
}
func (p *Parser) parseSentence() *ast.SentenceStatement {
	var st string
	var dollarMap map[string]string
	for p.expectPeek(token.STRING) || p.expectPeek(token.PLUS) || p.expectPeek(token.DOLLAR) {
		switch tk := p.peekToken; tk.Type {
		case token.STRING:
			st += tk.Literal
			p.nextToken()
		case token.PLUS:
			p.nextToken()
		case token.DOLLAR:
			if dollarMap == nil {
				dollarMap = make(map[string]string)
			}
			p.nextToken()
			st += "$" + p.peekToken.Literal
			p.nextToken()
			dollarMap[p.curToken.Literal] = ""
		}
		//TODO: 先处理没有dollar的情况，然后再处理有dollar的情况。
		//TODO: 没有dollar的情况下，就是switch (string) (plus) 然后得出结果，返回一个ast
		//TODO: 不算单纯的parse，但是可能算是优化过了
		//TODO：还是相当于对每次连接做一个新的AST了，
		//TODO: 把这段结合在一起，原本的也是应该算在eval中，所以抄一下书的String的前部分。

	}
	stmt := &ast.SentenceStatement{Token: token.Token{Type: token.STRING, Literal: st}, Value: st, DollarMap: dollarMap}
	return stmt
}
func (p *Parser) parseStepStatement() *ast.StepStatement {
	stmt := &ast.StepStatement{Token: p.curToken}
	p.nextToken()
	name := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	stmt.Name = name
	i := 0
	//TODO:需要用这个i,算出silence可能。
	for !p.curTokenIs(token.STEP) && !p.curTokenIs(token.EOF) {
		ostmt := p.parseStatement(-1)

		if ostmt != nil {
			stmt.ALLStatement = append(stmt.ALLStatement, ostmt)
			if stmt.CaseBranch == nil {
				stmt.CaseBranch = make(map[string]string)
			}
			if s, ok := ostmt.(*ast.SilenceStatement); ok {
				// if s.TokenLiteral() == "Silence" {
				stmt.CaseBranch[s.TokenLiteral()] = s.Expression.TokenLiteral()
				// }
			}
			if t, ok := ostmt.(*ast.BranchStatement); ok {
				stmt.CaseBranch[t.Expression.(*ast.BranchCase).Case] = t.Expression.(*ast.BranchCase).Branch
			}
		}

		i++
		p.nextToken()
	}
	return stmt
}
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
func (p *Parser) Errors() []string {
	return p.errors
}
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		return true
	} else {
		p.peekError(t)
		return false
	}
}
func checkError(p *Parser) {
	if len(p.errors) != 0 {
		for _, v := range p.errors {
			fmt.Println(v)
		}
	}
}
