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

type RegionSslCertificateObservation struct {
	CertificateID *int64 `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type RegionSslCertificateParameters struct {

	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	// +kubebuilder:validation:Required
	CertificateSecretRef v1.SecretKeySelector `json:"certificateSecretRef" tf:"-"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	//
	// These are in the same namespace as the managed SSL certificates.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// The write-only private key in PEM format.
	// +kubebuilder:validation:Required
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created regional ssl certificate should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// RegionSslCertificateSpec defines the desired state of RegionSslCertificate
type RegionSslCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionSslCertificateParameters `json:"forProvider"`
}

// RegionSslCertificateStatus defines the observed state of RegionSslCertificate.
type RegionSslCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionSslCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionSslCertificate is the Schema for the RegionSslCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type RegionSslCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionSslCertificateSpec   `json:"spec"`
	Status            RegionSslCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionSslCertificateList contains a list of RegionSslCertificates
type RegionSslCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionSslCertificate `json:"items"`
}

// Repository type metadata.
var (
	RegionSslCertificateKind             = "RegionSslCertificate"
	RegionSslCertificateGroupKind        = schema.GroupKind{Group: Group, Kind: RegionSslCertificateKind}.String()
	RegionSslCertificateKindAPIVersion   = RegionSslCertificateKind + "." + GroupVersion.String()
	RegionSslCertificateGroupVersionKind = GroupVersion.WithKind(RegionSslCertificateKind)
)

func init() {
	SchemeBuilder.Register(&RegionSslCertificate{}, &RegionSslCertificateList{})
}
