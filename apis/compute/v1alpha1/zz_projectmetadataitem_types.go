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

type ProjectMetadataItemObservation struct {
}

type ProjectMetadataItemParameters struct {

	// The metadata key to set.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The value to set for the given metadata key.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// ProjectMetadataItemSpec defines the desired state of ProjectMetadataItem
type ProjectMetadataItemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectMetadataItemParameters `json:"forProvider"`
}

// ProjectMetadataItemStatus defines the observed state of ProjectMetadataItem.
type ProjectMetadataItemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectMetadataItemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectMetadataItem is the Schema for the ProjectMetadataItems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type ProjectMetadataItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectMetadataItemSpec   `json:"spec"`
	Status            ProjectMetadataItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectMetadataItemList contains a list of ProjectMetadataItems
type ProjectMetadataItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectMetadataItem `json:"items"`
}

// Repository type metadata.
var (
	ProjectMetadataItemKind             = "ProjectMetadataItem"
	ProjectMetadataItemGroupKind        = schema.GroupKind{Group: Group, Kind: ProjectMetadataItemKind}.String()
	ProjectMetadataItemKindAPIVersion   = ProjectMetadataItemKind + "." + GroupVersion.String()
	ProjectMetadataItemGroupVersionKind = GroupVersion.WithKind(ProjectMetadataItemKind)
)

func init() {
	SchemeBuilder.Register(&ProjectMetadataItem{}, &ProjectMetadataItemList{})
}
