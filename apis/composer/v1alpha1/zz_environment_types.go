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

type ConfigObservation struct {
	AirflowURI *string `json:"airflowUri,omitempty" tf:"airflow_uri,omitempty"`

	DagGcsPrefix *string `json:"dagGcsPrefix,omitempty" tf:"dag_gcs_prefix,omitempty"`

	GkeCluster *string `json:"gkeCluster,omitempty" tf:"gke_cluster,omitempty"`
}

type ConfigParameters struct {

	// The configuration used for the Kubernetes Engine cluster.
	// +kubebuilder:validation:Optional
	NodeConfig []NodeConfigParameters `json:"nodeConfig,omitempty" tf:"node_config,omitempty"`

	// The number of nodes in the Kubernetes Engine cluster that will be used to run this environment. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	// +kubebuilder:validation:Optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The configuration used for the Private IP Cloud Composer environment.
	// +kubebuilder:validation:Optional
	PrivateEnvironmentConfig []PrivateEnvironmentConfigParameters `json:"privateEnvironmentConfig,omitempty" tf:"private_environment_config,omitempty"`

	// The configuration settings for software inside the environment.
	// +kubebuilder:validation:Optional
	SoftwareConfig []SoftwareConfigParameters `json:"softwareConfig,omitempty" tf:"software_config,omitempty"`
}

type EnvironmentObservation struct {
}

