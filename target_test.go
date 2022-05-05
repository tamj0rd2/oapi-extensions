package lib_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestReplacement(t *testing.T) {
	assert.NoError(t, doSomething())
}

func doSomething() error {
	fset := token.NewFileSet() // positions are relative to fset
	file, err := parser.ParseFile(fset, "target.go", nil, 0)
	if err != nil {
		return err
	}

	ast.Print(fset, file)

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(file, func(n ast.Node) bool {

		// var s string
		// switch x := n.(type) {
		// case *ast.BasicLit:
		// 	s = x.Value
		// case *ast.Ident:
		// 	s = x.Name
		// }
		// if s != "" {
		// 	fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		// }
		// return true
	})

	return nil
}
