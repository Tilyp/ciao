// Copyright 2018 Caicloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package simple

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/caicloud/ciao/pkg/types"
)

// Interpreter is the type for the simple interpreter.
type Interpreter struct {
	FrameworkPrefix   string
	WorkerPrefix      string
	PSPrefix          string
	MasterPrefix      string
	CleanPolicyPrefix string
}

// New returns a new interpreter.
func New() *Interpreter {
	return &Interpreter{
		FrameworkPrefix:   "%framework=",
		WorkerPrefix:      "%worker=",
		PSPrefix:          "%ps=",
		MasterPrefix:      "%master=",
		CleanPolicyPrefix: "%cleanPolicy=",
	}
}

// Preprocess interprets the magic commands.
func (i Interpreter) Preprocess(code string) (*types.Parameter, error) {
	param := &types.Parameter{}
	lines := strings.Split(code, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if string(line[0]) == "%" {
			if err := i.parseMagicCommand(param, line); err != nil {
				return nil, err
			}
		}
	}
	return param, nil
}

func (i Interpreter) parseMagicCommand(param *types.Parameter, line string) error {
	var err error
	switch {
	case strings.Contains(line, i.FrameworkPrefix):
		param.Framework = types.FrameworkType(line[len(i.FrameworkPrefix):])
	case strings.Contains(line, i.WorkerPrefix):
		param.WorkerCount, err = strconv.Atoi(line[len(i.WorkerPrefix):])
		if err != nil {
			return err
		}
	case strings.Contains(line, i.PSPrefix):
		param.PSCount, err = strconv.Atoi(line[len(i.PSPrefix):])
		if err != nil {
			return err
		}
	case strings.Contains(line, i.MasterPrefix):
		param.MasterCount, err = strconv.Atoi(line[len(i.MasterPrefix):])
		if err != nil {
			return err
		}
	case strings.Contains(line, i.CleanPolicyPrefix):
		// Set default clean pod policy to None.
		param.CleanPolicy = types.CleanPodPolicyNone
		policy := line[len(i.CleanPolicyPrefix):]
		if policy == types.CleanPodPolicyAll || policy == types.CleanPodPolicyRunning {
			param.CleanPolicy = policy
		}
	}

	return nil
}

// PreprocessedCode gets the preprocessed code ( the code without magic commands.)
func (i Interpreter) PreprocessedCode(code string) string {
	lines := strings.Split(code, "\n")
	res := ""
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if string(line[0]) != "%" {
			res = fmt.Sprintf("%s\n%s", res, line)
		}
	}
	return res
}
