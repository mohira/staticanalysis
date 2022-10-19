package manyargs_test

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"manyargs"
	"testing"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, manyargs.Analyzer, "a")
}
