package storage

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_storage_bucket_object", func(r *config.Resource) {
		// Note(turkenh): We have to modify schema of
		// "customer_encryption", since it is struct marked as sensitive.
		// https://github.com/crossplane-contrib/terrajet/issues/100#issuecomment-966892273
		r.TerraformResource.Schema["customer_encryption"].Sensitive = false
	})

	p.AddResourceConfigurator("google_storage_bucket", func(r *config.Resource) {
		// Note(turkenh): Workaround for a terraform gcp provider. See the
		// following for more details: https://github.com/crossplane-contrib/provider-jet-gcp/issues/12#issuecomment-987040044
		r.ExternalName = config.IdentifierFromProvider
	})
}
