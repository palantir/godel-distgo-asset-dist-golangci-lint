package fail

fail

/*
This is a non-compiling file that has been added to explicitly ensure that CI fails.
It also contains the command that caused the failure and its output.
Remove this file if debugging locally.

go mod operation failed. This may mean that there are legitimate dependency issues with the "go.mod" definition in the repository and the updates performed by the gomod check. This branch can be cloned locally to debug the issue.

Command that caused error:
./godelw mod

Output:
go: finding module for package github.com/golangci/gofmt/gofmt
go: github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/goformatters/gofmt imports
	github.com/golangci/gofmt/gofmt: module github.com/golangci/gofmt@latest found (v0.0.0-20251215234548-e7be49a5ab4d), but does not contain package github.com/golangci/gofmt/gofmt

*/
