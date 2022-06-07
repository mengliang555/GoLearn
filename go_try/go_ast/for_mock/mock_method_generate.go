package for_mock

import (
	"go/ast"
	"go/token"
	"go_ast"
)

var mockDataIdent = &ast.Ident{Name: mockDataName}
var mockDataIdentForPointer = &ast.Ident{Name: "&" + mockDataName}

func GenerateMockExprForMockFunc(val *go_ast.FuncStruct) []ast.Stmt {
	ans := make([]ast.Stmt, 0, 0)
	ans = append(ans)
	for _, v := range val.RespParamList {
		if _, ok := basicMapForReturn[v.ParamType]; !ok {
			ans = append(ans, &ast.AssignStmt{
				Lhs:    []ast.Expr{&ast.Ident{Name: lowerForFirstCharacter(v.ParamType)}},
				TokPos: 0,
				Tok:    token.DEFINE,
				Rhs:    []ast.Expr{&ast.CallExpr{Fun: &ast.Ident{Name: "new"}, Args: []ast.Expr{&ast.Ident{Name: v.ParamType}}}},
			})
		}
	}
	ans = append(ans, generateReturnValue(val))
	return ans
}
