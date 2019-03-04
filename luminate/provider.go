package luminate

import (
	"fmt"

	"github.com/facette/logger"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

type luminateProvider struct {
	log *logger.Logger
}

var providerLogFile = "terraform-provider-luminate.log"

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"debug": {
				Type:        schema.TypeBool,
				Description: fmt.Sprintf("Enable provider debug logging (logs to file %s)", providerLogFile),
				Optional:    true,
				Default:     false,
			},
			"tenant_name": {
				Type:        schema.TypeString,
				Description: "tenant name",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TENANT_NAME", nil),
			},
			"client_id": {
				Type:        schema.TypeString,
				Description: "client ID for API access",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", nil),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Description: "client secret for API access",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_SECRET", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"luminate_site":      resourceLuminateSite(),
			"luminate_connector": resourceLuminateConnector(),
			"luminate_app_http":  resourceLuminateAppHttp(),
			"luminate_app_ssh":   resourceLuminateAppSsh(),
			"luminate_app_tcp":   resourceLuminateAppTcp(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		TenantName:   d.Get("tenant_name").(string),
		ClientID:     d.Get("client_id").(string),
		ClientSecret: d.Get("client_secret").(string),
	}

	return config.Client()
}
