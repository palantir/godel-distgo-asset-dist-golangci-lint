<p align="right">
<a href="https://autorelease.general.dmz.palantir.tech/palantir/godel-distgo-asset-dist-golangci-lint"><img src="https://img.shields.io/badge/Perform%20an-Autorelease-success.svg" alt="Autorelease"></a>
</p>

# godel-distgo-asset-dist-golangci-lint
godel-distgo-asset-dist-golangci-lint is an asset for the gödel [dist plugin](https://github.com/palantir/distgo). It provides the `golangci-lint`
dister, which is a dister that creates os-arch-bin distributions for the golangci-lint plugin with custom modules.

This dister encapsulates the logic for building a `golangci-lint` binary with custom modules, as documented by the
[golangci-lint documentation](https://golangci-lint.run/plugins/module-plugins/). This dister amalgomates the
`golangci-lint` binary internally and runs the `golangci-lint custom` command to build the binary with the custom
modules and has support for building the binary for multiple OS/Arch output targets.

The version of `golangci-lint` that is amalgomated (which is the version that is used to run the `custom` command) is
controlled by the version of the `github.com/golangci/golangci-lint` module dependency in `go.mod`. The
`github.com/golangci/golangci-lint/v2/cmd/golangci-lint` module is added as a [`tool` dependency](https://go.dev/ref/mod#go-mod-file-tool)
in `go.mod`.

## Output
For each OS/Arch specified in the dister configuration, the dister produces a dist output named
`{Product}-{Version}-{OS}-{Arch}.tar.gz` that contains the single `golangci-lint` binary with the name `{Product}`.

## Configuration
The configuration for the dister is defined in the `GolangCILint` struct in `golangcilint/config/config.go` and is as
follows:

```
type Config struct {
	config.OSArchBin `yaml:",inline"`

	// Environment is a map of environment variables to set when building the binary.
	Environment map[string]string `yaml:"environment,omitempty"`

	// Version is the golangci-lint version of the source used to build the binary.
	// If empty, uses a default version hard-coded in the dister.
	// Corresponds to the internal golangci-lint configuration field github.com/golangci/golangci-lint/pkg/commands/internal.Configuration.Version (https://github.com/golangci/golangci-lint/blob/7ad7949ca9bf236ee4f349de2cb384d5f7c90b08/pkg/commands/internal/configuration.go#L20).
	Version string `yaml:"version,omitempty"`

	// Plugins configuration that specifies the plugins to include in the custom build.
	// Corresponds to the internal golangci-lint configuration field github.com/golangci/golangci-lint/pkg/commands/internal.Configuration.Plugins (https://github.com/golangci/golangci-lint/blob/7ad7949ca9bf236ee4f349de2cb384d5f7c90b08/pkg/commands/internal/configuration.go#L29).
	Plugins []*Plugin `yaml:"plugins,omitempty"`
}
```

The `Plugin` struct type matches the one defined by [`golangci-lint`](https://github.com/golangci/golangci-lint/blob/7ad7949ca9bf236ee4f349de2cb384d5f7c90b08/pkg/commands/internal/configuration.go#L79):

```
type Plugin struct {
	// Module name.
	Module string `yaml:"module"`

	// Import to use.
	Import string `yaml:"import,omitempty"`

	// Version of the module.
	// Only for module available through a Go proxy.
	Version string `yaml:"version,omitempty"`

	// Path to the local module.
	// Only for local module.
	Path string `yaml:"path,omitempty"`
}
```

The dister configuration renders the content of the `Path` field of plugins as a Go template, with the function
`{{ProjectDir}}` defined to be the path to the root of the project being built.

Here is an example of what a typical configuration for the dister might look like in `dist-plugin.yml`:

```
products:
  golangci-lint-palantir:
    dist:
      disters:
        type: golangci-lint
        config:
          # defines the OS/Arch combinations for which binaries and distributions are built
          os-archs:
            - os: darwin
              arch: amd64
            - os: darwin
              arch: arm64
            - os: linux
              arch: amd64
            - os: linux
              arch: arm64
          # set as environment variables when building the binary
          environment:
            CGO_ENABLED: "0"
            GOFLAGS: "-mod=readonly"
          # version of golangci-lint source to use
          version: v2.4.0
          # plugins to include in the custom build
          plugins:
            # a plugin from a Go proxy
            - module: 'github.com/golangci/plugin1'
              import: 'github.com/golangci/plugin1/foo'
              version: v1.0.0
            
            # a plugin from local source
            - module: 'github.com/golangci/plugin2'
              path: '{ProjectDir}/my/local/path/plugin2'
```
