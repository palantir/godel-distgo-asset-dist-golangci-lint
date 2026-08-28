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

package creator

import (
	"testing"

	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/golangcilint"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/golangcilint/config"
	"github.com/stretchr/testify/require"
)

func TestGolangCILintVersionConfiguration(t *testing.T) {
	upgradedCfg, err := config.UpgradeConfig([]byte("golangci-lint-version: v2.13.1\n"))
	require.NoError(t, err)

	dister, err := GolangCILint().Creator()(upgradedCfg)
	require.NoError(t, err)
	require.Equal(t, "v2.13.1", dister.(*golangcilint.Dister).Version)
}
