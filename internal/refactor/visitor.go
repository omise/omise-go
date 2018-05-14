package main

import "go/ast"

type visitor struct {
	err error
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if ident, ok := node.(*ast.Ident); ok {
		if ident.Name == "e" {
			ident.Name = "err"
		}
	}

	return v
}
