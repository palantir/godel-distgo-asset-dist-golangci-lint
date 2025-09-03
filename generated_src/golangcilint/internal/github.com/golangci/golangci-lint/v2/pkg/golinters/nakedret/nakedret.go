package nakedret

import (
	"github.com/alexkohler/nakedret/v2"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.NakedretSettings) *goanalysis.Linter {
	cfg := &nakedret.NakedReturnRunner{}

	if settings != nil {
		// SkipTestFiles is intentionally ignored => should be managed with `linters.exclusions.rules`.
		cfg.MaxLength = settings.MaxFuncLines
	}

	return goanalysis.
		NewLinterFromAnalyzer(nakedret.NakedReturnAnalyzer(cfg)).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
