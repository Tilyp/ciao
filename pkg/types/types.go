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

package types

import (
	"github.com/caicloud/ciao/pkg/resource"
)

// Parameter is the type for the parameter of the job.
type Parameter struct {
	Framework   FrameworkType
	PSCount     int
	WorkerCount int
	// MasterCount is for PyTorchJob.
	MasterCount  int
	GenerateName string
	Image        string
	CleanPolicy  string
	Resource     resource.JobResource
}

// Job is the type for general jobs (PyTorchJob, TFJob).
type Job struct {
	Framework FrameworkType
	Name      string
	PS        int
	// Master is for PyTorchJob.
	Master int
	Worker int
}

// FrameworkType is the type for types of the frameworks.
type FrameworkType string

const (
	// FrameworkTypeTensorFlow defines tensorflow type.
	FrameworkTypeTensorFlow = "tensorflow"
	// FrameworkTypePyTorch defines pytorch type.
	FrameworkTypePyTorch = "pytorch"
)

// CleanPodPolicy describes how to deal with pods when the job is finished.
const (
	CleanPodPolicyAll     = "all"
	CleanPodPolicyRunning = "running"
	CleanPodPolicyNone    = "none"
)
