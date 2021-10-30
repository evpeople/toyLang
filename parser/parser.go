package parser

import (
	"evpeople/toyLang/ast"
	"evpeople/toyLang/lexer"
	"evpeople/toyLang/token"
	"fmt"
)

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
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.STEP:
		return p.parseStepStatement()
	case token.SPEAK:
		return p.parseSpeakStatement()
	default:
		return nil
	}
}
func (p *Parser) parseSpeakStatement() *ast.SpeakStatement {
	stmt := &ast.SpeakStatement{Token: p.curToken}
	stmt.Expression = p.parseSentence()
	return stmt
}
func (p *Parser) parseSentence() *ast.SentenceStatement {
	var st string
	for p.peekTokenIs(token.STRING) || p.peekTokenIs(token.PLUS) {
		switch tk := p.peekToken; tk.Type {
		case token.STRING:
			st += tk.Literal
			p.nextToken()
		case token.PLUS:
			p.nextToken()
		}

		//TODO: 先处理没有dollar的情况，然后再处理有dollar的情况。
		//TODO: 没有dollar的情况下，就是switch (string) (plus) 然后得出结果，返回一个ast
		//TODO: 不算单纯的parse，但是可能算是优化过了
		//TODO：还是相当于对每次连接做一个新的AST了，
		//TODO: 把这段结合在一起，原本的也是应该算在eval中，所以抄一下书的String的前部分。

	}
	stmt := &ast.SentenceStatement{Token: token.Token{Type: token.STRING, Literal: st}, Value: st}
	return stmt
}
func (p *Parser) parseStepStatement() *ast.StepStatement {
	stmt := &ast.StepStatement{Token: p.curToken}
	p.nextToken()
	name := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	stmt.Name = name
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
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
