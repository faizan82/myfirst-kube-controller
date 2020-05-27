/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

const (
	// PhasePending indicates no action on resource has been performed yet
	PhasePending = "PENDING"
	// PhaseRunning denotes an action being performed through cr
	PhaseRunning = "RUNNING"
	// PhaseDone denotes completion
	PhaseDone = "DONE"
)

// At is a specification for a At resource
type At struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AtSpec   `json:"spec"`
	Status AtStatus `json:"status"`
}

// AtSpec is the spec for a At resource
type AtSpec struct {
	Schedule string `json:"schedule"`
	Command  string `json:"command"`
}

// AtStatus is the status for a At resource
type AtStatus struct {
	Phase string `json:"phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AtList is a list of At resources
type AtList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []At `json:"items"`
}
