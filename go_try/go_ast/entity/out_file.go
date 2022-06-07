package entity

import "go/ast"

type OutFileFormat struct {
	Package      string
	ImportRecord []*ast.ImportSpec
	InitFunc     *ast.FuncDecl
	MainFunc     *ast.FuncDecl
	NormalFunc   []*ast.FuncDecl
	DefineType   []*ast.Spec
}

type optionForF func(out *OutFileFormat)

func WithPackage(packageN string) optionForF {
	return func(out *OutFileFormat) {
		out.Package = packageN
	}
}

func WithImportRecord(ImportRecord []*ast.ImportSpec) optionForF {
	return func(out *OutFileFormat) {
		out.ImportRecord = ImportRecord
	}
}

func WithInitFunc(InitFunc *ast.FuncDecl) optionForF {
	return func(out *OutFileFormat) {
		out.InitFunc = InitFunc
	}
}

func WithMainFunc(MainFunc *ast.FuncDecl) optionForF {
	return func(out *OutFileFormat) {
		out.MainFunc = MainFunc
	}
}

func WithNormalFunc(NormalFunc []*ast.FuncDecl) optionForF {
	return func(out *OutFileFormat) {
		out.NormalFunc = NormalFunc
	}
}

func BuildOutFileFormat(options ...optionForF) *OutFileFormat {
	ans := new(OutFileFormat)
	for _, v := range options {
		v(ans)
	}
	return ans
}

func (out *OutFileFormat) GenerateFile() *ast.File {
	return &ast.File{
		Name:    &ast.Ident{Name: out.Package},
		Decls:   out.getDeclsAll(),
		Imports: out.ImportRecord,
	}
}

func (oo *OutFileFormat) getDeclsAll() []ast.Decl {
	ans := []ast.Decl{oo.InitFunc, oo.MainFunc}
	for _, v := range oo.NormalFunc {
		ans = append(ans, v)
	}
	return ans
}
