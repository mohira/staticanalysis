package manyargs

import (
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "manyargs",
	Doc:      "引数が多い関数を見つけるよ",
	Run:      run,
	Requires: nil,
}

func run(pass *analysis.Pass) (any, error) {
	return nil, nil
}
