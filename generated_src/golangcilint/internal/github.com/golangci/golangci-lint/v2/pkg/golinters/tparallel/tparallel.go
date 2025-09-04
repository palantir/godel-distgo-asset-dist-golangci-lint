package tparallel

import (
	"github.com/moricho/tparallel"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(tparallel.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
