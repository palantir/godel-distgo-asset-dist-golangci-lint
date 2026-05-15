package clickhouselint

import (
	"github.com/ClickHouse/clickhouse-go-linter/passes/chbatchclose"
	"github.com/ClickHouse/clickhouse-go-linter/passes/chrowserr"
	"golang.org/x/tools/go/analysis"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"clickhouselint",
		"Detects common mistakes with the ClickHouse native Go driver API.",
		[]*analysis.Analyzer{chrowserr.NewAnalyzer(), chbatchclose.NewAnalyzer()},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
