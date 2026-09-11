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
-: # github.com/palantir/godel-distgo-asset-dist-golangci-lint/golangcilint/integration_test [github.com/palantir/godel-distgo-asset-dist-golangci-lint/golangcilint/integration_test.test]
golangcilint/integration_test/integration_test.go:100:17: cannot use func(projectDir string) string {…} (value of type func(projectDir string) string) as func(t *testing.T, projectDir string) string value in struct literal
golangcilint/integration_test/integration_test.go:105:15: cannot use func(projectDir string) {…} (value of type func(projectDir string)) as func(t *testing.T, projectDir string) value in struct literal
golangcilint/integration_test/integration_test.go:189:17: cannot use func(projectDir string) string {…} (value of type func(projectDir string) string) as func(t *testing.T, projectDir string) string value in struct literal
golangcilint/integration_test/integration_test.go:192:15: cannot use func(projectDir string) {…} (value of type func(projectDir string)) as func(t *testing.T, projectDir string) value in struct literal
golangcilint/integration_test/integration_test.go:100:17: cannot use (func(projectDir string) string literal) (value of type func(projectDir string) string) as func(t *testing.T, projectDir string) string value in struct literal
golangcilint/integration_test/integration_test.go:105:15: cannot use (func(projectDir string) literal) (value of type func(projectDir string)) as func(t *testing.T, projectDir string) value in struct literal
golangcilint/integration_test/integration_test.go:189:17: cannot use (func(projectDir string) string literal) (value of type func(projectDir string) string) as func(t *testing.T, projectDir string) string value in struct literal
golangcilint/integration_test/integration_test.go:192:15: cannot use (func(projectDir string) literal) (value of type func(projectDir string)) as func(t *testing.T, projectDir string) value in struct literal
5 issues:
* compiles: 5

*/
