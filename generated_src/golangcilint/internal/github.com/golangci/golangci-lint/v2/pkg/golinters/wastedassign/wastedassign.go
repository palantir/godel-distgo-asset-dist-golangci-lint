package wastedassign

import (
	"github.com/sanposhiho/wastedassign/v2"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(wastedassign.Analyzer).
		WithDesc("Finds wasted assignment statements").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
