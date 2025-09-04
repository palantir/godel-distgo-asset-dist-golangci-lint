package err113

import (
	"github.com/Djarvur/go-err113"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(err113.NewAnalyzer()).
		WithDesc("Check errors handling expressions").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
