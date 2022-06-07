package for_mock

import (
	"go_ast/parse_file"
	"go_ast/util"
)

func Get(FilePath string) {
	fileParse := parse_file.RefFileParse()
	fileParse.InjectFile(FilePath)
	funcList := parse_file.RefFileParse().LoadFileGetFuncList()
	//for _, v := range fileParse.LoadFileGetImport() {
	//	println(util.PrintTheValue(v))
	//}
	for _, v := range GetNormalFuncList(funcList) {
		println(util.PrintTheValue(v))
	}
	//println(util.PrintTheValue(func_ast.GenerateFunc(GetInitFunc(funcList))))
	//file := entity.BuildOutFileFormat(
	//	entity.WithImportRecord(fileParse.LoadFileGetImport()),
	//	entity.WithPackage(parse_file.RefFileParse().GetPackage()),
	//	entity.WithInitFunc(func_ast.GenerateFunc(GetInitFunc(funcList))),
	//	entity.WithNormalFunc(GetNormalFuncList(funcList)),
	//)
	//util.PrintTheValue(file.InitFunc)
	//for _, v := range file.NormalFunc {
	//	util.PrintTheValue(v)
	//}
	//for _, v := range file.ImportRecord {
	//	util.PrintTheValue(v)
	//}
	//util.PrintTheValue(file.GenerateFile())
}
