package for_mock

import "go/ast"

const (
	errParam        = "err"
	errType         = "error"
	mockDataName    = "dataForMock"
	gUnit           = "gunit"
	GetMockDataFunc = "GetMockData"
	RegisterMockFun = "RegisterMockMethod"
	nilVal          = "nil"
	panicFunc       = "panic"
	newVal          = "new"
)

var basicMapForReturn = map[string]func(i int) ast.Expr{
	"string": func(i int) ast.Expr {
		return &ast.CallExpr{
			Fun: &ast.Ident{
				Name: "string",
			},
			Args: []ast.Expr{getArrayIndex(mockDataIdent, i)},
		}
	},
	"error": func(i int) ast.Expr {
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "errors"},
				Sel: &ast.Ident{Name: "New"},
			},
			Args: []ast.Expr{getArrayIndex(mockDataIdent, i)},
		}
	},
	"int64": func(i int) ast.Expr {
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "errors"},
				Sel: &ast.Ident{Name: "New"},
			},
			Args: []ast.Expr{getArrayIndex(mockDataIdent, i)},
		}
	},
	"int32": func(i int) ast.Expr {
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "errors"},
				Sel: &ast.Ident{Name: "New"},
			},
			Args: []ast.Expr{getArrayIndex(mockDataIdent, i)},
		}
	},
	"int": func(i int) ast.Expr {
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "errors"},
				Sel: &ast.Ident{Name: "New"},
			},
			Args: []ast.Expr{getArrayIndex(mockDataIdent, i)},
		}
	},
}
