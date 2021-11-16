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

type Hl7V2StoreIamPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type Hl7V2StoreIamPolicyParameters struct {

	// +kubebuilder:validation:Required
	Hl7V2StoreID *string `json:"hl7V2StoreId" tf:"hl7_v2_store_id,omitempty"`

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`
}

// Hl7V2StoreIamPolicySpec defines the desired state of Hl7V2StoreIamPolicy
type Hl7V2StoreIamPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     Hl7V2StoreIamPolicyParameters `json:"forProvider"`
}

// Hl7V2StoreIamPolicyStatus defines the observed state of Hl7V2StoreIamPolicy.
type Hl7V2StoreIamPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        Hl7V2StoreIamPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Hl7V2StoreIamPolicy is the Schema for the Hl7V2StoreIamPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type Hl7V2StoreIamPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Hl7V2StoreIamPolicySpec   `json:"spec"`
	Status            Hl7V2StoreIamPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Hl7V2StoreIamPolicyList contains a list of Hl7V2StoreIamPolicys
type Hl7V2StoreIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hl7V2StoreIamPolicy `json:"items"`
}

// Repository type metadata.
var (
	Hl7V2StoreIamPolicyKind             = "Hl7V2StoreIamPolicy"
	Hl7V2StoreIamPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: Hl7V2StoreIamPolicyKind}.String()
	Hl7V2StoreIamPolicyKindAPIVersion   = Hl7V2StoreIamPolicyKind + "." + GroupVersion.String()
	Hl7V2StoreIamPolicyGroupVersionKind = GroupVersion.WithKind(Hl7V2StoreIamPolicyKind)
)

func init() {
	SchemeBuilder.Register(&Hl7V2StoreIamPolicy{}, &Hl7V2StoreIamPolicyList{})
}
