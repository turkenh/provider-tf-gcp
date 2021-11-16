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

type DailyScheduleObservation struct {
}

type DailyScheduleParameters struct {

	// The number of days between snapshots.
	// +kubebuilder:validation:Required
	DaysInCycle *int64 `json:"daysInCycle" tf:"days_in_cycle,omitempty"`

	// This must be in UTC format that resolves to one of
	// 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example,
	// both 13:00-5 and 08:00 are valid.
	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type DayOfWeeksObservation struct {
}

type DayOfWeeksParameters struct {

	// The day of the week to create the snapshot. e.g. MONDAY Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	// +kubebuilder:validation:Required
	Day *string `json:"day" tf:"day,omitempty"`

	// Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type GroupPlacementPolicyObservation struct {
}

type GroupPlacementPolicyParameters struct {

	// The number of availability domains instances will be spread across. If two instances are in different
	// availability domain, they will not be put in the same low latency network
	// +kubebuilder:validation:Optional
	AvailabilityDomainCount *int64 `json:"availabilityDomainCount,omitempty" tf:"availability_domain_count,omitempty"`

	// Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
	// Specify 'COLLOCATED' to enable collocation. Can only be specified with 'vm_count'. If compute instances are created
	// with a COLLOCATED policy, then exactly 'vm_count' instances must be created at the same time with the resource policy
	// attached. Possible values: ["COLLOCATED"]
	// +kubebuilder:validation:Optional
	Collocation *string `json:"collocation,omitempty" tf:"collocation,omitempty"`

	// Number of vms in this placement group.
	// +kubebuilder:validation:Optional
	VMCount *int64 `json:"vmCount,omitempty" tf:"vm_count,omitempty"`
}

type HourlyScheduleObservation struct {
}

type HourlyScheduleParameters struct {

	// The number of hours between snapshots.
	// +kubebuilder:validation:Required
	HoursInCycle *int64 `json:"hoursInCycle" tf:"hours_in_cycle,omitempty"`

	// Time within the window to start the operations.
	// It must be in an hourly format "HH:MM",
	// where HH : [00-23] and MM : [00] GMT.
	// eg: 21:00
	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type InstanceSchedulePolicyObservation struct {
}

type InstanceSchedulePolicyParameters struct {

	// The expiration time of the schedule. The timestamp is an RFC3339 string.
	// +kubebuilder:validation:Optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// The start time of the schedule. The timestamp is an RFC3339 string.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone to be used in interpreting the schedule. The value of this field must be a time zone name
	// from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	// +kubebuilder:validation:Required
	TimeZone *string `json:"timeZone" tf:"time_zone,omitempty"`

	// Specifies the schedule for starting instances.
	// +kubebuilder:validation:Optional
	VMStartSchedule []VMStartScheduleParameters `json:"vmStartSchedule,omitempty" tf:"vm_start_schedule,omitempty"`

	// Specifies the schedule for stopping instances.
	// +kubebuilder:validation:Optional
	VMStopSchedule []VMStopScheduleParameters `json:"vmStopSchedule,omitempty" tf:"vm_stop_schedule,omitempty"`
}

type ResourcePolicyObservation struct {
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ResourcePolicyParameters struct {

	// An optional description of this resource. Provide this property when you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Resource policy for instances used for placement configuration.
	// +kubebuilder:validation:Optional
	GroupPlacementPolicy []GroupPlacementPolicyParameters `json:"groupPlacementPolicy,omitempty" tf:"group_placement_policy,omitempty"`

	// Resource policy for scheduling instance operations.
	// +kubebuilder:validation:Optional
	InstanceSchedulePolicy []InstanceSchedulePolicyParameters `json:"instanceSchedulePolicy,omitempty" tf:"instance_schedule_policy,omitempty"`

	// The name of the resource, provided by the client when initially creating
	// the resource. The resource name must be 1-63 characters long, and comply
	// with RFC1035. Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the
	// first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character,
	// which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where resource policy resides.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Policy for creating snapshots of persistent disks.
	// +kubebuilder:validation:Optional
	SnapshotSchedulePolicy []SnapshotSchedulePolicyParameters `json:"snapshotSchedulePolicy,omitempty" tf:"snapshot_schedule_policy,omitempty"`
}

type RetentionPolicyObservation struct {
}

type RetentionPolicyParameters struct {

	// Maximum age of the snapshot that is allowed to be kept.
	// +kubebuilder:validation:Required
	MaxRetentionDays *int64 `json:"maxRetentionDays" tf:"max_retention_days,omitempty"`

	// Specifies the behavior to apply to scheduled snapshots when
	// the source disk is deleted. Default value: "KEEP_AUTO_SNAPSHOTS" Possible values: ["KEEP_AUTO_SNAPSHOTS", "APPLY_RETENTION_POLICY"]
	// +kubebuilder:validation:Optional
	OnSourceDiskDelete *string `json:"onSourceDiskDelete,omitempty" tf:"on_source_disk_delete,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// The policy will execute every nth day at the specified time.
	// +kubebuilder:validation:Optional
	DailySchedule []DailyScheduleParameters `json:"dailySchedule,omitempty" tf:"daily_schedule,omitempty"`

	// The policy will execute every nth hour starting at the specified time.
	// +kubebuilder:validation:Optional
	HourlySchedule []HourlyScheduleParameters `json:"hourlySchedule,omitempty" tf:"hourly_schedule,omitempty"`

	// Allows specifying a snapshot time for each day of the week.
	// +kubebuilder:validation:Optional
	WeeklySchedule []WeeklyScheduleParameters `json:"weeklySchedule,omitempty" tf:"weekly_schedule,omitempty"`
}

type SnapshotPropertiesObservation struct {
}

type SnapshotPropertiesParameters struct {

	// Whether to perform a 'guest aware' snapshot.
	// +kubebuilder:validation:Optional
	GuestFlush *bool `json:"guestFlush,omitempty" tf:"guest_flush,omitempty"`

	// A set of key-value pairs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Cloud Storage bucket location to store the auto snapshot
	// (regional or multi-regional)
	// +kubebuilder:validation:Optional
	StorageLocations []*string `json:"storageLocations,omitempty" tf:"storage_locations,omitempty"`
}

type SnapshotSchedulePolicyObservation struct {
}

type SnapshotSchedulePolicyParameters struct {

	// Retention policy applied to snapshots created by this resource policy.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Contains one of an 'hourlySchedule', 'dailySchedule', or 'weeklySchedule'.
	// +kubebuilder:validation:Required
	Schedule []ScheduleParameters `json:"schedule" tf:"schedule,omitempty"`

	// Properties with which the snapshots are created, such as labels.
	// +kubebuilder:validation:Optional
	SnapshotProperties []SnapshotPropertiesParameters `json:"snapshotProperties,omitempty" tf:"snapshot_properties,omitempty"`
}

type VMStartScheduleObservation struct {
}

type VMStartScheduleParameters struct {

	// Specifies the frequency for the operation, using the unix-cron format.
	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type VMStopScheduleObservation struct {
}

type VMStopScheduleParameters struct {

	// Specifies the frequency for the operation, using the unix-cron format.
	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type WeeklyScheduleObservation struct {
}

type WeeklyScheduleParameters struct {

	// May contain up to seven (one for each day of the week) snapshot times.
	// +kubebuilder:validation:Required
	DayOfWeeks []DayOfWeeksParameters `json:"dayOfWeeks" tf:"day_of_weeks,omitempty"`
}

// ResourcePolicySpec defines the desired state of ResourcePolicy
type ResourcePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourcePolicyParameters `json:"forProvider"`
}

// ResourcePolicyStatus defines the observed state of ResourcePolicy.
type ResourcePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourcePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicy is the Schema for the ResourcePolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type ResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourcePolicySpec   `json:"spec"`
	Status            ResourcePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicyList contains a list of ResourcePolicys
type ResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePolicy `json:"items"`
}

// Repository type metadata.
var (
	ResourcePolicyKind             = "ResourcePolicy"
	ResourcePolicyGroupKind        = schema.GroupKind{Group: Group, Kind: ResourcePolicyKind}.String()
	ResourcePolicyKindAPIVersion   = ResourcePolicyKind + "." + GroupVersion.String()
	ResourcePolicyGroupVersionKind = GroupVersion.WithKind(ResourcePolicyKind)
)

func init() {
	SchemeBuilder.Register(&ResourcePolicy{}, &ResourcePolicyList{})
}
