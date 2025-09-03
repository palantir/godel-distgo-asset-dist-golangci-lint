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

package v0

import (
	"github.com/palantir/distgo/dister/osarchbin/config"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

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

// Plugin represents information about a plugin.
// Corresponds to the internal golangci-lint configuration struct github.com/golangci/golangci-lint/pkg/commands/internal.Plugin (https://github.com/golangci/golangci-lint/blob/7ad7949ca9bf236ee4f349de2cb384d5f7c90b08/pkg/commands/internal/configuration.go#L79).
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

func UpgradeConfig(cfgBytes []byte) ([]byte, error) {
	var cfg Config
	if err := yaml.UnmarshalStrict(cfgBytes, &cfg); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal dist-sls-asset v0 configuration")
	}
	return cfgBytes, nil
}
