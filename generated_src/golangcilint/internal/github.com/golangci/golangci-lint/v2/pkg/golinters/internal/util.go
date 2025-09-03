package internal

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goanalysis"
)

func FormatCode(code string) string {
	if strings.Contains(code, "`") {
		return code	// TODO: properly escape or remove
	}

	return fmt.Sprintf("`%s`", code)
}

func GetGoFileNames(pass *analysis.Pass) []string {
	var filenames []string

	for _, f := range pass.Files {
		position, b := goanalysis.GetGoFilePosition(pass, f)
		if !b {
			continue
		}

		filenames = append(filenames, position.Filename)
	}

	return filenames
}
