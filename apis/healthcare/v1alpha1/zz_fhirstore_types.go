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

type BigqueryDestinationObservation struct {
}

type BigqueryDestinationParameters struct {

	// BigQuery URI to a dataset, up to 2000 characters long, in the format bq://projectId.bqDatasetId
	// +kubebuilder:validation:Required
	DatasetURI *string `json:"datasetUri" tf:"dataset_uri,omitempty"`

	// The configuration for the exported BigQuery schema.
	// +kubebuilder:validation:Required
	SchemaConfig []SchemaConfigParameters `json:"schemaConfig" tf:"schema_config,omitempty"`
}

type FhirStoreNotificationConfigObservation struct {
}

type FhirStoreNotificationConfigParameters struct {

	// The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client.
	// PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
	// It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
	// was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
	// project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
	// Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
	// +kubebuilder:validation:Required
	PubsubTopic *string `json:"pubsubTopic" tf:"pubsub_topic,omitempty"`
}

type FhirStoreObservation struct {
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type FhirStoreParameters struct {

	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	// +kubebuilder:validation:Required
	Dataset *string `json:"dataset" tf:"dataset,omitempty"`

	// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
	// creation. The default value is false, meaning that the API will enforce referential integrity and fail the
	// requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
	// will skip referential integrity check. Consequently, operations that rely on references, such as
	// Patient.get$everything, will not return all the results if broken references exist.
	//
	// ** Changing this property may recreate the FHIR store (removing all data) **
	// +kubebuilder:validation:Optional
	DisableReferentialIntegrity *bool `json:"disableReferentialIntegrity,omitempty" tf:"disable_referential_integrity,omitempty"`

	// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
	// of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
	// versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
	// cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
	// attempts to read the historical versions.
	//
	// ** Changing this property may recreate the FHIR store (removing all data) **
	// +kubebuilder:validation:Optional
	DisableResourceVersioning *bool `json:"disableResourceVersioning,omitempty" tf:"disable_resource_versioning,omitempty"`

	// Whether to allow the bulk import API to accept history bundles and directly insert historical resource
	// versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
	// occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
	// will fail with an error.
	//
	// ** Changing this property may recreate the FHIR store (removing all data) **
	//
	// ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
	// +kubebuilder:validation:Optional
	EnableHistoryImport *bool `json:"enableHistoryImport,omitempty" tf:"enable_history_import,omitempty"`

	// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
	// operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
	// the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
	// logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
	// identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
	// notifications.
	// +kubebuilder:validation:Optional
	EnableUpdateCreate *bool `json:"enableUpdateCreate,omitempty" tf:"enable_update_create,omitempty"`

	// User-supplied key-value pairs used to organize FHIR stores.
	//
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
	// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//
	// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
	// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//
	// No more than 64 labels can be associated with a given store.
	//
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The resource name for the FhirStore.
	//
	// ** Changing this property may recreate the FHIR store (removing all data) **
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A nested object resource
	// +kubebuilder:validation:Optional
	NotificationConfig []FhirStoreNotificationConfigParameters `json:"notificationConfig,omitempty" tf:"notification_config,omitempty"`

	// A list of streaming configs that configure the destinations of streaming export for every resource mutation in
	// this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
	// resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
	// from the list, the server stops streaming to that location. Before adding a new config, you must add the required
	// bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
	// the order of dozens of seconds) is expected before the results show up in the streaming destination.
	// +kubebuilder:validation:Optional
	StreamConfigs []StreamConfigsParameters `json:"streamConfigs,omitempty" tf:"stream_configs,omitempty"`

	// The FHIR specification version. Possible values: ["DSTU2", "STU3", "R4"]
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type SchemaConfigObservation struct {
}

type SchemaConfigParameters struct {

	// The depth for all recursive structures in the output analytics schema. For example, concept in the CodeSystem
	// resource is a recursive structure; when the depth is 2, the CodeSystem table will have a column called
	// concept.concept but not concept.concept.concept. If not specified or set to 0, the server will use the default
	// value 2. The maximum depth allowed is 5.
	// +kubebuilder:validation:Required
	RecursiveStructureDepth *int64 `json:"recursiveStructureDepth" tf:"recursive_structure_depth,omitempty"`

	// Specifies the output schema type. Only ANALYTICS is supported at this time.
	// * ANALYTICS: Analytics schema defined by the FHIR community.
	// See https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md. Default value: "ANALYTICS" Possible values: ["ANALYTICS"]
	// +kubebuilder:validation:Optional
	SchemaType *string `json:"schemaType,omitempty" tf:"schema_type,omitempty"`
}

type StreamConfigsObservation struct {
}

type StreamConfigsParameters struct {

	// The destination BigQuery structure that contains both the dataset location and corresponding schema config.
	// The output is organized in one table per resource type. The server reuses the existing tables (if any) that
	// are named after the resource types, e.g. "Patient", "Observation". When there is no existing table for a given
	// resource type, the server attempts to create one.
	// See the [streaming config reference](https://cloud.google.com/healthcare/docs/reference/rest/v1beta1/projects.locations.datasets.fhirStores#streamconfig) for more details.
	// +kubebuilder:validation:Required
	BigqueryDestination []BigqueryDestinationParameters `json:"bigqueryDestination" tf:"bigquery_destination,omitempty"`

	// Supply a FHIR resource type (such as "Patient" or "Observation"). See
	// https://www.hl7.org/fhir/valueset-resource-types.html for a list of all FHIR resource types. The server treats
	// an empty list as an intent to stream all the supported resource types in this FHIR store.
	// +kubebuilder:validation:Optional
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

// FhirStoreSpec defines the desired state of FhirStore
type FhirStoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FhirStoreParameters `json:"forProvider"`
}

// FhirStoreStatus defines the observed state of FhirStore.
type FhirStoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FhirStoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FhirStore is the Schema for the FhirStores API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type FhirStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FhirStoreSpec   `json:"spec"`
	Status            FhirStoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FhirStoreList contains a list of FhirStores
type FhirStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FhirStore `json:"items"`
}

// Repository type metadata.
var (
	FhirStoreKind             = "FhirStore"
	FhirStoreGroupKind        = schema.GroupKind{Group: Group, Kind: FhirStoreKind}.String()
	FhirStoreKindAPIVersion   = FhirStoreKind + "." + GroupVersion.String()
	FhirStoreGroupVersionKind = GroupVersion.WithKind(FhirStoreKind)
)

func init() {
	SchemeBuilder.Register(&FhirStore{}, &FhirStoreList{})
}
