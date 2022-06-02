package func_ast

import (
	"go/ast"
	"go_ast"
	"go_ast/for_mock"
	"go_ast/util"
	"testing"
)

var funcStructTest = &go_ast.FuncStruct{
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

var generateStruct = []*go_ast.FuncStruct{
	{
		MethodName: "SetUserCoinTryAndRetry",
		ReqParamList: []*go_ast.Param{
			{"_", "cssBizImpl"},
			{"ctx", "context.Context"},
			{"req", "*payment_model_ext.SetUserCoinTryReq"},
		},
		RespParamList: []*go_ast.Param{
			{"", "string"},
			{"", "error"},
		},
	},
	{
		MethodName: "SetUserCoinConfirm",
		ReqParamList: []*go_ast.Param{
			{"_", "cssBizImpl"},
			{"ctx", "context.Context"},
			{"transactionId", "string"},
		},
		RespParamList: []*go_ast.Param{
			{"", "error"},
		},
	},
}

var mockGenerateForInit = []*go_ast.FuncStruct{
	{
		MethodName:    "init",
		ReqParamList:  nil,
		RespParamList: nil,
		Body:          for_mock.GenerateMockExprForInit(generateStruct),
	},
}

func TestGenerateFunc(t *testing.T) {
	println(util.PrintTheValue(GenerateFunc(funcStructTest, &ast.ReturnStmt{
		Results: []ast.Expr{&ast.Ident{Name: "getTheWorld(data[0])"}, &ast.Ident{Name: "shijie"}},
	})))
}

func TestGenerateMockFuncInit(t *testing.T) {
	println(util.PrintTheValue(GenerateFunc(mockGenerateForInit[0])))
}

func TestGenerateMockFuncMethod(t *testing.T) {
	generateStruct[0].Body = for_mock.GenerateMockExprForMockFunc(generateStruct[0])
	println(util.PrintTheValue(GenerateFunc(generateStruct[0])))
}
