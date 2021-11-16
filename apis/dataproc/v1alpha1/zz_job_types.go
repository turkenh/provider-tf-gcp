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

type HadoopConfigObservation struct {
}

type HadoopConfigParameters struct {

	// HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
	// +kubebuilder:validation:Optional
	ArchiveUris []*string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`

	// The arguments to pass to the driver.
	// +kubebuilder:validation:Optional
	Args []*string `json:"args,omitempty" tf:"args,omitempty"`

	// HCFS URIs of files to be copied to the working directory of Spark drivers and distributed tasks. Useful for naively parallel tasks.
	// +kubebuilder:validation:Optional
	FileUris []*string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`

	// HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.
	// +kubebuilder:validation:Optional
	JarFileUris []*string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`

	// The runtime logging config of the job
	// +kubebuilder:validation:Optional
	LoggingConfig []LoggingConfigParameters `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`

	// The class containing the main method of the driver. Must be in a provided jar or jar that is already on the classpath. Conflicts with main_jar_file_uri
	// +kubebuilder:validation:Optional
	MainClass *string `json:"mainClass,omitempty" tf:"main_class,omitempty"`

	// The HCFS URI of jar file containing the driver jar. Conflicts with main_class
	// +kubebuilder:validation:Optional
	MainJarFileURI *string `json:"mainJarFileUri,omitempty" tf:"main_jar_file_uri,omitempty"`

	// A mapping of property names to values, used to configure Spark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type HiveConfigObservation struct {
}

type HiveConfigParameters struct {

	// Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.
	// +kubebuilder:validation:Optional
	ContinueOnFailure *bool `json:"continueOnFailure,omitempty" tf:"continue_on_failure,omitempty"`

	// HCFS URIs of jar files to add to the CLASSPATH of the Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive SerDes and UDFs.
	// +kubebuilder:validation:Optional
	JarFileUris []*string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`

	// A mapping of property names and values, used to configure Hive. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/hive/conf/hive-site.xml, and classes in user code.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// HCFS URI of file containing Hive script to execute as the job. Conflicts with query_list
	// +kubebuilder:validation:Optional
	QueryFileURI *string `json:"queryFileUri,omitempty" tf:"query_file_uri,omitempty"`

	// The list of Hive queries or statements to execute as part of the job. Conflicts with query_file_uri
	// +kubebuilder:validation:Optional
	QueryList []*string `json:"queryList,omitempty" tf:"query_list,omitempty"`

	// Mapping of query variable names to values (equivalent to the Hive command: SET name="value";).
	// +kubebuilder:validation:Optional
	ScriptVariables map[string]*string `json:"scriptVariables,omitempty" tf:"script_variables,omitempty"`
}

