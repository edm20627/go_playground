package main

import (
	"github.com/edm20627/go_playground/analysis/myanalyzer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(myanalyzer.Analyzer) }
