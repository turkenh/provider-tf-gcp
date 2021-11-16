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

type KeyRingIamBindingConditionObservation struct {
}

type KeyRingIamBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type KeyRingIamBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type KeyRingIamBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []KeyRingIamBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	KeyRingID *string `json:"keyRingId" tf:"key_ring_id,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// KeyRingIamBindingSpec defines the desired state of KeyRingIamBinding
type KeyRingIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyRingIamBindingParameters `json:"forProvider"`
}

// KeyRingIamBindingStatus defines the observed state of KeyRingIamBinding.
type KeyRingIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyRingIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeyRingIamBinding is the Schema for the KeyRingIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type KeyRingIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyRingIamBindingSpec   `json:"spec"`
	Status            KeyRingIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyRingIamBindingList contains a list of KeyRingIamBindings
type KeyRingIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyRingIamBinding `json:"items"`
}

// Repository type metadata.
var (
	KeyRingIamBindingKind             = "KeyRingIamBinding"
	KeyRingIamBindingGroupKind        = schema.GroupKind{Group: Group, Kind: KeyRingIamBindingKind}.String()
	KeyRingIamBindingKindAPIVersion   = KeyRingIamBindingKind + "." + GroupVersion.String()
	KeyRingIamBindingGroupVersionKind = GroupVersion.WithKind(KeyRingIamBindingKind)
)

func init() {
	SchemeBuilder.Register(&KeyRingIamBinding{}, &KeyRingIamBindingList{})
}
