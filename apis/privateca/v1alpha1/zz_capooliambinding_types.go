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

type CaPoolIamBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type CaPoolIamBindingParameters struct {

	// +kubebuilder:validation:Required
	CaPool *string `json:"caPool" tf:"ca_pool,omitempty"`

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

// CaPoolIamBindingSpec defines the desired state of CaPoolIamBinding
type CaPoolIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CaPoolIamBindingParameters `json:"forProvider"`
}

// CaPoolIamBindingStatus defines the observed state of CaPoolIamBinding.
type CaPoolIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CaPoolIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CaPoolIamBinding is the Schema for the CaPoolIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type CaPoolIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CaPoolIamBindingSpec   `json:"spec"`
	Status            CaPoolIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CaPoolIamBindingList contains a list of CaPoolIamBindings
type CaPoolIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CaPoolIamBinding `json:"items"`
}

// Repository type metadata.
var (
	CaPoolIamBindingKind             = "CaPoolIamBinding"
	CaPoolIamBindingGroupKind        = schema.GroupKind{Group: Group, Kind: CaPoolIamBindingKind}.String()
	CaPoolIamBindingKindAPIVersion   = CaPoolIamBindingKind + "." + GroupVersion.String()
	CaPoolIamBindingGroupVersionKind = GroupVersion.WithKind(CaPoolIamBindingKind)
)

func init() {
	SchemeBuilder.Register(&CaPoolIamBinding{}, &CaPoolIamBindingList{})
}
