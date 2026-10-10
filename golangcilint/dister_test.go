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

package golangcilint

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectGoToolchain(t *testing.T) {
	for _, test := range []struct {
		name  string
		goMod string
		want  string
	}{
		{
			name:  "toolchain directive",
			goMod: "module example.com/foo\n\ngo 1.26.0\n\ntoolchain go1.27.0\n",
			want:  "go1.27.0",
		},
		{
			name:  "Go directive with patch version",
			goMod: "module example.com/foo\n\ngo 1.27.1\n",
			want:  "go1.27.1",
		},
		{
			name:  "Go directive without patch version",
			goMod: "module example.com/foo\n\ngo 1.27\n",
			want:  "go1.27.0",
		},
		{
			name:  "no version directive",
			goMod: "module example.com/foo\n",
			want:  runtime.Version(),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			projectDir := t.TempDir()
			require.NoError(t, os.WriteFile(filepath.Join(projectDir, "go.mod"), []byte(test.goMod), 0o644))

			got, err := projectGoToolchain(projectDir)
			require.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
