package paralleltest

import (
	"github.com/kunwardeep/paralleltest/pkg/paralleltest"
	"golang.org/x/tools/go/analysis"
)

var Analyzer *analysis.Analyzer

func init() {
	a := paralleltest.NewAnalyzer()
	flags := a.Flags
	if err := flags.Set("ignoreloopVar", "true"); err != nil {
		panic(err)
	}
	a.Flags = flags
	Analyzer = a
}
