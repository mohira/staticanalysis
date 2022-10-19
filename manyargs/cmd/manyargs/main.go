package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"
	"manyargs"
)

func main() {
	unitchecker.Main(manyargs.Analyzer)
}
