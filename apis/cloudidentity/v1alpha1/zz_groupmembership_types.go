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

type GroupMembershipObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type GroupMembershipParameters struct {

	// The name of the Group to create this membership in.
	// +kubebuilder:validation:Required
	Group *string `json:"group" tf:"group,omitempty"`

	// EntityKey of the member.
	// +kubebuilder:validation:Optional
	PreferredMemberKey []PreferredMemberKeyParameters `json:"preferredMemberKey,omitempty" tf:"preferred_member_key,omitempty"`

	// The MembershipRoles that apply to the Membership.
	// Must not contain duplicate MembershipRoles with the same name.
	// +kubebuilder:validation:Required
	Roles []RolesParameters `json:"roles" tf:"roles,omitempty"`
}

type PreferredMemberKeyObservation struct {
}

type PreferredMemberKeyParameters struct {

	// The ID of the entity.
	//
	// For Google-managed entities, the id must be the email address of an existing
	// group or user.
	//
	// For external-identity-mapped entities, the id must be a string conforming
	// to the Identity Source's requirements.
	//
	// Must be unique within a namespace.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// The namespace in which the entity exists.
	//
	// If not specified, the EntityKey represents a Google-managed entity
	// such as a Google user or a Google Group.
	//
	// If specified, the EntityKey represents an external-identity-mapped group.
	// The namespace must correspond to an identity source created in Admin Console
	// and must be in the form of 'identitysources/{identity_source_id}'.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type RolesObservation struct {
}

type RolesParameters struct {

	// The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER. Possible values: ["OWNER", "MANAGER", "MEMBER"]
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// GroupMembershipSpec defines the desired state of GroupMembership
type GroupMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupMembershipParameters `json:"forProvider"`
}

// GroupMembershipStatus defines the observed state of GroupMembership.
type GroupMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMembership is the Schema for the GroupMemberships API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type GroupMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupMembershipSpec   `json:"spec"`
	Status            GroupMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMembershipList contains a list of GroupMemberships
type GroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupMembership `json:"items"`
}

// Repository type metadata.
var (
	GroupMembershipKind             = "GroupMembership"
	GroupMembershipGroupKind        = schema.GroupKind{Group: Group, Kind: GroupMembershipKind}.String()
	GroupMembershipKindAPIVersion   = GroupMembershipKind + "." + GroupVersion.String()
	GroupMembershipGroupVersionKind = GroupVersion.WithKind(GroupMembershipKind)
)

func init() {
	SchemeBuilder.Register(&GroupMembership{}, &GroupMembershipList{})
}
