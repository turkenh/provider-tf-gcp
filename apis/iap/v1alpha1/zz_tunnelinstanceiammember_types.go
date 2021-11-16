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

type TunnelInstanceIamMemberConditionObservation struct {
}

type TunnelInstanceIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type TunnelInstanceIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type TunnelInstanceIamMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []TunnelInstanceIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Instance *string `json:"instance" tf:"instance,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// TunnelInstanceIamMemberSpec defines the desired state of TunnelInstanceIamMember
type TunnelInstanceIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TunnelInstanceIamMemberParameters `json:"forProvider"`
}

// TunnelInstanceIamMemberStatus defines the observed state of TunnelInstanceIamMember.
type TunnelInstanceIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TunnelInstanceIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TunnelInstanceIamMember is the Schema for the TunnelInstanceIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type TunnelInstanceIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TunnelInstanceIamMemberSpec   `json:"spec"`
	Status            TunnelInstanceIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TunnelInstanceIamMemberList contains a list of TunnelInstanceIamMembers
type TunnelInstanceIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TunnelInstanceIamMember `json:"items"`
}

// Repository type metadata.
var (
	TunnelInstanceIamMemberKind             = "TunnelInstanceIamMember"
	TunnelInstanceIamMemberGroupKind        = schema.GroupKind{Group: Group, Kind: TunnelInstanceIamMemberKind}.String()
	TunnelInstanceIamMemberKindAPIVersion   = TunnelInstanceIamMemberKind + "." + GroupVersion.String()
	TunnelInstanceIamMemberGroupVersionKind = GroupVersion.WithKind(TunnelInstanceIamMemberKind)
)

func init() {
	SchemeBuilder.Register(&TunnelInstanceIamMember{}, &TunnelInstanceIamMemberList{})
}
