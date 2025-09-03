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

package main

import (
	"os"

	"github.com/palantir/amalgomate/amalgomated"
	"github.com/palantir/distgo/dister"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/golangcilint/config"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/golangcilint/creator"
	"github.com/palantir/pkg/cobracli"
)

func main() {
	os.Exit(amalgomated.RunApp(os.Args, nil, amalgomated.NewCmdLibrary(golangcilint.Instance()), pluginMain))
}

func pluginMain(osArgs []string) int {
	os.Args = osArgs
	return cobracli.ExecuteWithDefaultParams(dister.AssetRootCmd(creator.GolangCILint(), config.UpgradeConfig, "golangci-lint dist"))
}
