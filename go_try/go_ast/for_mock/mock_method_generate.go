package for_mock

import (
	"go/ast"
	"go/token"
	"go_ast"
)

var errParam = "err"
var mockDataName = "dataForMock"

func GenerateMockExprForInit(val []*go_ast.FuncStruct) []ast.Stmt {
	ans := make([]ast.Stmt, 0, 0)
	for _, v := range val {
		ans = append(ans, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "gunit"},
					Sel: &ast.Ident{Name: "RegisterMockMethod"},
				},
				Args:     []ast.Expr{&ast.Ident{Name: convertStrToString(v.MethodName)}, &ast.Ident{Name: "Mock_" + v.MethodName}},
				Ellipsis: 0,
				Rparen:   0,
			},
		})
	}
	return ans
}

// not support func/interface
func GenerateMockExprForMockFunc(val *go_ast.FuncStruct) []ast.Stmt {
	ans := make([]ast.Stmt, 0, 0)
	ans = append(ans, &ast.AssignStmt{
		Lhs:    []ast.Expr{&ast.Ident{Name: mockDataName}},
		TokPos: 0,
		Tok:    0,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "gunit"},
				Sel: &ast.Ident{Name: "GetMockData"},
			},
		}},
	})
	for _, v := range val.RespParamList {
		switch v.ParamType {
		case "func", "interface":
			panic("not support")
		case "struct":
			ans = append(ans, forStructIf())
		case "string":
			ans = append(ans, &ast.ReturnStmt{
				Results: []ast.Expr{&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						Sel: &ast.Ident{Name: "string"},
					},
					Args: []ast.Expr{&ast.Ident{Name: mockDataName}},
				}},
			})
		case "error":
			ans = append(ans, &ast.ReturnStmt{
				Results: []ast.Expr{&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						Sel: &ast.Ident{Name: "string"},
					},
					Args: []ast.Expr{&ast.Ident{Name: mockDataName}},
				}},
			})
		}
	}
	return ans
}

func forStructIf() *ast.IfStmt {
	val := &ast.IfStmt{
		Init: &ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: errParam}},
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "json"},
					Sel: &ast.Ident{Name: "Unmarshal"},
				},
				Args: []ast.Expr{&ast.Ident{Name: "data[0]"}, &ast.Ident{Name: "&dataForMock"}},
			}},
		},
		Cond: &ast.BinaryExpr{
			X: &ast.Ident{
				Name: errParam,
			},
			OpPos: 0,
			Op:    token.NEQ,
			Y: &ast.Ident{
				Name: "nil",
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						Sel: &ast.Ident{Name: "panic"},
					},
					Args: []ast.Expr{&ast.Ident{Name: errParam}},
				},
			}},
		},
		Else: nil,
	}
	return val
}

func convertStrToString(str string) string {
	return "\"" + str + "\""
}

//remove
func genCallExprSimple(methodName string, param ...string) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:  &ast.SelectorExpr{Sel: &ast.Ident{Name: methodName}},
		Args: generateExprList(generateIdentList(param...)...),
	}
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
