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

package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/palantir/amalgomate/amalgomated"
	"github.com/palantir/godel-distgo-asset-dist-golangci-lint/generated_src/golangcilint"
	"github.com/pkg/errors"
)

func RunAmalgomatedGolangCILint(args []string, configureCmd CmdConfigurer) (int, error) {
	runner, err := AmalgomatedGolangCILintCmdRunner(args, configureCmd)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to create amalgomated golangci-lint command")
	}
	return runner(), nil
}

type CmdConfigurer func(*exec.Cmd)

func AmalgomatedGolangCILintCmdRunner(args []string, configureCmd CmdConfigurer) (func() int, error) {
	cmd, err := newAmalgomatedGolangCILintCmd(args)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create amalgomated golangci-lint command")
	}
	if configureCmd != nil {
		configureCmd(cmd)
	}

	return func() int {
		if err := cmd.Run(); err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				return exitErr.ExitCode()
			}
			if cmd.Stderr != nil {
				_, _ = fmt.Fprintf(cmd.Stderr, "command %v failed with error: %v\n", cmd.Args, errors.Wrapf(err, "run error"))
			}
			return 1
		}
		return 0
	}, nil
}

func newAmalgomatedGolangCILintCmd(args []string) (*exec.Cmd, error) {
	pathToSelf, err := os.Executable()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find self executable")
	}

	amalgomatedCmds := golangcilint.Instance().Cmds()
	if len(amalgomatedCmds) != 1 {
		return nil, errors.Errorf("expected exactly one amalgomated command, got %d: %v", len(amalgomatedCmds), amalgomatedCmds)
	}

	return exec.Command(pathToSelf, append([]string{amalgomated.ProxyCmdPrefix + amalgomatedCmds[0]}, args...)...), nil
}
