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

type ContainerImageObservation struct {
}

type ContainerImageParameters struct {

	// The path to the container image repository.
	// For example: gcr.io/{project_id}/{imageName}
	// +kubebuilder:validation:Required
	Repository *string `json:"repository" tf:"repository,omitempty"`

	// The tag of the container image. If not specified, this defaults to the latest tag.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type EnvironmentObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`
}

type EnvironmentParameters struct {

	// Use a container image to start the notebook instance.
	// +kubebuilder:validation:Optional
	ContainerImage []ContainerImageParameters `json:"containerImage,omitempty" tf:"container_image,omitempty"`

	// A brief description of this environment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Display name of this environment for the UI.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// A reference to the zone where the machine resides.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name specified for the Environment instance.
	// Format: projects/{project_id}/locations/{location}/environments/{environmentId}
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	// The path must be a URL or Cloud Storage path. Example: "gs://path-to-file/file-name"
	// +kubebuilder:validation:Optional
	PostStartupScript *string `json:"postStartupScript,omitempty" tf:"post_startup_script,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Use a Compute Engine VM image to start the notebook instance.
	// +kubebuilder:validation:Optional
	VMImage []VMImageParameters `json:"vmImage,omitempty" tf:"vm_image,omitempty"`
}

type VMImageObservation struct {
}

type VMImageParameters struct {

	// Use this VM image family to find the image; the newest image in this family will be used.
	// +kubebuilder:validation:Optional
	ImageFamily *string `json:"imageFamily,omitempty" tf:"image_family,omitempty"`

	// Use VM image name to find the image.
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// The name of the Google Cloud project that this VM image belongs to.
	// Format: projects/{project_id}
	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`
}

// EnvironmentSpec defines the desired state of Environment
type EnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentParameters `json:"forProvider"`
}

// EnvironmentStatus defines the observed state of Environment.
type EnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Environment is the Schema for the Environments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EnvironmentSpec   `json:"spec"`
	Status            EnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentList contains a list of Environments
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentKind             = "Environment"
	EnvironmentGroupKind        = schema.GroupKind{Group: Group, Kind: EnvironmentKind}.String()
	EnvironmentKindAPIVersion   = EnvironmentKind + "." + GroupVersion.String()
	EnvironmentGroupVersionKind = GroupVersion.WithKind(EnvironmentKind)
)

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
