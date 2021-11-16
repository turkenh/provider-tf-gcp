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

type ProjectIAMAuditConfigAuditLogConfigObservation struct {
}

type ProjectIAMAuditConfigAuditLogConfigParameters struct {

	// Identities that do not cause logging for this type of permission. Each entry can have one of the following values:user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com. serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com. group:{emailid}: An email address that represents a Google group. For example, admins@example.com. domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	// +kubebuilder:validation:Optional
	ExemptedMembers []*string `json:"exemptedMembers,omitempty" tf:"exempted_members,omitempty"`

	// Permission type for which logging is to be configured. Must be one of DATA_READ, DATA_WRITE, or ADMIN_READ.
	// +kubebuilder:validation:Required
	LogType *string `json:"logType" tf:"log_type,omitempty"`
}

type ProjectIAMAuditConfigObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type ProjectIAMAuditConfigParameters struct {

	// The configuration for logging of each type of permission. This can be specified multiple times.
	// +kubebuilder:validation:Required
	AuditLogConfig []ProjectIAMAuditConfigAuditLogConfigParameters `json:"auditLogConfig" tf:"audit_log_config,omitempty"`

	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// Service which will be enabled for audit logging. The special value allServices covers all services.
	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

// ProjectIAMAuditConfigSpec defines the desired state of ProjectIAMAuditConfig
type ProjectIAMAuditConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectIAMAuditConfigParameters `json:"forProvider"`
}

// ProjectIAMAuditConfigStatus defines the observed state of ProjectIAMAuditConfig.
type ProjectIAMAuditConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectIAMAuditConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMAuditConfig is the Schema for the ProjectIAMAuditConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type ProjectIAMAuditConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectIAMAuditConfigSpec   `json:"spec"`
	Status            ProjectIAMAuditConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMAuditConfigList contains a list of ProjectIAMAuditConfigs
type ProjectIAMAuditConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectIAMAuditConfig `json:"items"`
}

// Repository type metadata.
var (
	ProjectIAMAuditConfigKind             = "ProjectIAMAuditConfig"
	ProjectIAMAuditConfigGroupKind        = schema.GroupKind{Group: Group, Kind: ProjectIAMAuditConfigKind}.String()
	ProjectIAMAuditConfigKindAPIVersion   = ProjectIAMAuditConfigKind + "." + GroupVersion.String()
	ProjectIAMAuditConfigGroupVersionKind = GroupVersion.WithKind(ProjectIAMAuditConfigKind)
)

func init() {
	SchemeBuilder.Register(&ProjectIAMAuditConfig{}, &ProjectIAMAuditConfigList{})
}
