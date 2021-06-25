package main

import (
	"github.com/edm20627/go_playground/analysis/dupimport"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(dupimport.Analyzer) }
