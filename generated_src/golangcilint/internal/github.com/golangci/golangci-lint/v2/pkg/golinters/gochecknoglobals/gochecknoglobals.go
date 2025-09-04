package gochecknoglobals

import (
	"4d63.com/gochecknoglobals/checknoglobals"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(checknoglobals.Analyzer()).
		WithDesc("Check that no global variables exist.").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
