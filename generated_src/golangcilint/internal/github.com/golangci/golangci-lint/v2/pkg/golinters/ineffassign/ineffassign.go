package ineffassign

import (
	"github.com/gordonklaus/ineffassign/pkg/ineffassign"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(ineffassign.Analyzer).
		WithDesc("Detects when assignments to existing variables are not used").
		WithLoadMode(goanalysis.LoadModeSyntax)
}
