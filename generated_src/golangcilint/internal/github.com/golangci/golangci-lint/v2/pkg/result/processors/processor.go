package processors

import (
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/result"
)

const typeCheckName = "typecheck"

type Processor interface {
	Process(issues []*result.Issue) ([]*result.Issue, error)
	Name() string
	Finish()
}
