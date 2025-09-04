package contextcheck

import (
	"github.com/kkHAIKE/contextcheck"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/lint/linter"
)

func New() *goanalysis.Linter {
	analyzer := contextcheck.NewAnalyzer(contextcheck.Configuration{})

	return goanalysis.
		NewLinterFromAnalyzer(analyzer).
		WithContextSetter(func(lintCtx *linter.Context) {
			analyzer.Run = contextcheck.NewRun(lintCtx.Packages, false)
		}).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
