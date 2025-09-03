package noctx

import (
	"github.com/sonatard/noctx"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(noctx.Analyzer).
		WithDesc("Detects function and method with missing usage of context.Context").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
