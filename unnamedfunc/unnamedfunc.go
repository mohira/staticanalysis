package unnamedfunc

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "unnamedfunc",
	Doc:      "名前無し引数を使っているパッケージ関数を探してその使いみちを知る",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		f := n.(*ast.FuncDecl)

		if f.Recv != nil {
			return
		}

		var isUnnamed bool
		for _, p := range f.Type.Params.List {
			if p.Names == nil {
				isUnnamed = true
			}
		}

		if isUnnamed {
			pass.Reportf(f.Pos(), "%s 名前無し引数を使った関数", f.Name.Name)
		}
	})

	return nil, nil
}
