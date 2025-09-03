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

package config

import (
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/golangcilint"
	v0 "github.com/palantir/godel-distgo-asset-dist-golangci-lint/golangcilint/config/internal/v0"
)

// Default version of golangci-lint to compile for output binary if no version is specified in configuration.
// Should generally track/match the latest golangci-lint release: https://github.com/golangci/golangci-lint/releases
const defaultVersion = "v2.4.0"

type GolangCILint v0.Config

func (cfg *GolangCILint) ToDister() (*golangcilint.Dister, error) {
	var plugins []*golangcilint.Plugin
	for _, plugin := range cfg.Plugins {
		plugins = append(plugins, &golangcilint.Plugin{
			Module:  plugin.Module,
			Import:  plugin.Import,
			Version: plugin.Version,
			Path:    plugin.Path,
		})
	}

	version := cfg.Version
	// if version in configuration is empty, use a default version
	if version == "" {
		version = defaultVersion
	}

	return &golangcilint.Dister{
		OSArchs:     cfg.OSArchs,
		Environment: cfg.Environment,
		Version:     version,
		Plugins:     plugins,
	}, nil
}
