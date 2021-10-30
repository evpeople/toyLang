package parser

import (
	"evpeople/toyLang/ast"
	"evpeople/toyLang/lexer"
	"testing"
)

func TestStep(t *testing.T) {
	input := `
	Step welcome
	Step bad
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParserProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"welcome"},
		{"bad"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testStepStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}
func testStepStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "Step" {
		t.Errorf("want 'Step', got=%q", s.TokenLiteral())
		return false
	}
	stepStmt, ok := s.(*ast.StepStatement)
	if !ok {
		t.Errorf("want *ast.StepStatement, got=%T", s)
		return false
	}
	if stepStmt.Name.Value != name {
		t.Errorf("want %s, got=%s", name, stepStmt.Name)
		return false
	}
	return true
}
func TestSpeak(t *testing.T) {
	input := `
	Speak 'hello wo pao de zui kuai'
	Speak 'hello'+' world'
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParserProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"hello wo pao de zui kuai"},
		{"hello world"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testSpeakStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}
func testSpeakStatement(t *testing.T, s ast.Statement, sentence string) bool {
	if s.TokenLiteral() != "Speak" {
		t.Errorf("want 'Speak', got=%q", s.TokenLiteral())
		return false
	}
	speakStmt, ok := s.(*ast.SpeakStatement)
	if !ok {
		t.Errorf("want *ast.SpeakStatement, got=%T", s)
		return false
	}
	if speakStmt.Expression.TokenLiteral() != sentence {
		t.Errorf("want %s, got=%s", sentence, speakStmt.Expression.TokenLiteral())
		return false
	}
	return true
}
