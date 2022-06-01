package func_ast

import (
	"go/ast"
	"go_ast"
)

func getFuncList(astFile *ast.File) []*go_ast.FuncStruct {
	ans := make([]*go_ast.FuncStruct, 0, 1)
	for _, v := range astFile.Decls {
		ast.Inspect(v, func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.FuncDecl:
				funcStr := new(go_ast.FuncStruct)
				funcStr.MethodName = v.Name.Name
				funcStr.ReqParamList = getParamList(v.Type.Params)
				funcStr.RespParamList = getParamList(v.Type.Results)
				ans = append(ans, funcStr)
			}
			return true
		})
	}
	return ans
}

func getParamList(fieldList *ast.FieldList) []*go_ast.Param {
	if fieldList == nil || fieldList.List == nil {
		return nil
	}
	paramList := make([]*go_ast.Param, 0, len(fieldList.List))
	for _, reqParamList := range fieldList.List {
		typeName := getIdent(reqParamList.Type)
		if reqParamList.Names == nil {
			paramList = append(paramList, &go_ast.Param{
				ParamName: "",
				ParamType: typeName,
			})
			continue
		}
		for _, v := range reqParamList.Names {
			paramList = append(paramList, &go_ast.Param{
				ParamName: v.Name,
				ParamType: typeName,
			})
		}
	}
	return paramList
}

func getIdent(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return getIdent(t.X)
	case *ast.SelectorExpr:
		return getIdent(t.X) + "." + t.Sel.Name
	}
	return "invalid type"
}