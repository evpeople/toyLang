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
	default:
		return nil
	}
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
func (p *Parser) expectPeek(t token.TokenType) bool { //通过检测下一个token的类型来保证整体的正确性，这时候不易debug，因为返回了个nil
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
