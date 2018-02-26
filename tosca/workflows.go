// Copyright 2018 Bull S.A.S. Atos Technologies - Bull, Rue Jean Jaures, B.P.68, 78340, Les Clayes-sous-Bois, France.
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

package tosca

// An Workflow is the representation of a TOSCA Workflow
//
// Currently Workflows are not part of the TOSCA specification
type Workflow struct {
	Steps map[string]Step `yaml:"steps,omitempty" json:"steps,omitempty"`
}

// An Step is the representation of a TOSCA Workflow Step
//
// Currently Workflows are not part of the TOSCA specification
type Step struct {
	Target             string     `yaml:"target,omitempty" json:"target,omitempty"`
	TargetRelationShip string     `yaml:"target_relationship,omitempty" json:"target_relationship,omitempty"`
	Activities         []Activity `yaml:"activities" json:"activities"`
	OnSuccess          []string   `yaml:"on_success,omitempty" json:"on_success,omitempty"`
	OperationHost      string     `yaml:"operation_host,omitempty" json:"operation_host,omitempty"`
}

// An Activity is the representation of a TOSCA Workflow Step Activity
//
// Currently Workflows are not part of the TOSCA specification
type Activity struct {
	SetState      string `yaml:"set_state,omitempty" json:"set_state,omitempty"`
	Delegate      string `yaml:"delegate,omitempty" json:"delegate,omitempty"`
	CallOperation string `yaml:"call_operation,omitempty" json:"call_operation,omitempty"`
	Inline        string `yaml:"inline,omitempty" json:"inline,omitempty"`
}
