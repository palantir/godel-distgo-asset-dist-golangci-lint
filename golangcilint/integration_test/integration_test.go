// Copyright 2025 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integration

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/nmiyake/pkg/gofiles"
	"github.com/palantir/distgo/dister/distertester"
	"github.com/palantir/godel/v2/framework/pluginapitester"
	"github.com/palantir/godel/v2/pkg/products"
	"github.com/palantir/pkg/specdir"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	distgoPluginLocator  = "com.palantir.distgo:dist-plugin:1.66.0"
	distgoPluginResolver = "https://github.com/{{index GroupParts 1}}/{{index GroupParts 2}}/releases/download/v{{Version}}/{{Product}}-{{Version}}-{{OS}}-{{Arch}}.tgz"
)

func TestDist(t *testing.T) {
	t.Setenv("INTOTO_DISABLED", "true")
	t.Setenv("ITAR", "")
	t.Setenv("RELEASE_TYPE", "")

	const godelYML = `exclude:
  names:
    - "\\..+"
    - "vendor"
  paths:
    - "godel"
`

	pluginProvider, err := pluginapitester.NewPluginProviderFromLocator(distgoPluginLocator, distgoPluginResolver)
	require.NoError(t, err)

	assetPath, err := products.Bin("dist-golangci-lint-asset")
	require.NoError(t, err)

	distertester.RunAssetDistTest(t,
		pluginProvider,
		pluginapitester.NewAssetProvider(assetPath),
		[]distertester.TestCase{
			{
				Name: "generates golangci-lint distribution",
				Specs: []gofiles.GoFileSpec{
					{
						RelPath: "go.mod",
						Src:     "module foo",
					},
					{
						RelPath: "foo/foo.go",
						Src:     `package main; func main() {}`,
					},
				},
				ConfigFiles: map[string]string{
					"godel/config/godel.yml": godelYML,
					"godel/config/dist-plugin.yml": `
products:
  foo:
    dist:
      disters:
        type: golangci-lint
        config:
          os-archs:
            - os: darwin
              arch: amd64
            - os: linux
              arch: amd64
          # set as environment variables when building the binary
          environment:
            CGO_ENABLED: "0"
            GOFLAGS: "-mod=readonly"
          golangci-lint-version: v2.13.0
          plugins:
            - module: 'github.com/golangci/example-plugin-module-linter'
              version: v0.1.0
    publish:
      group-id: com.test.group
`,
				},
				WantOutput: func(projectDir string) string {
					return `Creating distribution for foo at out/dist/foo/1.0.0/golangci-lint/foo-1.0.0-darwin-amd64.tgz, out/dist/foo/1.0.0/golangci-lint/foo-1.0.0-linux-amd64.tgz
Finished creating golangci-lint distribution for foo
`
				},
				Validate: func(projectDir string) {
					// verify that work directory and output TGZs created
					wantLayout := specdir.NewLayoutSpec(
						specdir.Dir(specdir.LiteralName("golangci-lint"), "",
							specdir.File(specdir.LiteralName("foo-1.0.0-darwin-amd64.tgz"), ""),
							specdir.File(specdir.LiteralName("foo-1.0.0-linux-amd64.tgz"), ""),
							specdir.Dir(specdir.LiteralName("foo-1.0.0"), "",
								specdir.Dir(specdir.LiteralName("darwin-amd64"), "",
									specdir.File(specdir.LiteralName("foo"), ""),
								),
								specdir.Dir(specdir.LiteralName("linux-amd64"), "",
									specdir.File(specdir.LiteralName("foo"), ""),
								),
							),
						), true,
					)
					assert.NoError(t, wantLayout.Validate(filepath.Join(projectDir, "out", "dist", "foo", "1.0.0", "golangci-lint"), nil))
				},
			},
		},
	)
}

func TestDistUsesTargetProjectGoToolchain(t *testing.T) {
	const (
		supportedGoVersion = "1.26"
		targetGoToolchain  = "go1.27.0"
	)

	t.Setenv("INTOTO_DISABLED", "true")
	t.Setenv("ITAR", "")
	t.Setenv("RELEASE_TYPE", "")

	pluginProvider, err := pluginapitester.NewPluginProviderFromLocator(distgoPluginLocator, distgoPluginResolver)
	require.NoError(t, err)

	assetPath, err := products.Bin("dist-golangci-lint-asset")
	require.NoError(t, err)

	hostOSArch := runtime.GOOS + "-" + runtime.GOARCH
	distertester.RunAssetDistTest(t,
		pluginProvider,
		pluginapitester.NewAssetProvider(assetPath),
		[]distertester.TestCase{
			{
				Name: "uses target project Go toolchain",
				Specs: []gofiles.GoFileSpec{
					{
						RelPath: "go.mod",
						Src:     "module foo\n\ngo " + supportedGoVersion + ".0\n\ntoolchain " + targetGoToolchain,
					},
					{
						RelPath: "foo/foo.go",
						Src:     `package main; func main() {}`,
					},
				},
				ConfigFiles: map[string]string{
					"godel/config/godel.yml": `exclude:
  names:
    - "\\..+"
    - "vendor"
  paths:
    - "godel"
`,
					"godel/config/dist-plugin.yml": fmt.Sprintf(`
products:
  foo:
    dist:
      disters:
        type: golangci-lint
        config:
          os-archs:
            - os: %s
              arch: %s
          environment:
            CGO_ENABLED: "0"
            GOFLAGS: "-mod=readonly"
          plugins:
            - module: 'github.com/golangci/example-plugin-module-linter'
              version: v0.1.0
    publish:
      group-id: com.test.group
`, runtime.GOOS, runtime.GOARCH),
				},
				WantOutput: func(projectDir string) string {
					return fmt.Sprintf("Creating distribution for foo at out/dist/foo/1.0.0/golangci-lint/foo-1.0.0-%s.tgz\nFinished creating golangci-lint distribution for foo\n", hostOSArch)
				},
				Validate: func(projectDir string) {
					executableName := "foo"
					if runtime.GOOS == "windows" {
						executableName += ".exe"
					}
					binaryPath := filepath.Join(projectDir, "out", "dist", "foo", "1.0.0", "golangci-lint", "foo-1.0.0", hostOSArch, executableName)

					versionOutput, err := exec.Command("go", "version", "-m", binaryPath).CombinedOutput()
					t.Logf("go version -m %s:\n%s", binaryPath, versionOutput)
					require.NoError(t, err)
					assert.Contains(t, string(versionOutput), ": "+targetGoToolchain+"\n")

					lintDir := t.TempDir()
					require.NoError(t, os.WriteFile(filepath.Join(lintDir, "go.mod"), []byte("module example.com/integration\n\ngo "+supportedGoVersion+".0\n"), 0o644))
					require.NoError(t, os.WriteFile(filepath.Join(lintDir, "main.go"), []byte("package integration\n"), 0o644))

					lintCmd := exec.Command(binaryPath, "run")
					lintCmd.Dir = lintDir
					lintOutput, err := lintCmd.CombinedOutput()
					t.Logf("golangci-lint output:\n%s", lintOutput)
					require.NoError(t, err)
				},
			},
		},
	)
}
