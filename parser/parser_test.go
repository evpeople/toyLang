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
