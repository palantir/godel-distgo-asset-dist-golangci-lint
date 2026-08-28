// Copyright 2026 Palantir Technologies, Inc.
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

//go:generate go run $GOFILE

package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

const golangCILintModule = "github.com/golangci/golangci-lint/v2"

func main() {
	goModBytes, err := os.ReadFile("../../../go.mod")
	if err != nil {
		panic(fmt.Errorf("read go.mod: %w", err))
	}

	goMod, err := modfile.Parse("go.mod", goModBytes, nil)
	if err != nil {
		panic(fmt.Errorf("parse go.mod: %w", err))
	}

	for _, requirement := range goMod.Require {
		if requirement.Mod.Path == golangCILintModule {
			if err := os.WriteFile("golangci-lint-version.txt", []byte(requirement.Mod.Version), 0o644); err != nil {
				panic(fmt.Errorf("write golangci-lint version: %w", err))
			}
			return
		}
	}

	panic(fmt.Errorf("module %q is not required by go.mod", golangCILintModule))
}
