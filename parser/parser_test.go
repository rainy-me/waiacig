package parser

import (
	"testing"
	"waiig/ast"
	"waiig/lexer"
)

// func TestLetStatement(t *testing.T) {
// 	input := `
// 	let x = 5;
// 	let y = 10;
// 	let foobar = 838383;
// 	`
// 	l := lexer.NewLexer(input)
// 	p := NewParser(l)
// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)
// 	if program == nil {
// 		t.Fatalf("ParseProgram() returned nil")
// 	}
// 	if len(program.Statements) != 3 {
// 		t.Fatalf("program.Statements does not contains 3 statement. got=%d",
// 			len(program.Statements))
// 	}
// 	tests := []struct {
// 		expectedIdentifier string
// 	}{
// 		{"x"},
// 		{"y"},
// 		{"foobar"},
// 	}
// 	for i, tt := range tests {
// 		statement := program.Statements[i]
// 		if !testLetStatement(t, statement, tt.expectedIdentifier) {
// 			return
// 		}
// 	}

// }

// func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
// 	if s.TokenLiteral() != "let" {
// 		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
// 		return false
// 	}
// 	statement, ok := s.(*ast.LetStatement)
// 	if !ok {
// 		t.Errorf("s not *ast.LetStatement. got=%T", s)
// 		return false
// 	}
// 	if statement.Name.Value != name {
// 		t.Errorf("statement.Name.Value not '%s'. got=%s", name, statement.Name.Value)
// 		return false
// 	}
// 	if statement.Name.TokenLiteral() != name {
// 		t.Errorf("s.Name not '%s'. got=%s", name, statement.Name)
// 		return false
// 	}
// 	return true
// }

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parder has %d errors", len(errors))
	for i, msg := range errors {
		t.Errorf("%d) parser error: %q", i+1, msg)
	}
	t.FailNow()
}

// func TestReturnStatement(t *testing.T) {
// 	input := `
// return 5; return 10; return 993322; `
// 	l := lexer.NewLexer(input)
// 	p := NewParser(l)
// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)
// 	if len(program.Statements) != 3 {
// 		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
// 			len(program.Statements))
// 	}

// 	for i, statement := range program.Statements {
// 		returnStatement, ok := statement.(*ast.ReturnStatement)
// 		if !ok {
// 			t.Errorf("%d) statement not *ast.returnStatement. got=%T", i+1, statement)
// 			continue
// 		}
// 		if returnStatement.TokenLiteral() != "return" {
// 			t.Errorf("returnStatement.TokenLiteral not 'return', got %q", returnStatement.TokenLiteral())
// 		}
// 	}
// }

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	ident, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", statement.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}
}
