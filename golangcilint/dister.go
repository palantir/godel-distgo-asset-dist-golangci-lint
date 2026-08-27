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

package golangcilint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/palantir/distgo/dister/osarchbin"
	"github.com/palantir/distgo/distgo"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/runner"
	"github.com/palantir/godel/v2/pkg/osarch"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const TypeName = "golangci-lint"

// Configuration represents the configuration file.
// Corresponds to the internal golangci-lint configuration struct github.com/golangci/golangci-lint/pkg/commands/internal.Configuration (https://github.com/golangci/golangci-lint/blob/7ad7949ca9bf236ee4f349de2cb384d5f7c90b08/pkg/commands/internal/configuration.go#L18).
type Configuration struct {
	// golangci-lint version.
	Version string `yaml:"version"`

	// Name of the binary.
	Name string `yaml:"name,omitempty"`

	// Destination is the path to a directory to store the binary.
	Destination string `yaml:"destination,omitempty"`

	// Plugins information.
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

type Dister struct {
	// The OS/Arch targets for which to build the golangci-lint binary.
	OSArchs []osarch.OSArch

	// Environment variables to set when building the golangci-lint binary.
	Environment map[string]string

	// The version of golangci-lint source that is fetched to build the binary.
	Version string

	// The plugins to include in the golangci-lint binary.
	Plugins []*Plugin
}

func (d *Dister) TypeName() (string, error) {
	return TypeName, nil
}

func (d *Dister) Artifacts(renderedName string) ([]string, error) {
	// delegate to the osarchbin dister
	return d.getOSArchBinDister().Artifacts(renderedName)
}

func (d *Dister) PackagingExtension() (string, error) {
	// delegate to the osarchbin dister
	return d.getOSArchBinDister().PackagingExtension()
}

func (d *Dister) getOSArchBinDister() distgo.Dister {
	return osarchbin.New(d.OSArchs...)
}

func (d *Dister) RunDist(distID distgo.DistID, productTaskOutputInfo distgo.ProductTaskOutputInfo) ([]byte, error) {
	goToolchain, err := projectGoToolchain(productTaskOutputInfo.Project.ProjectDir)
	if err != nil {
		return nil, errors.Wrap(err, "failed to determine target project Go toolchain")
	}

	distWorkDir := productTaskOutputInfo.ProductDistWorkDirs()[distID]
	outputPathsForOSArchs := make(map[string][]string)
	for _, osArch := range d.OSArchs {
		currOutputBuildDir := filepath.Join(distWorkDir, osArch.String())
		if err := os.MkdirAll(currOutputBuildDir, 0755); err != nil {
			return nil, errors.Wrapf(err, "failed to create directory for golangci-lint build for OS/Arch %s", osArch.String())
		}

		binaryName := productTaskOutputInfo.Product.Name
		configFilesWritten, err := d.writeGolangCILintCustomBuildConfig(currOutputBuildDir, binaryName, productTaskOutputInfo.Project.ProjectDir)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to write golangci-lint custom build configuration")
		}

		// run amalgomated build in directory with OS/Arch set
		outputBuf := &bytes.Buffer{}
		if rVal, err := runner.RunAmalgomatedGolangCILint([]string{"custom", "-v"}, func(cmd *exec.Cmd) {
			cmd.Stdout = outputBuf
			cmd.Stderr = outputBuf
			envVars := cmd.Environ()
			for k, v := range d.Environment {
				envVars = setEnvironmentVariable(envVars, k, v)
			}
			envVars = setEnvironmentVariable(envVars, "GOTOOLCHAIN", goToolchain)
			envVars = setEnvironmentVariable(envVars, "GOOS", osArch.OS)
			envVars = setEnvironmentVariable(envVars, "GOARCH", osArch.Arch)
			cmd.Env = envVars
			cmd.Dir = currOutputBuildDir
		}); rVal != 0 || err != nil {
			if err != nil {
				return nil, errors.Wrapf(err, "error running amalgomated golangci-lint custom build")
			}
			return nil, errors.New(fmt.Sprintf("running amalgomated golangci-lint custom build failed with exit code %d. Output: %s", rVal, outputBuf.String()))
		}

		// clean up any config files that were written
		for _, configFilePath := range configFilesWritten {
			if err := os.Remove(configFilePath); err != nil {
				return nil, errors.Wrapf(err, "failed to remove golangci-lint custom build configuration file %s", configFilePath)
			}
		}

		binaryPath := filepath.Join(distWorkDir, osArch.String(), distgo.ExecutableName(binaryName, osArch.OS))
		outputPathsForOSArchs[osArch.String()] = append(outputPathsForOSArchs[osArch.String()], binaryPath)
	}
	jsonBytes, err := json.Marshal(outputPathsForOSArchs)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal outputPathsForOSArchs as JSON")
	}
	return jsonBytes, nil
}

func projectGoToolchain(projectDir string) (string, error) {
	cmd := exec.Command("go", "env", "GOVERSION")
	cmd.Dir = projectDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrapf(err, "command %q failed: %s", strings.Join(cmd.Args, " "), strings.TrimSpace(string(output)))
	}

	goToolchain := strings.TrimSpace(string(output))
	if !strings.HasPrefix(goToolchain, "go1.") {
		return "", errors.Errorf("command %q returned invalid Go toolchain version %q", strings.Join(cmd.Args, " "), goToolchain)
	}
	return goToolchain, nil
}

func setEnvironmentVariable(environment []string, key, value string) []string {
	prefix := key + "="
	result := make([]string, 0, len(environment)+1)
	for _, environmentVariable := range environment {
		if !strings.HasPrefix(environmentVariable, prefix) {
			result = append(result, environmentVariable)
		}
	}
	return append(result, fmt.Sprintf("%s=%s", key, value))
}

func (d *Dister) GenerateDistArtifacts(distID distgo.DistID, productTaskOutputInfo distgo.ProductTaskOutputInfo, runDistResult []byte) error {
	// delegate to the osarchbin dister
	return d.getOSArchBinDister().GenerateDistArtifacts(distID, productTaskOutputInfo, runDistResult)
}

// writes the golangci-lint custom build configuration to the specified directory and returns the paths of the written
// configuration files.
func (d *Dister) writeGolangCILintCustomBuildConfig(dir, renderedName, projectDir string) ([]string, error) {
	pluginsWithRenderedPath := make([]*Plugin, len(d.Plugins))
	for i, currPlugin := range d.Plugins {
		pluginCopy := *currPlugin
		renderedPath, err := distgo.RenderTemplate(pluginCopy.Path, nil, distgo.TemplateValueFunction("ProjectDir", projectDir))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to render golangci-lint plugin path %s as template", pluginCopy.Path)
		}
		pluginCopy.Path = renderedPath
		pluginsWithRenderedPath[i] = &pluginCopy
	}
	configYAMLBytes, err := yaml.Marshal(Configuration{
		Version: d.Version,
		Name:    renderedName,
		Plugins: pluginsWithRenderedPath,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal golangci-lint configuration as YAML")
	}
	configFilePath := filepath.Join(dir, ".custom-gcl.yml")
	if err := os.WriteFile(configFilePath, configYAMLBytes, 0644); err != nil {
		return nil, errors.Wrapf(err, "failed to write golangci-lint configuration to %s", configFilePath)
	}
	return []string{configFilePath}, nil
}
