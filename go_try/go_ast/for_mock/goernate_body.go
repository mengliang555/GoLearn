package for_mock

import (
	"go/ast"
	"go/token"
	"go_ast"
	"go_ast/func_ast"
)

func GetInitFunc(val []*go_ast.FuncStruct) *go_ast.FuncStruct {
	return &go_ast.FuncStruct{MethodName: "init", Body: generateMockExprForInitFunction(val)}
}

func generateMockExprForInitFunction(val []*go_ast.FuncStruct) []ast.Stmt {
	ans := make([]ast.Stmt, 0, 0)
	for _, v := range val {
		if len(v.MethodName) <= 0 || (v.MethodName[0] >= 'a' && v.MethodName[0] <= 'z') || (v.MethodName[0] == '_') {
			continue
		}
		ans = append(ans, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: gUnit}, Sel: &ast.Ident{Name: RegisterMockFun}},
				Args: []ast.Expr{&ast.Ident{Name: convertStrToString(v.MethodName)}, &ast.Ident{Name: "Mock_" + v.MethodName}},
			},
		})
	}
	return ans
}

func GetNormalFuncList(val []*go_ast.FuncStruct) []*ast.FuncDecl {
	ans := make([]*ast.FuncDecl, 0, len(val))
	for _, v := range val {
		if len(v.MethodName) <= 0 || (v.MethodName[0] >= 'a' && v.MethodName[0] <= 'z') || (v.MethodName[0] == '_') {
			continue
		}
		ans = append(ans, func_ast.GenerateFunc(normalFunction(v)))
	}
	return ans
}

func normalFunction(str *go_ast.FuncStruct) *go_ast.FuncStruct {
	mockData := getMockData()
	mockBody := generateTheBodyForStruct(str)
	returnExpr := generateReturnValue(str)
	str.Body = append([]ast.Stmt{mockData}, append(mockBody, returnExpr)...)
	return str
}

func generateTheBodyForStruct(val *go_ast.FuncStruct) []ast.Stmt {
	var flag = false
	responseList := make([]ast.Stmt, 0, len(val.RespParamList))
	for i, v := range val.RespParamList {
		if _, ok := basicMapForReturn[v.ParamType]; !ok {
			if !flag {
				flag = true
				responseList = append(responseList, &ast.DeclStmt{
					Decl: &ast.GenDecl{
						Tok:   token.VAR,
						Specs: []ast.Spec{&ast.ValueSpec{Names: []*ast.Ident{{Name: errParam}}, Type: &ast.Ident{Name: errType}}},
					},
				})
			}
			responseList = append(responseList, getAssignStmt(lowerForFirstCharacter(v.ParamType), []ast.Expr{&ast.CallExpr{
				Fun:  &ast.Ident{Name: newVal},
				Args: []ast.Expr{&ast.Ident{Name: v.ParamType}},
			}}))
			responseList = append(responseList, convertByteToStructWithIfStmt(i, lowerForFirstCharacter(v.ParamType)))
		}
	}
	return responseList
}

func generateReturnValue(val *go_ast.FuncStruct) *ast.ReturnStmt {
	responseList := make([]ast.Expr, 0, len(val.RespParamList))
	for i, v := range val.RespParamList {
		if val, ok := basicMapForReturn[v.ParamType]; ok {
			responseList = append(responseList, val(i))
		} else {
			responseList = append(responseList, &ast.Ident{Name: lowerForFirstCharacter(v.ParamType)})
		}
	}
	return &ast.ReturnStmt{
		Results: responseList,
	}
}

func convertByteToStructWithIfStmt(index int, paramName string) *ast.IfStmt {
	val := &ast.IfStmt{
		Init: &ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: errParam}},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "json"}, Sel: &ast.Ident{Name: "Unmarshal"}},
				Args: []ast.Expr{getArrayIndex(mockDataIdent, index), &ast.Ident{Name: "&" + paramName}},
			}},
		},
		Cond: &ast.BinaryExpr{
			X:  &ast.Ident{Name: errParam},
			Op: token.NEQ,
			Y:  &ast.Ident{Name: nilVal},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun:  &ast.Ident{Name: panicFunc},
					Args: []ast.Expr{&ast.Ident{Name: errParam}},
				},
			}},
		},
		Else: nil,
	}
	return val
}

func getMockData() ast.Stmt {
	return getAssignStmt(mockDataName, []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: gUnit}, Sel: &ast.Ident{Name: GetMockDataFunc}}}})
}
func getAssignStmt(left string, right []ast.Expr) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs:    []ast.Expr{&ast.Ident{Name: left}},
		TokPos: 0,
		Tok:    token.DEFINE,
		Rhs:    right,
	}
}
