package evaluator

import (
	"waiig/ast"
	"waiig/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		switch node.Operator {
		case "!":
			return evalBangOperatorExpression(right)
		case "-":
			return evalMinusPrefixOperatorExpression(right)
		default:
			return NULL
		}
	case *ast.Boolean:
		if node.Value {
			return TRUE
		}
		return FALSE
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
	}
	return result
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}
