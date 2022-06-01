package func_ast

import (
	"go/ast"
	"go_ast"
	"go_ast/util"
	"testing"
)

var funcStruct = &go_ast.FuncStruct{
	MethodName: "test_method",
	ReqParamList: []*go_ast.Param{
		{"hello", "string"},
		{"World", "string"},
	},
	RespParamList: []*go_ast.Param{
		{"", "string"},
		{"", "string"},
	},
}

func TestGenerateFunc(t *testing.T) {
	println(util.PrintTheValue(GenerateFunc(funcStruct, &ast.ReturnStmt{
		Results: []ast.Expr{&ast.Ident{Name: "getTheWorld(data[0])"}, &ast.Ident{Name: "shijie"}},
	})))
}
