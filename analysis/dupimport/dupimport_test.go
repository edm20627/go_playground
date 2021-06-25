package dupimport_test

import (
	"testing"

	"github.com/edm20627/go_playground/analysis/dupimport"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, dupimport.Analyzer, "a")
}
