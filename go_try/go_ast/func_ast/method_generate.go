package func_ast

import (
	"go/ast"
	"go_ast"
)

// 无复杂逻辑
func GenerateFunc(funcStruct *go_ast.FuncStruct, exprList ...ast.Stmt) *ast.FuncDecl {
	ans := new(ast.FuncDecl)
	ans.Name = generateIdent(funcStruct.MethodName)
	ans.Type = generateFuncType(funcStruct)
	ans.Body = &ast.BlockStmt{
		List: append(funcStruct.Body, exprList...),
	}
	return ans
}

func generateIdent(name string) *ast.Ident {
	return &ast.Ident{
		NamePos: 0,
		Name:    name,
		Obj:     nil,
	}
}

func generateFuncType(funcStruct *go_ast.FuncStruct) *ast.FuncType {
	val := new(ast.FuncType)
	val.Params = genFieldList(funcStruct.ReqParamList)
	val.Results = genFieldList(funcStruct.RespParamList)
	return val
}

func genFieldList(paramList []*go_ast.Param) *ast.FieldList {
	val := new(ast.FieldList)
	val.List = make([]*ast.Field, 0, len(paramList))
	for _, v := range paramList {
		val.List = append(val.List, &ast.Field{
			Names: []*ast.Ident{generateIdent(v.ParamName)},
			Type:  generateIdent(v.ParamType),
		})
	}
	return val
}
