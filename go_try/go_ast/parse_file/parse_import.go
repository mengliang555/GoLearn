package parse_file

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go_ast"
)

type FileParse interface {
	LoadFileGetFuncList() []*go_ast.FuncStruct
	GetPackage() string
	LoadFileGetImport() []*ast.ImportSpec
	InjectFile(filePath string)
}

var _FileParse FileParse

func RefFileParse() FileParse {
	return _FileParse
}

type FileParseImpl struct {
	astFile *ast.File
}

func (f *FileParseImpl) GetPackage() string {
	return f.astFile.Name.Name
}

func (f *FileParseImpl) InjectFile(filePath string) {
	astFile, err := parser.ParseFile(token.NewFileSet(), filePath, nil, 0)
	if err != nil {
		panic(err)
	}
	f.astFile = astFile
}

func (f *FileParseImpl) LoadFileGetFuncList() []*go_ast.FuncStruct {
	ans := make([]*go_ast.FuncStruct, 0, 0)
	f.traverseAllNode(func(node ast.Node) {
		switch v := node.(type) {
		case *ast.FuncDecl:
			funcStr := new(go_ast.FuncStruct)
			funcStr.MethodName = v.Name.Name
			funcStr.ReqParamList = getParamList(v.Type.Params)
			funcStr.RespParamList = getParamList(v.Type.Results)
			ans = append(ans, funcStr)
		}
	})
	return ans
}

func (f *FileParseImpl) LoadFileGetImport() []*ast.ImportSpec {
	ans := make([]*ast.ImportSpec, 0, 0)
	f.traverseAllNode(func(node ast.Node) {
		switch v := node.(type) {
		case *ast.ImportSpec:
			ans = append(ans, v)
		}
	})
	return ans
}

func init() {
	_FileParse = new(FileParseImpl)
}

func (f *FileParseImpl) traverseAllNode(deal func(node ast.Node)) {
	ast.Inspect(f.astFile, func(node ast.Node) bool {
		deal(node)
		return true
	})
}

func getParamList(fieldList *ast.FieldList) []*go_ast.Param {
	if fieldList == nil || fieldList.List == nil {
		return nil
	}
	paramList := make([]*go_ast.Param, 0, len(fieldList.List))
	for _, reqParamList := range fieldList.List {
		typeName := getIdent(reqParamList.Type)
		if reqParamList.Names == nil {
			paramList = append(paramList, &go_ast.Param{ParamType: typeName})
			continue
		}
		for _, v := range reqParamList.Names {
			paramList = append(paramList, &go_ast.Param{ParamName: v.Name, ParamType: typeName})
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