type JobObservation struct {
	DriverControlsFilesURI *string `json:"driverControlsFilesUri,omitempty" tf:"driver_controls_files_uri,omitempty"`

	DriverOutputResourceURI *string `json:"driverOutputResourceUri,omitempty" tf:"driver_output_resource_uri,omitempty"`

	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`
}

type JobParameters struct {

	// By default, you can only delete inactive jobs within Dataproc. Setting this to true, and calling destroy, will ensure that the job is first cancelled before issuing the delete.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// The config of Hadoop job
	// +kubebuilder:validation:Optional
	HadoopConfig []HadoopConfigParameters `json:"hadoopConfig,omitempty" tf:"hadoop_config,omitempty"`

	// The config of hive job
	// +kubebuilder:validation:Optional
	HiveConfig []HiveConfigParameters `json:"hiveConfig,omitempty" tf:"hive_config,omitempty"`

	// Optional. The labels to associate with this job.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The config of pag job.
	// +kubebuilder:validation:Optional
	PigConfig []PigConfigParameters `json:"pigConfig,omitempty" tf:"pig_config,omitempty"`

	// The config of job placement.
	// +kubebuilder:validation:Required
	Placement []PlacementParameters `json:"placement" tf:"placement,omitempty"`

	// The project in which the cluster can be found and jobs subsequently run against. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The config of pySpark job.
	// +kubebuilder:validation:Optional
	PysparkConfig []PysparkConfigParameters `json:"pysparkConfig,omitempty" tf:"pyspark_config,omitempty"`

	// The reference of the job
	// +kubebuilder:validation:Optional
	Reference []ReferenceParameters `json:"reference,omitempty" tf:"reference,omitempty"`

	// The Cloud Dataproc region. This essentially determines which clusters are available for this job to be submitted to. If not specified, defaults to global.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Optional. Job scheduling configuration.
	// +kubebuilder:validation:Optional
	Scheduling []SchedulingParameters `json:"scheduling,omitempty" tf:"scheduling,omitempty"`

	// The config of the Spark job.
	// +kubebuilder:validation:Optional
	SparkConfig []SparkConfigParameters `json:"sparkConfig,omitempty" tf:"spark_config,omitempty"`

	// The config of SparkSql job
	// +kubebuilder:validation:Optional
	SparksqlConfig []SparksqlConfigParameters `json:"sparksqlConfig,omitempty" tf:"sparksql_config,omitempty"`
}

type LoggingConfigObservation struct {
}

type LoggingConfigParameters struct {

	// Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	// +kubebuilder:validation:Required
	DriverLogLevels map[string]*string `json:"driverLogLevels" tf:"driver_log_levels,omitempty"`
}

type PigConfigLoggingConfigObservation struct {
}

type PigConfigLoggingConfigParameters struct {

	// Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	// +kubebuilder:validation:Required
	DriverLogLevels map[string]*string `json:"driverLogLevels" tf:"driver_log_levels,omitempty"`
}

type PigConfigObservation struct {
}

type PigConfigParameters struct {

	// Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.
	// +kubebuilder:validation:Optional
	ContinueOnFailure *bool `json:"continueOnFailure,omitempty" tf:"continue_on_failure,omitempty"`

	// HCFS URIs of jar files to add to the CLASSPATH of the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig UDFs.
	// +kubebuilder:validation:Optional
	JarFileUris []*string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`

	// The runtime logging config of the job
	// +kubebuilder:validation:Optional
	LoggingConfig []PigConfigLoggingConfigParameters `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`

	// A mapping of property names to values, used to configure Pig. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/pig/conf/pig.properties, and classes in user code.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// HCFS URI of file containing Hive script to execute as the job. Conflicts with query_list
	// +kubebuilder:validation:Optional
	QueryFileURI *string `json:"queryFileUri,omitempty" tf:"query_file_uri,omitempty"`

	// The list of Hive queries or statements to execute as part of the job. Conflicts with query_file_uri
	// +kubebuilder:validation:Optional
	QueryList []*string `json:"queryList,omitempty" tf:"query_list,omitempty"`

	// Mapping of query variable names to values (equivalent to the Pig command: name=[value]).
	// +kubebuilder:validation:Optional
	ScriptVariables map[string]*string `json:"scriptVariables,omitempty" tf:"script_variables,omitempty"`
}

type PlacementObservation struct {
	ClusterUUID *string `json:"clusterUuid,omitempty" tf:"cluster_uuid,omitempty"`
}

type PlacementParameters struct {

	// The name of the cluster where the job will be submitted
	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`
}

type PysparkConfigLoggingConfigObservation struct {
}

type PysparkConfigLoggingConfigParameters struct {

	// Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	// +kubebuilder:validation:Required
	DriverLogLevels map[string]*string `json:"driverLogLevels" tf:"driver_log_levels,omitempty"`
}

type PysparkConfigObservation struct {
}

type PysparkConfigParameters struct {

	// Optional. HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip
	// +kubebuilder:validation:Optional
	ArchiveUris []*string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`

	// Optional. The arguments to pass to the driver. Do not include arguments, such as --conf, that can be set as job properties, since a collision may occur that causes an incorrect job submission
	// +kubebuilder:validation:Optional
	Args []*string `json:"args,omitempty" tf:"args,omitempty"`

	// Optional. HCFS URIs of files to be copied to the working directory of Python drivers and distributed tasks. Useful for naively parallel tasks
	// +kubebuilder:validation:Optional
	FileUris []*string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`

	// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks
	// +kubebuilder:validation:Optional
	JarFileUris []*string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`

	// The runtime logging config of the job
	// +kubebuilder:validation:Optional
	LoggingConfig []PysparkConfigLoggingConfigParameters `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`

	// Required. The HCFS URI of the main Python file to use as the driver. Must be a .py file
	// +kubebuilder:validation:Required
	MainPythonFileURI *string `json:"mainPythonFileUri" tf:"main_python_file_uri,omitempty"`

	// Optional. A mapping of property names to values, used to configure PySpark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// Optional. HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip
	// +kubebuilder:validation:Optional
	PythonFileUris []*string `json:"pythonFileUris,omitempty" tf:"python_file_uris,omitempty"`
}

type ReferenceObservation struct {
}

type ReferenceParameters struct {

	// The job ID, which must be unique within the project. The job ID is generated by the server upon job submission or provided by the user as a means to perform retries without creating duplicate jobs
	// +kubebuilder:validation:Optional
	JobID *string `json:"jobId,omitempty" tf:"job_id,omitempty"`
}

