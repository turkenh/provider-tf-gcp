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

type AppProfileObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AppProfileParameters struct {

	// The unique name of the app profile in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	// +kubebuilder:validation:Required
	AppProfileID *string `json:"appProfileId" tf:"app_profile_id,omitempty"`

	// Long form description of the use case for this app profile.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If true, ignore safety checks when deleting/updating the app profile.
	// +kubebuilder:validation:Optional
	IgnoreWarnings *bool `json:"ignoreWarnings,omitempty" tf:"ignore_warnings,omitempty"`

	// The name of the instance to create the app profile within.
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
	// in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
	// consistency to improve availability.
	// +kubebuilder:validation:Optional
	MultiClusterRoutingUseAny *bool `json:"multiClusterRoutingUseAny,omitempty" tf:"multi_cluster_routing_use_any,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Use a single-cluster routing policy.
	// +kubebuilder:validation:Optional
	SingleClusterRouting []SingleClusterRoutingParameters `json:"singleClusterRouting,omitempty" tf:"single_cluster_routing,omitempty"`
}

type SingleClusterRoutingObservation struct {
}

type SingleClusterRoutingParameters struct {

	// If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
	// It is unsafe to send these requests to the same table/row/column in multiple clusters.
	// +kubebuilder:validation:Optional
	AllowTransactionalWrites *bool `json:"allowTransactionalWrites,omitempty" tf:"allow_transactional_writes,omitempty"`

	// The cluster to which read/write requests should be routed.
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`
}

// AppProfileSpec defines the desired state of AppProfile
type AppProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppProfileParameters `json:"forProvider"`
}

// AppProfileStatus defines the observed state of AppProfile.
type AppProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppProfile is the Schema for the AppProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type AppProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppProfileSpec   `json:"spec"`
	Status            AppProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppProfileList contains a list of AppProfiles
type AppProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppProfile `json:"items"`
}

// Repository type metadata.
var (
	AppProfileKind             = "AppProfile"
	AppProfileGroupKind        = schema.GroupKind{Group: Group, Kind: AppProfileKind}.String()
	AppProfileKindAPIVersion   = AppProfileKind + "." + GroupVersion.String()
	AppProfileGroupVersionKind = GroupVersion.WithKind(AppProfileKind)
)

func init() {
	SchemeBuilder.Register(&AppProfile{}, &AppProfileList{})
}
