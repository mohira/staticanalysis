package main

import (
	"unnamedfunc"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(unnamedfunc.Analyzer) }
