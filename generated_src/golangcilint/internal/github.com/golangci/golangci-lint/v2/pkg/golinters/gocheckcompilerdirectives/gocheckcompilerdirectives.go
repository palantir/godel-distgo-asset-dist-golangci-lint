package gocheckcompilerdirectives

import (
	"4d63.com/gocheckcompilerdirectives/checkcompilerdirectives"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(checkcompilerdirectives.Analyzer()).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
