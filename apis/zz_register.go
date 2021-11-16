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

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1accessapproval "github.com/crossplane-contrib/provider-tf-gcp/apis/accessapproval/v1alpha1"
	v1alpha1accesscontextmanager "github.com/crossplane-contrib/provider-tf-gcp/apis/accesscontextmanager/v1alpha1"
	v1alpha1activedirectory "github.com/crossplane-contrib/provider-tf-gcp/apis/activedirectory/v1alpha1"
	v1alpha1apigee "github.com/crossplane-contrib/provider-tf-gcp/apis/apigee/v1alpha1"
	v1alpha1appengine "github.com/crossplane-contrib/provider-tf-gcp/apis/appengine/v1alpha1"
	v1alpha1assuredworkloads "github.com/crossplane-contrib/provider-tf-gcp/apis/assuredworkloads/v1alpha1"
	v1alpha1bigquery "github.com/crossplane-contrib/provider-tf-gcp/apis/bigquery/v1alpha1"
	v1alpha1bigtable "github.com/crossplane-contrib/provider-tf-gcp/apis/bigtable/v1alpha1"
	v1alpha1billing "github.com/crossplane-contrib/provider-tf-gcp/apis/billing/v1alpha1"
	v1alpha1binaryauthorization "github.com/crossplane-contrib/provider-tf-gcp/apis/binaryauthorization/v1alpha1"
	v1alpha1cloudasset "github.com/crossplane-contrib/provider-tf-gcp/apis/cloudasset/v1alpha1"
	v1alpha1cloudbuild "github.com/crossplane-contrib/provider-tf-gcp/apis/cloudbuild/v1alpha1"
	v1alpha1cloudfunctions "github.com/crossplane-contrib/provider-tf-gcp/apis/cloudfunctions/v1alpha1"
	v1alpha1cloudidentity "github.com/crossplane-contrib/provider-tf-gcp/apis/cloudidentity/v1alpha1"
	v1alpha1cloudiot "github.com/crossplane-contrib/provider-tf-gcp/apis/cloudiot/v1alpha1"
	v1alpha1cloudplatform "github.com/crossplane-contrib/provider-tf-gcp/apis/cloudplatform/v1alpha1"
	v1alpha1cloudrun "github.com/crossplane-contrib/provider-tf-gcp/apis/cloudrun/v1alpha1"
	v1alpha1cloudscheduler "github.com/crossplane-contrib/provider-tf-gcp/apis/cloudscheduler/v1alpha1"
	v1alpha1cloudtasks "github.com/crossplane-contrib/provider-tf-gcp/apis/cloudtasks/v1alpha1"
	v1alpha1composer "github.com/crossplane-contrib/provider-tf-gcp/apis/composer/v1alpha1"
	v1alpha1compute "github.com/crossplane-contrib/provider-tf-gcp/apis/compute/v1alpha1"
	v1alpha1container "github.com/crossplane-contrib/provider-tf-gcp/apis/container/v1alpha1"
	v1alpha1datacatalog "github.com/crossplane-contrib/provider-tf-gcp/apis/datacatalog/v1alpha1"
	v1alpha1dataflow "github.com/crossplane-contrib/provider-tf-gcp/apis/dataflow/v1alpha1"
	v1alpha1datalossprevention "github.com/crossplane-contrib/provider-tf-gcp/apis/datalossprevention/v1alpha1"
	v1alpha1dataproc "github.com/crossplane-contrib/provider-tf-gcp/apis/dataproc/v1alpha1"
	v1alpha1datastore "github.com/crossplane-contrib/provider-tf-gcp/apis/datastore/v1alpha1"
	v1alpha1deploymentmanager "github.com/crossplane-contrib/provider-tf-gcp/apis/deploymentmanager/v1alpha1"
	v1alpha1dialogflow "github.com/crossplane-contrib/provider-tf-gcp/apis/dialogflow/v1alpha1"
	v1alpha1dns "github.com/crossplane-contrib/provider-tf-gcp/apis/dns/v1alpha1"
	v1alpha1endpoints "github.com/crossplane-contrib/provider-tf-gcp/apis/endpoints/v1alpha1"
	v1alpha1essentialcontacts "github.com/crossplane-contrib/provider-tf-gcp/apis/essentialcontacts/v1alpha1"
	v1alpha1eventarc "github.com/crossplane-contrib/provider-tf-gcp/apis/eventarc/v1alpha1"
	v1alpha1filestore "github.com/crossplane-contrib/provider-tf-gcp/apis/filestore/v1alpha1"
	v1alpha1firestore "github.com/crossplane-contrib/provider-tf-gcp/apis/firestore/v1alpha1"
	v1alpha1gameservices "github.com/crossplane-contrib/provider-tf-gcp/apis/gameservices/v1alpha1"
	v1alpha1gkehub "github.com/crossplane-contrib/provider-tf-gcp/apis/gkehub/v1alpha1"
	v1alpha1healthcare "github.com/crossplane-contrib/provider-tf-gcp/apis/healthcare/v1alpha1"
	v1alpha1iap "github.com/crossplane-contrib/provider-tf-gcp/apis/iap/v1alpha1"
	v1alpha1identityplatform "github.com/crossplane-contrib/provider-tf-gcp/apis/identityplatform/v1alpha1"
	v1alpha1kms "github.com/crossplane-contrib/provider-tf-gcp/apis/kms/v1alpha1"
	v1alpha1logging "github.com/crossplane-contrib/provider-tf-gcp/apis/logging/v1alpha1"
	v1alpha1memcache "github.com/crossplane-contrib/provider-tf-gcp/apis/memcache/v1alpha1"
	v1alpha1mlengine "github.com/crossplane-contrib/provider-tf-gcp/apis/mlengine/v1alpha1"
	v1alpha1monitoring "github.com/crossplane-contrib/provider-tf-gcp/apis/monitoring/v1alpha1"
	v1alpha1networkmanagement "github.com/crossplane-contrib/provider-tf-gcp/apis/networkmanagement/v1alpha1"
	v1alpha1networkservices "github.com/crossplane-contrib/provider-tf-gcp/apis/networkservices/v1alpha1"
	v1alpha1notebooks "github.com/crossplane-contrib/provider-tf-gcp/apis/notebooks/v1alpha1"
	v1alpha1orgpolicy "github.com/crossplane-contrib/provider-tf-gcp/apis/orgpolicy/v1alpha1"
	v1alpha1os "github.com/crossplane-contrib/provider-tf-gcp/apis/os/v1alpha1"
	v1alpha1privateca "github.com/crossplane-contrib/provider-tf-gcp/apis/privateca/v1alpha1"
	v1alpha1pubsub "github.com/crossplane-contrib/provider-tf-gcp/apis/pubsub/v1alpha1"
	v1alpha1redis "github.com/crossplane-contrib/provider-tf-gcp/apis/redis/v1alpha1"
	v1alpha1resourcemanager "github.com/crossplane-contrib/provider-tf-gcp/apis/resourcemanager/v1alpha1"
	v1alpha1scc "github.com/crossplane-contrib/provider-tf-gcp/apis/scc/v1alpha1"
	v1alpha1secretmanager "github.com/crossplane-contrib/provider-tf-gcp/apis/secretmanager/v1alpha1"
	v1alpha1servicenetworking "github.com/crossplane-contrib/provider-tf-gcp/apis/servicenetworking/v1alpha1"
	v1alpha1sourcerepo "github.com/crossplane-contrib/provider-tf-gcp/apis/sourcerepo/v1alpha1"
	v1alpha1spanner "github.com/crossplane-contrib/provider-tf-gcp/apis/spanner/v1alpha1"
	v1alpha1sql "github.com/crossplane-contrib/provider-tf-gcp/apis/sql/v1alpha1"
	v1alpha1storage "github.com/crossplane-contrib/provider-tf-gcp/apis/storage/v1alpha1"
	v1alpha1tags "github.com/crossplane-contrib/provider-tf-gcp/apis/tags/v1alpha1"
	v1alpha1tpu "github.com/crossplane-contrib/provider-tf-gcp/apis/tpu/v1alpha1"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-gcp/apis/v1alpha1"
	v1alpha1vertex "github.com/crossplane-contrib/provider-tf-gcp/apis/vertex/v1alpha1"
	v1alpha1vpcaccess "github.com/crossplane-contrib/provider-tf-gcp/apis/vpcaccess/v1alpha1"
	v1alpha1workflows "github.com/crossplane-contrib/provider-tf-gcp/apis/workflows/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1accessapproval.SchemeBuilder.AddToScheme,
		v1alpha1accesscontextmanager.SchemeBuilder.AddToScheme,
		v1alpha1activedirectory.SchemeBuilder.AddToScheme,
		v1alpha1apigee.SchemeBuilder.AddToScheme,
		v1alpha1appengine.SchemeBuilder.AddToScheme,
		v1alpha1assuredworkloads.SchemeBuilder.AddToScheme,
		v1alpha1bigquery.SchemeBuilder.AddToScheme,
		v1alpha1bigtable.SchemeBuilder.AddToScheme,
		v1alpha1billing.SchemeBuilder.AddToScheme,
		v1alpha1binaryauthorization.SchemeBuilder.AddToScheme,
		v1alpha1cloudasset.SchemeBuilder.AddToScheme,
		v1alpha1cloudbuild.SchemeBuilder.AddToScheme,
		v1alpha1cloudfunctions.SchemeBuilder.AddToScheme,
		v1alpha1cloudidentity.SchemeBuilder.AddToScheme,
		v1alpha1cloudiot.SchemeBuilder.AddToScheme,
		v1alpha1cloudplatform.SchemeBuilder.AddToScheme,
		v1alpha1cloudrun.SchemeBuilder.AddToScheme,
		v1alpha1cloudscheduler.SchemeBuilder.AddToScheme,
		v1alpha1cloudtasks.SchemeBuilder.AddToScheme,
		v1alpha1composer.SchemeBuilder.AddToScheme,
		v1alpha1compute.SchemeBuilder.AddToScheme,
		v1alpha1container.SchemeBuilder.AddToScheme,
		v1alpha1datacatalog.SchemeBuilder.AddToScheme,
		v1alpha1dataflow.SchemeBuilder.AddToScheme,
		v1alpha1datalossprevention.SchemeBuilder.AddToScheme,
		v1alpha1dataproc.SchemeBuilder.AddToScheme,
		v1alpha1datastore.SchemeBuilder.AddToScheme,
		v1alpha1deploymentmanager.SchemeBuilder.AddToScheme,
		v1alpha1dialogflow.SchemeBuilder.AddToScheme,
		v1alpha1dns.SchemeBuilder.AddToScheme,
		v1alpha1endpoints.SchemeBuilder.AddToScheme,
		v1alpha1essentialcontacts.SchemeBuilder.AddToScheme,
		v1alpha1eventarc.SchemeBuilder.AddToScheme,
		v1alpha1filestore.SchemeBuilder.AddToScheme,
		v1alpha1firestore.SchemeBuilder.AddToScheme,
		v1alpha1gameservices.SchemeBuilder.AddToScheme,
		v1alpha1gkehub.SchemeBuilder.AddToScheme,
		v1alpha1healthcare.SchemeBuilder.AddToScheme,
		v1alpha1iap.SchemeBuilder.AddToScheme,
		v1alpha1identityplatform.SchemeBuilder.AddToScheme,
		v1alpha1kms.SchemeBuilder.AddToScheme,
		v1alpha1logging.SchemeBuilder.AddToScheme,
		v1alpha1memcache.SchemeBuilder.AddToScheme,
		v1alpha1mlengine.SchemeBuilder.AddToScheme,
		v1alpha1monitoring.SchemeBuilder.AddToScheme,
		v1alpha1networkmanagement.SchemeBuilder.AddToScheme,
		v1alpha1networkservices.SchemeBuilder.AddToScheme,
		v1alpha1notebooks.SchemeBuilder.AddToScheme,
		v1alpha1orgpolicy.SchemeBuilder.AddToScheme,
		v1alpha1os.SchemeBuilder.AddToScheme,
		v1alpha1privateca.SchemeBuilder.AddToScheme,
		v1alpha1pubsub.SchemeBuilder.AddToScheme,
		v1alpha1redis.SchemeBuilder.AddToScheme,
		v1alpha1resourcemanager.SchemeBuilder.AddToScheme,
		v1alpha1scc.SchemeBuilder.AddToScheme,
		v1alpha1secretmanager.SchemeBuilder.AddToScheme,
		v1alpha1servicenetworking.SchemeBuilder.AddToScheme,
		v1alpha1sourcerepo.SchemeBuilder.AddToScheme,
		v1alpha1spanner.SchemeBuilder.AddToScheme,
		v1alpha1sql.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha1tags.SchemeBuilder.AddToScheme,
		v1alpha1tpu.SchemeBuilder.AddToScheme,
		v1alpha1vertex.SchemeBuilder.AddToScheme,
		v1alpha1vpcaccess.SchemeBuilder.AddToScheme,
		v1alpha1workflows.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
