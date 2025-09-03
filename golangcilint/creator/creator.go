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
	"bytes"

	"github.com/palantir/distgo/dister"
	"github.com/palantir/distgo/distgo"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/golangcilint"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/golangcilint/config"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func GolangCILint() dister.Creator {
	return dister.NewCreator(
		golangcilint.TypeName,
		func(cfgYML []byte) (distgo.Dister, error) {
			decoder := yaml.NewDecoder(bytes.NewReader(cfgYML))
			decoder.KnownFields(true)

			var distCfg config.GolangCILint
			if err := decoder.Decode(&distCfg); err != nil {
				return nil, errors.Wrapf(err, "failed to unmarshal dist-golangci-lint-asset configuration")
			}
			return distCfg.ToDister()
		},
	)
}
