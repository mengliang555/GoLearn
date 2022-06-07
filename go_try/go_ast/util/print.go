package util

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
)

func PrintTheValue(node ast.Node) string {
	if node == nil {
		return ""
	}
	byteStream := bytes.NewBufferString("")
	err := format.Node(byteStream, token.NewFileSet(), node)
	if err != nil {
		return ""
	}
	return byteStream.String()
}
