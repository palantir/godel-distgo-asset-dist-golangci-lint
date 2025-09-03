package errchkjson

import (
	"github.com/breml/errchkjson"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.ErrChkJSONSettings) *goanalysis.Linter {
	cfg := map[string]any{
		"omit-safe": true,
	}

	if settings != nil {
		cfg = map[string]any{
			"omit-safe":		!settings.CheckErrorFreeEncoding,
			"report-no-exported":	settings.ReportNoExported,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(errchkjson.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
