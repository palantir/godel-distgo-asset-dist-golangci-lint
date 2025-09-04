package nilerr

import (
	"github.com/gostaticanalysis/nilerr"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(nilerr.Analyzer).
		WithDesc("Find the code that returns nil even if it checks that the error is not nil.").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
