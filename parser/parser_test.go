package parser

import (
	"fmt"
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
	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	literal, ok := statement.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", statement.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15}}
	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		exp, ok := statement.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("statement is not ast.PrefixExpression. got=%T", statement.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}
	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5}, {"5 - 5;", 5, "-", 5}, {"5 * 5;", 5, "*", 5}, {"5 / 5;", 5, "/", 5}, {"5 > 5;", 5, ">", 5}, {"5 < 5;", 5, "<", 5}, {"5 == 5;", 5, "==", 5}, {"5 != 5;", 5, "!=", 5},
	}
	for _, tt := range infixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		exp, ok := statement.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", statement.Expression)
		}
		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		}, {
			"!-a",
			"(!(-a))"},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)"},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}
	for _, tt := range tests {

		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		actual := program.String()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}
