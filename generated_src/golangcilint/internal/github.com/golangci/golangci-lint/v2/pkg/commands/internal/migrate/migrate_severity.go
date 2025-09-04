package migrate

import (
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/versionone"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/versiontwo"
)

func toSeverity(old *versionone.Config) versiontwo.Severity {
	var rules []versiontwo.SeverityRule

	for _, rule := range old.Severity.Rules {
		names := convertStaticcheckLinterNames(convertAlternativeNames(rule.Linters))
		if len(rule.Linters) > 0 && len(names) == 0 {
			continue
		}

		rules = append(rules, versiontwo.SeverityRule{
			BaseRule: versiontwo.BaseRule{
				Linters:	names,
				Path:		rule.Path,
				PathExcept:	rule.PathExcept,
				Text:		rule.Text,
				Source:		rule.Source,
			},
			Severity:	rule.Severity,
		})
	}

	return versiontwo.Severity{
		Default:	old.Severity.Default,
		Rules:		rules,
	}
}
