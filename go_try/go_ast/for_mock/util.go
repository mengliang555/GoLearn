package for_mock

import (
	"go/ast"
	"strconv"
	"strings"
)

//数组
func getArrayIndex(arrayName ast.Expr, index int) *ast.IndexExpr {
	ans := new(ast.IndexExpr)
	ans.X = arrayName
	ans.Index = &ast.Ident{
		Name: strconv.Itoa(index),
	}
	return ans
}

func convertStrToString(str string) string {
	return "\"" + str + "\""
}
func lowerForFirstCharacter(str string) string {
	if len(str) <= 0 {
		return str
	}
	builder := strings.Builder{}
	for i, v := range str {
		if i == 0 && ('A' <= v && v <= 'Z') {
			builder.WriteRune(v + 32)
		} else {
			builder.WriteRune(v)
		}
	}
	return builder.String()
}

func generateIdentList(param ...string) []*ast.Ident {
	if len(param) == 0 {
		return nil
	}
	ans := make([]*ast.Ident, 0, len(param))
	for _, v := range param {
		ans = append(ans, &ast.Ident{
			Name: v,
		})
	}
	return ans
}

func generateExprList(param ...*ast.Ident) []ast.Expr {
	if len(param) == 0 {
		return nil
	}
	ans := make([]ast.Expr, 0, len(param))
	for _, v := range param {
		ans = append(ans, v)
	}
	return ans
}
