package whitespace

import (
	"github.com/ultraware/whitespace"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.WhitespaceSettings) *goanalysis.Linter {
	var wsSettings whitespace.Settings
	if settings != nil {
		wsSettings = whitespace.Settings{
			MultiIf:	settings.MultiIf,
			MultiFunc:	settings.MultiFunc,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(whitespace.NewAnalyzer(&wsSettings)).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
