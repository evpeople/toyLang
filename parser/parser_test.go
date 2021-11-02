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
	checkError(p)
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

func TestListen(t *testing.T) {
	input := `
	Listen 5, 20
	Listen 6,20
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParserProgram()
	checkError(p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		begin    string
		lastTime string
	}{
		{"5", "20"},
		{"6", "20"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		sentence := "Start is " + tt.begin + "\nEnd is " + tt.lastTime
		if !testListenStatement(t, stmt, sentence) {
			return
		}
	}
}
func testListenStatement(t *testing.T, s ast.Statement, sentence string) bool {
	if s.TokenLiteral() != "Listen" {
		t.Errorf("want 'Listen', got=%q", s.TokenLiteral())
		return false
	}
	listenStmt, ok := s.(*ast.ListenStatement)
	if !ok {
		t.Errorf("want *ast.ListenStatement, got=%T", s)
		return false
	}
	if listenStmt.Expression.TokenLiteral() != sentence {
		t.Errorf("want %s, got=%s", sentence, listenStmt.Expression.TokenLiteral())
		return false
	}
	return true
}

func TestBranch(t *testing.T) {
	input := `
	Branch 'happy',Proc
	Branch 'bad',omp
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParserProgram()
	checkError(p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		Case   string
		Branch string
	}{
		{"happy", "Proc"},
		{"bad", "omp"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		sentence := "Case is " + tt.Case + "\nBranch is " + tt.Branch
		if !testBranchStatement(t, stmt, sentence) {
			return
		}
	}
}
func testBranchStatement(t *testing.T, s ast.Statement, sentence string) bool {
	if s.TokenLiteral() != "Branch" {
		t.Errorf("want 'Branch', got=%q", s.TokenLiteral())
		return false
	}
	BranchStmt, ok := s.(*ast.BranchStatement)
	if !ok {
		t.Errorf("want *ast.ListenStatement, got=%T", s)
		return false
	}
	if BranchStmt.Expression.TokenLiteral() != sentence {
		t.Errorf("want %s, got=%s", sentence, BranchStmt.Expression.TokenLiteral())
		return false
	}
	return true
}

func TestSilence(t *testing.T) {
	input := `
	Silence  silence
	Silence  s
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParserProgram()
	checkError(p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		Branch string
	}{
		{"silence"},
		{"s"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		sentence := "Branch is " + tt.Branch
		if !testSilenceStatement(t, stmt, sentence) {
			return
		}
	}
}
func testSilenceStatement(t *testing.T, s ast.Statement, sentence string) bool {
	if s.TokenLiteral() != "Silence" {
		t.Errorf("want 'Branch', got=%q", s.TokenLiteral())
		return false
	}
	BranchStmt, ok := s.(*ast.SilenceStatement)
	if !ok {
		t.Errorf("want *ast.ListenStatement, got=%T", s)
		return false
	}
	if BranchStmt.Expression.TokenLiteral() != sentence {
		t.Errorf("want %s, got=%s", sentence, BranchStmt.Expression.TokenLiteral())
		return false
	}
	return true
}
