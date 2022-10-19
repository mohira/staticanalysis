package manyargs

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "manyargs",
	Doc:      "引数が多い関数を見つけるよ",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var criteria = 5

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncType)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		f := n.(*ast.FuncType)

		if isTooManyArgs(f) {
			pass.Reportf(f.Pos(), "引数が多いな！")
		}
	})
	return nil, nil
}

func isTooManyArgs(f *ast.FuncType) bool {
	return f.Params.NumFields() >= criteria
}
