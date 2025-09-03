package durationcheck

import (
	"github.com/charithe/durationcheck"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(durationcheck.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
