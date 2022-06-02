package go_try

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
)

// 解析String为一个表达式，用于后续的规则判断

func ParseStrForBinaryExpr(str string) bool {
	parseFile, err := parser.ParseExpr(str)
	panicErr(err)
	rootVal, ok := parseFile.(*ast.BinaryExpr)
	if !ok {
		panic("str is invalid BinaryExpr")
	}
	return parseExprForBinaryExpr(rootVal)
}

func parseExprForBinaryExpr(rootVal *ast.BinaryExpr) bool {
	left := parseExprForBinaryExpr(rootVal.X.(*ast.BinaryExpr))
	right := parseExprForBinaryExpr(rootVal.Y.(*ast.BinaryExpr))
	return left || right
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintStrExpr(str string) {
	parseFile, err := parser.ParseExpr(str)
	panicErr(err)
	ast.Print(token.NewFileSet(), parseFile)
}

var allowedOperator = map[token.Token]struct{}{token.EQL: {}, token.NEQ: {}, token.LSS: {}, token.GTR: {}, token.LEQ: {}, token.GEQ: {}}

func OperateCompare(left, right any, operator token.Token) (bool, error) {
	if _, ok := allowedOperator[operator]; !ok {
		return false, errors.New("not allowed operator")
	}
	switch operator {
	case token.EQL, token.NEQ:
		return left == right, nil
	case token.LSS, token.GTR, token.LEQ, token.GEQ:
		return compareForNumber(left, right, operator), nil
	}
	return false, nil
}

func compareForNumber[T Number](left, right any, operator token.Token) bool {
	switch operator {
	case token.LSS:
		return left.(T) < right.(T)
	case token.GTR:
		return left.(T) > right.(T)
	case token.LEQ:
		return left.(T) <= right.(T)
	case token.GEQ:
		return left.(T) >= right.(T)
	default:
		return false
	}
}
