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

type GameServerDeploymentObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GameServerDeploymentParameters struct {

	// A unique id for the deployment.
	// +kubebuilder:validation:Required
	DeploymentID *string `json:"deploymentId" tf:"deployment_id,omitempty"`

	// Human readable description of the game server deployment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The labels associated with this game server deployment. Each label is a
	// key-value pair.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Location of the Deployment.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// GameServerDeploymentSpec defines the desired state of GameServerDeployment
type GameServerDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GameServerDeploymentParameters `json:"forProvider"`
}

// GameServerDeploymentStatus defines the observed state of GameServerDeployment.
type GameServerDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GameServerDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GameServerDeployment is the Schema for the GameServerDeployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type GameServerDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameServerDeploymentSpec   `json:"spec"`
	Status            GameServerDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameServerDeploymentList contains a list of GameServerDeployments
type GameServerDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameServerDeployment `json:"items"`
}

// Repository type metadata.
var (
	GameServerDeploymentKind             = "GameServerDeployment"
	GameServerDeploymentGroupKind        = schema.GroupKind{Group: Group, Kind: GameServerDeploymentKind}.String()
	GameServerDeploymentKindAPIVersion   = GameServerDeploymentKind + "." + GroupVersion.String()
	GameServerDeploymentGroupVersionKind = GroupVersion.WithKind(GameServerDeploymentKind)
)

func init() {
	SchemeBuilder.Register(&GameServerDeployment{}, &GameServerDeploymentList{})
}
