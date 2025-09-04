package testableexamples

import (
	"github.com/maratori/testableexamples/pkg/testableexamples"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(testableexamples.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
