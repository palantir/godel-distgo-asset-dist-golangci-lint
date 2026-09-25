package fail

fail

/*
This is a non-compiling file that has been added to explicitly ensure that CI fails.
It also contains the command that caused the failure and its output.
Remove this file if debugging locally.

go mod operation failed. This may mean that there are legitimate dependency issues with the "go.mod" definition in the repository and the updates performed by the bump-go-dependencies check. This branch can be cloned locally to debug the issue.

Command that caused error:
./godelw lint --fix

Output:
-: # github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/golinters/gochecksumtype
generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/golinters/gochecksumtype/gochecksumtype.go:60:24: undefined: gochecksumtype.Config
generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/golinters/gochecksumtype/gochecksumtype.go:66:27: undefined: gochecksumtype.Run
generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/golinters/gochecksumtype/gochecksumtype.go:60:24: undefined: gochecksumtype.Config
generated_src/golangcilint/internal/github.com/golangci/golangci-lint/v2/pkg/golinters/gochecksumtype/gochecksumtype.go:66:27: undefined: gochecksumtype.Run
3 issues:
* compiles: 3

*/
