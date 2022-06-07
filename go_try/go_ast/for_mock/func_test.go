package for_mock

import (
	"go_ast"
	"go_ast/func_ast"
	"go_ast/util"
	"testing"
)

var generateStruct = []*go_ast.FuncStruct{
	{
		MethodName: "SetUserCoinTryAndRetry",
		ReqParamList: []*go_ast.Param{
			{"_", "cssBizImpl"},
			{"ctx", "context.Context"},
			{"req", "*payment_model_ext.SetUserCoinTryReq"},
		},
		RespParamList: []*go_ast.Param{
			{"", "ForTestStruct"},
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
			{"", "TestStruct"},
			{"", "TestStructTwo"},
		},
	},
}

func TestConvertByteToStructWithIfStmt(t *testing.T) {
	println(util.PrintTheValue(convertByteToStructWithIfStmt(1, "policy")))
}

func TestNormalFunction(t *testing.T) {
	println(util.PrintTheValue(func_ast.GenerateFunc(normalFunction(generateStruct[0]))))
}

func TestNormalFunction2(t *testing.T) {
	println(util.PrintTheValue(func_ast.GenerateFunc(normalFunction(generateStruct[1]))))
}

func TestGet(t *testing.T) {
	Get("/Users/mengliang.yang/shopee_insurance/insurance-business-bff/src/biz/impl/s3_biz_impl.go")
}