type SchedulingObservation struct {
}

type SchedulingParameters struct {

	// Maximum number of times per hour a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.
	// +kubebuilder:validation:Required
	MaxFailuresPerHour *int64 `json:"maxFailuresPerHour" tf:"max_failures_per_hour,omitempty"`

	// Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.
	// +kubebuilder:validation:Required
	MaxFailuresTotal *int64 `json:"maxFailuresTotal" tf:"max_failures_total,omitempty"`
}

type SparkConfigLoggingConfigObservation struct {
}

type SparkConfigLoggingConfigParameters struct {

	// Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	// +kubebuilder:validation:Required
	DriverLogLevels map[string]*string `json:"driverLogLevels" tf:"driver_log_levels,omitempty"`
}

type SparkConfigObservation struct {
}

type SparkConfigParameters struct {

	// HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.
	// +kubebuilder:validation:Optional
	ArchiveUris []*string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`

	// The arguments to pass to the driver.
	// +kubebuilder:validation:Optional
	Args []*string `json:"args,omitempty" tf:"args,omitempty"`

	// HCFS URIs of files to be copied to the working directory of Spark drivers and distributed tasks. Useful for naively parallel tasks.
	// +kubebuilder:validation:Optional
	FileUris []*string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`

	// HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.
	// +kubebuilder:validation:Optional
	JarFileUris []*string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`

	// The runtime logging config of the job
	// +kubebuilder:validation:Optional
	LoggingConfig []SparkConfigLoggingConfigParameters `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`

	// The class containing the main method of the driver. Must be in a provided jar or jar that is already on the classpath. Conflicts with main_jar_file_uri
	// +kubebuilder:validation:Optional
	MainClass *string `json:"mainClass,omitempty" tf:"main_class,omitempty"`

	// The HCFS URI of jar file containing the driver jar. Conflicts with main_class
	// +kubebuilder:validation:Optional
	MainJarFileURI *string `json:"mainJarFileUri,omitempty" tf:"main_jar_file_uri,omitempty"`

	// A mapping of property names to values, used to configure Spark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type SparksqlConfigLoggingConfigObservation struct {
}

type SparksqlConfigLoggingConfigParameters struct {

	// Optional. The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	// +kubebuilder:validation:Required
	DriverLogLevels map[string]*string `json:"driverLogLevels" tf:"driver_log_levels,omitempty"`
}

type SparksqlConfigObservation struct {
}

type SparksqlConfigParameters struct {

	// HCFS URIs of jar files to be added to the Spark CLASSPATH.
	// +kubebuilder:validation:Optional
	JarFileUris []*string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`

	// The runtime logging config of the job
	// +kubebuilder:validation:Optional
	LoggingConfig []SparksqlConfigLoggingConfigParameters `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`

	// A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// The HCFS URI of the script that contains SQL queries. Conflicts with query_list
	// +kubebuilder:validation:Optional
	QueryFileURI *string `json:"queryFileUri,omitempty" tf:"query_file_uri,omitempty"`

	// The list of SQL queries or statements to execute as part of the job. Conflicts with query_file_uri
	// +kubebuilder:validation:Optional
	QueryList []*string `json:"queryList,omitempty" tf:"query_list,omitempty"`

	// Mapping of query variable names to values (equivalent to the Spark SQL command: SET name="value";).
	// +kubebuilder:validation:Optional
	ScriptVariables map[string]*string `json:"scriptVariables,omitempty" tf:"script_variables,omitempty"`
}

type StatusObservation struct {
	Details *string `json:"details,omitempty" tf:"details,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	StateStartTime *string `json:"stateStartTime,omitempty" tf:"state_start_time,omitempty"`

	Substate *string `json:"substate,omitempty" tf:"substate,omitempty"`
}

type StatusParameters struct {
}

// JobSpec defines the desired state of Job
type JobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobParameters `json:"forProvider"`
}

// JobStatus defines the observed state of Job.
type JobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Job is the Schema for the Jobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobSpec   `json:"spec"`
	Status            JobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobList contains a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Job `json:"items"`
}

// Repository type metadata.
var (
	JobKind             = "Job"
	JobGroupKind        = schema.GroupKind{Group: Group, Kind: JobKind}.String()
	JobKindAPIVersion   = JobKind + "." + GroupVersion.String()
	JobGroupVersionKind = GroupVersion.WithKind(JobKind)
)

func init() {
	SchemeBuilder.Register(&Job{}, &JobList{})
}
