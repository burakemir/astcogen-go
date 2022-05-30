// Package astcogen is a library for AST based code generation.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

type Visitor struct {
	fset *token.FileSet
}

func (v *Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch x := n.(type) {
	case *ast.CallExpr:
		sel, ok := x.Fun.(*ast.SelectorExpr)
		if ok {
			if id, ok := sel.X.(*ast.Ident); ok && id.Name == "fmt" {
				x.Args = []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: `"hello codegen world!"`,
					},
				}
				fmt.Printf("// Rewritten call to Println() at %s\n", v.fset.Position(n.Pos()))
			}
		}
	}
	return v
}

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "example/hello.go", nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	visitor := &Visitor{fset: fset}
	ast.Walk(visitor, file)

	printer.Fprint(os.Stdout, fset, file)
}
