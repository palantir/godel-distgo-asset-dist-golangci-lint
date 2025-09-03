package nilnesserr

import (
	"github.com/alingse/nilnesserr"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/golinters/internal"
)

func New() *goanalysis.Linter {
	analyzer, err := nilnesserr.NewAnalyzer(nilnesserr.LinterSetting{})
	if err != nil {
		internal.LinterLogger.Fatalf("nilnesserr: create analyzer: %v", err)
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
