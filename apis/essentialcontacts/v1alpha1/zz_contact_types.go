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

type ContactObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ContactParameters struct {

	// The email address to send notifications to. This does not need to be a Google account.
	// +kubebuilder:validation:Required
	Email *string `json:"email" tf:"email,omitempty"`

	// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
	// +kubebuilder:validation:Required
	LanguageTag *string `json:"languageTag" tf:"language_tag,omitempty"`

	// The categories of notifications that the contact will receive communications for.
	// +kubebuilder:validation:Required
	NotificationCategorySubscriptions []*string `json:"notificationCategorySubscriptions" tf:"notification_category_subscriptions,omitempty"`

	// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
	// +kubebuilder:validation:Required
	Parent *string `json:"parent" tf:"parent,omitempty"`
}

// ContactSpec defines the desired state of Contact
type ContactSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContactParameters `json:"forProvider"`
}

// ContactStatus defines the observed state of Contact.
type ContactStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Contact is the Schema for the Contacts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type Contact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContactSpec   `json:"spec"`
	Status            ContactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContactList contains a list of Contacts
type ContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Contact `json:"items"`
}

// Repository type metadata.
var (
	ContactKind             = "Contact"
	ContactGroupKind        = schema.GroupKind{Group: Group, Kind: ContactKind}.String()
	ContactKindAPIVersion   = ContactKind + "." + GroupVersion.String()
	ContactGroupVersionKind = GroupVersion.WithKind(ContactKind)
)

func init() {
	SchemeBuilder.Register(&Contact{}, &ContactList{})
}
