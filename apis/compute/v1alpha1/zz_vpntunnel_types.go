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

type VpnTunnelObservation struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	DetailedStatus *string `json:"detailedStatus,omitempty" tf:"detailed_status,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	SharedSecretHash *string `json:"sharedSecretHash,omitempty" tf:"shared_secret_hash,omitempty"`

	TunnelID *string `json:"tunnelId,omitempty" tf:"tunnel_id,omitempty"`
}

type VpnTunnelParameters struct {

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// IKE protocol version to use when establishing the VPN tunnel with
	// peer VPN gateway.
	// Acceptable IKE versions are 1 or 2. Default version is 2.
	// +kubebuilder:validation:Optional
	IkeVersion *int64 `json:"ikeVersion,omitempty" tf:"ike_version,omitempty"`

	// Local traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example '192.168.0.0/16'. The ranges should be disjoint.
	// Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	LocalTrafficSelector []*string `json:"localTrafficSelector,omitempty" tf:"local_traffic_selector,omitempty"`

	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63
	// characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character
	// must be a lowercase letter, and all following characters must
	// be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// URL of the peer side external VPN gateway to which this VPN tunnel is connected.
	// +kubebuilder:validation:Optional
	PeerExternalGateway *string `json:"peerExternalGateway,omitempty" tf:"peer_external_gateway,omitempty"`

	// The interface ID of the external VPN gateway to which this VPN tunnel is connected.
	// +kubebuilder:validation:Optional
	PeerExternalGatewayInterface *int64 `json:"peerExternalGatewayInterface,omitempty" tf:"peer_external_gateway_interface,omitempty"`

	// URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
	// If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
	// ID in the peer GCP VPN gateway.
	// This field must reference a 'google_compute_ha_vpn_gateway' resource.
	// +kubebuilder:validation:Optional
	PeerGcpGateway *string `json:"peerGcpGateway,omitempty" tf:"peer_gcp_gateway,omitempty"`

	// IP address of the peer VPN gateway. Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	PeerIP *string `json:"peerIp,omitempty" tf:"peer_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Remote traffic selector to use when establishing the VPN tunnel with
	// peer VPN gateway. The value should be a CIDR formatted string,
	// for example '192.168.0.0/16'. The ranges should be disjoint.
	// Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	RemoteTrafficSelector []*string `json:"remoteTrafficSelector,omitempty" tf:"remote_traffic_selector,omitempty"`

	// URL of router resource to be used for dynamic routing.
	// +kubebuilder:validation:Optional
	Router *string `json:"router,omitempty" tf:"router,omitempty"`

	// Shared secret used to set the secure session between the Cloud VPN
	// gateway and the peer VPN gateway.
	// +kubebuilder:validation:Required
	SharedSecretSecretRef v1.SecretKeySelector `json:"sharedSecretSecretRef" tf:"-"`

	// URL of the Target VPN gateway with which this VPN tunnel is
	// associated.
	// +kubebuilder:validation:Optional
	TargetVpnGateway *string `json:"targetVpnGateway,omitempty" tf:"target_vpn_gateway,omitempty"`

	// URL of the VPN gateway with which this VPN tunnel is associated.
	// This must be used if a High Availability VPN gateway resource is created.
	// This field must reference a 'google_compute_ha_vpn_gateway' resource.
	// +kubebuilder:validation:Optional
	VpnGateway *string `json:"vpnGateway,omitempty" tf:"vpn_gateway,omitempty"`

	// The interface ID of the VPN gateway with which this VPN tunnel is associated.
	// +kubebuilder:validation:Optional
	VpnGatewayInterface *int64 `json:"vpnGatewayInterface,omitempty" tf:"vpn_gateway_interface,omitempty"`
}

// VpnTunnelSpec defines the desired state of VpnTunnel
type VpnTunnelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VpnTunnelParameters `json:"forProvider"`
}

// VpnTunnelStatus defines the observed state of VpnTunnel.
type VpnTunnelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VpnTunnelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpnTunnel is the Schema for the VpnTunnels API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type VpnTunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnTunnelSpec   `json:"spec"`
	Status            VpnTunnelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnTunnelList contains a list of VpnTunnels
type VpnTunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnTunnel `json:"items"`
}

// Repository type metadata.
var (
	VpnTunnelKind             = "VpnTunnel"
	VpnTunnelGroupKind        = schema.GroupKind{Group: Group, Kind: VpnTunnelKind}.String()
	VpnTunnelKindAPIVersion   = VpnTunnelKind + "." + GroupVersion.String()
	VpnTunnelGroupVersionKind = GroupVersion.WithKind(VpnTunnelKind)
)

func init() {
	SchemeBuilder.Register(&VpnTunnel{}, &VpnTunnelList{})
}
