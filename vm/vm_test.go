package vm

import (
	"fmt"
	"testing"
	"waiacig/ast"
	"waiacig/compiler"
	"waiacig/lexer"
	"waiacig/object"
	"waiacig/parser"
)

func parse(input string) *ast.Program {
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	return p.ParseProgram()
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)", actual, actual)
	}
	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
	}
	return nil
}

type vmTestCase struct {
	input    string
	expected interface{}
}

func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()
	for _, tt := range tests {
		program := parse(tt.input)
		c := compiler.NewCompiler()
		err := c.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}
		vm := NewVM(c.Bytecode())
		err = vm.Run()
		if err != nil {
			t.Fatalf("vm error: %s", err)
		}
		stackElem := vm.LastPoppedStackElem()
		testExpectedObject(t, tt.expected, stackElem)
	}
}

func testExpectedObject(t *testing.T,
	expected interface{},
	actual object.Object) {
	t.Helper()
	switch expected := expected.(type) {
	case int:
		err := testIntegerObject(int64(expected), actual)
		if err != nil {
			t.Errorf("testIntegerObject failed: %s", err)
		}
	}
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"1 * 2", 2},
		{"4 / 2", 2},
		{"50 / 2 * 2 + 10 - 5", 55}, {"5 * (2 + 10)", 60},
		{"5 + 5 + 5 + 5 - 10", 10}, {"2 * 2 * 2 * 2 * 2", 32}, {"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"5 * (2 + 10)", 60},
	}
	runVmTests(t, tests)
}