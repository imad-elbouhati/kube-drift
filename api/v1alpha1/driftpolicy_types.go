/*
Copyright 2025.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type TargetReference struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type GitSyncSource struct {
	// +kubebuilder:validation:'^https?:\/\/.+$'
	RepoURL  string `json:"repoUrl"`
	Path     string `json:"path"`
	Revision string `json:"revision"`
}

// DriftPolicySpec defines the desired state of DriftPolicy.
type DriftPolicySpec struct {
	// Reference to the Kubernetes resource to monitor
	TargetRef TargetReference `json:"targetRef"`

	// Git source of the truth
	SyncSource GitSyncSource `json:"syncSource"`
}

// DriftPolicyStatus defines the observed state of DriftPolicy.
// +kubebuilder:subresource:status
type DriftPolicyStatus struct {
	// LastRevision is the Git revision (commit hash or tag) last processed for drift check.
	LastRevision string `json:"lastRevision,omitempty"`

	DriftDetected bool `json:"driftDetected,omitempty"`
	
	LastChecked metav1.Time `json:"lastChecked,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DriftPolicy is the Schema for the driftpolicies API.
type DriftPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DriftPolicySpec   `json:"spec,omitempty"`
	Status DriftPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DriftPolicyList contains a list of DriftPolicy.
type DriftPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DriftPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DriftPolicy{}, &DriftPolicyList{})
}
