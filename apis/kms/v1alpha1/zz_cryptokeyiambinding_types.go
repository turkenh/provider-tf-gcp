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

type CryptoKeyIamBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type CryptoKeyIamBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	CryptoKeyID *string `json:"cryptoKeyId" tf:"crypto_key_id,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// CryptoKeyIamBindingSpec defines the desired state of CryptoKeyIamBinding
type CryptoKeyIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CryptoKeyIamBindingParameters `json:"forProvider"`
}

// CryptoKeyIamBindingStatus defines the observed state of CryptoKeyIamBinding.
type CryptoKeyIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CryptoKeyIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CryptoKeyIamBinding is the Schema for the CryptoKeyIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type CryptoKeyIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CryptoKeyIamBindingSpec   `json:"spec"`
	Status            CryptoKeyIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CryptoKeyIamBindingList contains a list of CryptoKeyIamBindings
type CryptoKeyIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CryptoKeyIamBinding `json:"items"`
}

// Repository type metadata.
var (
	CryptoKeyIamBindingKind             = "CryptoKeyIamBinding"
	CryptoKeyIamBindingGroupKind        = schema.GroupKind{Group: Group, Kind: CryptoKeyIamBindingKind}.String()
	CryptoKeyIamBindingKindAPIVersion   = CryptoKeyIamBindingKind + "." + GroupVersion.String()
	CryptoKeyIamBindingGroupVersionKind = GroupVersion.WithKind(CryptoKeyIamBindingKind)
)

func init() {
	SchemeBuilder.Register(&CryptoKeyIamBinding{}, &CryptoKeyIamBindingList{})
}
