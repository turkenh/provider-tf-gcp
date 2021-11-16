/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceIamPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type InstanceIamPolicyParameters struct {

	// +kubebuilder:validation:Required
	Instance *string `json:"instance" tf:"instance,omitempty"`

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// InstanceIamPolicySpec defines the desired state of InstanceIamPolicy
type InstanceIamPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceIamPolicyParameters `json:"forProvider"`
}

// InstanceIamPolicyStatus defines the observed state of InstanceIamPolicy.
type InstanceIamPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceIamPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIamPolicy is the Schema for the InstanceIamPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type InstanceIamPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceIamPolicySpec   `json:"spec"`
	Status            InstanceIamPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIamPolicyList contains a list of InstanceIamPolicys
type InstanceIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceIamPolicy `json:"items"`
}

// Repository type metadata.
var (
	InstanceIamPolicyKind             = "InstanceIamPolicy"
	InstanceIamPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: InstanceIamPolicyKind}.String()
	InstanceIamPolicyKindAPIVersion   = InstanceIamPolicyKind + "." + GroupVersion.String()
	InstanceIamPolicyGroupVersionKind = GroupVersion.WithKind(InstanceIamPolicyKind)
)

func init() {
	SchemeBuilder.Register(&InstanceIamPolicy{}, &InstanceIamPolicyList{})
}
