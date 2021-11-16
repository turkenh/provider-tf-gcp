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

type AccountIamMemberConditionObservation struct {
}

type AccountIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type AccountIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type AccountIamMemberParameters struct {

	// +kubebuilder:validation:Required
	BillingAccountID *string `json:"billingAccountId" tf:"billing_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	Condition []AccountIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// AccountIamMemberSpec defines the desired state of AccountIamMember
type AccountIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountIamMemberParameters `json:"forProvider"`
}

// AccountIamMemberStatus defines the observed state of AccountIamMember.
type AccountIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccountIamMember is the Schema for the AccountIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type AccountIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountIamMemberSpec   `json:"spec"`
	Status            AccountIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountIamMemberList contains a list of AccountIamMembers
type AccountIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccountIamMember `json:"items"`
}

// Repository type metadata.
var (
	AccountIamMemberKind             = "AccountIamMember"
	AccountIamMemberGroupKind        = schema.GroupKind{Group: Group, Kind: AccountIamMemberKind}.String()
	AccountIamMemberKindAPIVersion   = AccountIamMemberKind + "." + GroupVersion.String()
	AccountIamMemberGroupVersionKind = GroupVersion.WithKind(AccountIamMemberKind)
)

func init() {
	SchemeBuilder.Register(&AccountIamMember{}, &AccountIamMemberList{})
}
