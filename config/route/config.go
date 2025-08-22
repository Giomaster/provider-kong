package route

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("kong-gateway_route", func(r *config.Resource) {
		r.ShortGroup = "kong"
		r.ExternalName = config.IdentifierFromProvider
		r.TerraformResource.Schema["destinations"] = &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ip": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"port": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		}

		r.TerraformResource.Schema["service"] = &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		}

		r.TerraformResource.Schema["sources"] = &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ip": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"port": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		}

		r.AddSingletonListConversion("service[0]", "spec.forProvider.service")
		r.TerraformConversions = append(r.TerraformConversions, config.NewTFSingletonConversion())
		r.References["service.id"] = config.Reference{
			TerraformName: "kong-gateway_service",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"destinations", "service", "sources"},
		}
	})
}