type EnvironmentParameters struct {

	// Configuration parameters for this environment.
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: [a-z]([-a-z0-9]*[a-z0-9])?. Label values must be between 0 and 63 characters long and must conform to the regular expression ([a-z]([-a-z0-9]*[a-z0-9])?)?. No more than 64 labels can be associated with a given environment. Both keys and values must be <= 128 bytes in size.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the environment.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The location or Compute Engine region for the environment.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type IPAllocationPolicyObservation struct {
}

type IPAllocationPolicyParameters struct {

	// The IP address range used to allocate IP addresses to pods in the cluster. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true. Set to blank to have GKE choose a range with the default size. Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. Specify either cluster_secondary_range_name or cluster_ipv4_cidr_block but not both.
	// +kubebuilder:validation:Optional
	ClusterIPv4CidrBlock *string `json:"clusterIpv4CidrBlock,omitempty" tf:"cluster_ipv4_cidr_block,omitempty"`

	// The name of the cluster's secondary range used to allocate IP addresses to pods. Specify either cluster_secondary_range_name or cluster_ipv4_cidr_block but not both. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true.
	// +kubebuilder:validation:Optional
	ClusterSecondaryRangeName *string `json:"clusterSecondaryRangeName,omitempty" tf:"cluster_secondary_range_name,omitempty"`

	// The IP address range used to allocate IP addresses in this cluster. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true. Set to blank to have GKE choose a range with the default size. Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. Specify either services_secondary_range_name or services_ipv4_cidr_block but not both.
	// +kubebuilder:validation:Optional
	ServicesIPv4CidrBlock *string `json:"servicesIpv4CidrBlock,omitempty" tf:"services_ipv4_cidr_block,omitempty"`

	// The name of the services' secondary range used to allocate IP addresses to the cluster. Specify either services_secondary_range_name or services_ipv4_cidr_block but not both. For Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*, this field is applicable only when use_ip_aliases is true.
	// +kubebuilder:validation:Optional
	ServicesSecondaryRangeName *string `json:"servicesSecondaryRangeName,omitempty" tf:"services_secondary_range_name,omitempty"`

	// Whether or not to enable Alias IPs in the GKE cluster. If true, a VPC-native cluster is created. Defaults to true if the ip_allocation_policy block is present in config. This field is only supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*. Environments in newer versions always use VPC-native GKE clusters.
	// +kubebuilder:validation:Optional
	UseIPAliases *bool `json:"useIpAliases,omitempty" tf:"use_ip_aliases,omitempty"`
}

type NodeConfigObservation struct {
}

type NodeConfigParameters struct {

	// The disk size in GB used for node VMs. Minimum size is 20GB. If unspecified, defaults to 100GB. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	// +kubebuilder:validation:Optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// Configuration for controlling how IPs are allocated in the GKE cluster. Cannot be updated.
	// +kubebuilder:validation:Optional
	IPAllocationPolicy []IPAllocationPolicyParameters `json:"ipAllocationPolicy,omitempty" tf:"ip_allocation_policy,omitempty"`

	// The Compute Engine machine type used for cluster instances, specified as a name or relative resource name. For example: "projects/{project}/zones/{zone}/machineTypes/{machineType}". Must belong to the enclosing environment's project and region/zone. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	// +kubebuilder:validation:Optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// The Compute Engine machine type used for cluster instances, specified as a name or relative resource name. For example: "projects/{project}/zones/{zone}/machineTypes/{machineType}". Must belong to the enclosing environment's project and region/zone. The network must belong to the environment's project. If unspecified, the "default" network ID in the environment's project is used. If a Custom Subnet Network is provided, subnetwork must also be provided.
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The set of Google API scopes to be made available on all node VMs. Cannot be updated. If empty, defaults to ["https://www.googleapis.com/auth/cloud-platform"]. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	// +kubebuilder:validation:Optional
	OauthScopes []*string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`

	// The Google Cloud Platform Service Account to be used by the node VMs. If a service account is not specified, the "default" Compute Engine service account is used. Cannot be updated. If given, note that the service account must have roles/composer.worker for any GCP resources created under the Cloud Composer Environment.
	// +kubebuilder:validation:Optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// The Compute Engine subnetwork to be used for machine communications, , specified as a self-link, relative resource name (e.g. "projects/{project}/regions/{region}/subnetworks/{subnetwork}"), or by name. If subnetwork is provided, network must also be provided and the subnetwork must belong to the enclosing environment's project and region.
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// The list of instance tags applied to all node VMs. Tags are used to identify valid sources or targets for network firewalls. Each tag within the list must comply with RFC1035. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Compute Engine zone in which to deploy the VMs running the Apache Airflow software, specified as the zone name or relative resource name (e.g. "projects/{project}/zones/{zone}"). Must belong to the enclosing environment's project and region. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PrivateEnvironmentConfigObservation struct {
}

type PrivateEnvironmentConfigParameters struct {

	// The CIDR block from which IP range in tenant project will be reserved for Cloud SQL. Needs to be disjoint from web_server_ipv4_cidr_block.
	// +kubebuilder:validation:Optional
	CloudSQLIPv4CidrBlock *string `json:"cloudSqlIpv4CidrBlock,omitempty" tf:"cloud_sql_ipv4_cidr_block,omitempty"`

	// If true, access to the public endpoint of the GKE cluster is denied. If this field is set to true, ip_allocation_policy.use_ip_aliases must be set to true for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	// +kubebuilder:validation:Optional
	EnablePrivateEndpoint *bool `json:"enablePrivateEndpoint,omitempty" tf:"enable_private_endpoint,omitempty"`

	// The IP range in CIDR notation to use for the hosted master network. This range is used for assigning internal IP addresses to the cluster master or set of masters and to the internal load balancer virtual IP. This range must not overlap with any other ranges in use within the cluster's network. If left blank, the default value of '172.16.0.0/28' is used.
	// +kubebuilder:validation:Optional
	MasterIPv4CidrBlock *string `json:"masterIpv4CidrBlock,omitempty" tf:"master_ipv4_cidr_block,omitempty"`

	// The CIDR block from which IP range for web server will be reserved. Needs to be disjoint from master_ipv4_cidr_block and cloud_sql_ipv4_cidr_block. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	// +kubebuilder:validation:Optional
	WebServerIPv4CidrBlock *string `json:"webServerIpv4CidrBlock,omitempty" tf:"web_server_ipv4_cidr_block,omitempty"`
}

type SoftwareConfigObservation struct {
}

type SoftwareConfigParameters struct {

	// Apache Airflow configuration properties to override. Property keys contain the section and property names, separated by a hyphen, for example "core-dags_are_paused_at_creation". Section names must not contain hyphens ("-"), opening square brackets ("["), or closing square brackets ("]"). The property name must not be empty and cannot contain "=" or ";". Section and property names cannot contain characters: "." Apache Airflow configuration property names must be written in snake_case. Property values can contain any character, and can be written in any lower/upper case format. Certain Apache Airflow configuration property values are blacklisted, and cannot be overridden.
	// +kubebuilder:validation:Optional
	AirflowConfigOverrides map[string]*string `json:"airflowConfigOverrides,omitempty" tf:"airflow_config_overrides,omitempty"`

	// Additional environment variables to provide to the Apache Airflow schedulerf, worker, and webserver processes. Environment variable names must match the regular expression [a-zA-Z_][a-zA-Z0-9_]*. They cannot specify Apache Airflow software configuration overrides (they cannot match the regular expression AIRFLOW__[A-Z0-9_]+__[A-Z0-9_]+), and they cannot match any of the following reserved names: AIRFLOW_HOME C_FORCE_ROOT CONTAINER_NAME DAGS_FOLDER GCP_PROJECT GCS_BUCKET GKE_CLUSTER_NAME SQL_DATABASE SQL_INSTANCE SQL_PASSWORD SQL_PROJECT SQL_REGION SQL_USER.
	// +kubebuilder:validation:Optional
	EnvVariables map[string]*string `json:"envVariables,omitempty" tf:"env_variables,omitempty"`

	// The version of the software running in the environment. This encapsulates both the version of Cloud Composer functionality and the version of Apache Airflow. It must match the regular expression composer-[0-9]+\.[0-9]+(\.[0-9]+)?-airflow-[0-9]+\.[0-9]+(\.[0-9]+.*)?. The Cloud Composer portion of the version is a semantic version. The portion of the image version following 'airflow-' is an official Apache Airflow repository release name. See documentation for allowed release names.
	// +kubebuilder:validation:Optional
	ImageVersion *string `json:"imageVersion,omitempty" tf:"image_version,omitempty"`

	// Custom Python Package Index (PyPI) packages to be installed in the environment. Keys refer to the lowercase package name (e.g. "numpy"). Values are the lowercase extras and version specifier (e.g. "==1.12.0", "[devel,gcp_api]", "[devel]>=1.8.2, <1.9.2"). To specify a package without pinning it to a version specifier, use the empty string as the value.
	// +kubebuilder:validation:Optional
	PypiPackages map[string]*string `json:"pypiPackages,omitempty" tf:"pypi_packages,omitempty"`

	// The major version of Python used to run the Apache Airflow scheduler, worker, and webserver processes. Can be set to '2' or '3'. If not specified, the default is '2'. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*. Environments in newer versions always use Python major version 3.
	// +kubebuilder:validation:Optional
	PythonVersion *string `json:"pythonVersion,omitempty" tf:"python_version,omitempty"`

	// The number of schedulers for Airflow. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-2.*.*.
	// +kubebuilder:validation:Optional
	SchedulerCount *int64 `json:"schedulerCount,omitempty" tf:"scheduler_count,omitempty"`
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
