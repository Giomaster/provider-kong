package service

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("kong-gateway_service", func(r *config.Resource) {
		r.ShortGroup = "kong"
		r.ExternalName = config.IdentifierFromProvider
		r.TerraformResource.Schema["client_certificate"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		}

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"client_certificate"},
		}
	})
}
