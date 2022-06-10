// Package astcogen is a library for AST based code generation.
package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path"

	toml "github.com/pelletier/go-toml/v2"
)

// config is a struct that supplies parameters to the run of the astcogen tool.
// The config is a .toml file with the format
//   input = "<path to input file>"
//   output = "<path to output file>"
// Paths are interpreted as relative to the config file.
type Config struct {
	Input  map[string]string
	Output map[string]string
}

func (c Config) Check() error {
	if _, ok := c.Input["path"]; !ok {
		return errors.New("missing path in [input]")
	}
	if _, ok := c.Output["path"]; !ok {
		return errors.New("missing path in [output]")
	}
	return nil
}

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
	if len(os.Args) != 2 {
		log.Fatal("Need a .toml file as argument.")
	}
	configFilePath := os.Args[1]
	configDir := path.Dir(configFilePath)
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	var cfg Config
	err = toml.Unmarshal(configFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err = cfg.Check(); err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path.Join(configDir, cfg.Input["path"]), nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	visitor := &Visitor{fset: fset}
	ast.Walk(visitor, file)

	outputFile, err := os.Create(path.Join(configDir, cfg.Output["path"]))
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	printer.Fprint(outputFile, fset, file)
}
