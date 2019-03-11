package luminate

import (
	"context"
	"github.com/andriipetruk/go-luminate/luminate"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLuminateAppHttp() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_name": {
				Type:        schema.TypeString,
				Description: "Application name",
				Required:    true,
			},
			"internal_address": {
				Type:        schema.TypeString,
				Description: "Application dns address(svc)",
				Required:    true,
			},
			"site_id": {
				Type:        schema.TypeString,
				Description: "Site id",
				Required:    true,
			},
		},

		Create: resourceLuminateAppHttpCreate,
		Read:   resourceLuminateAppHttpRead,
		Update: resourceLuminateAppHttpUpdate,
		Delete: resourceLuminateAppHttpDelete,
	}
}

func resourceLuminateAppHttpCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)

	// -- Create http application
	newAppHttp := goluminate.AppHttpCreateRequest{Name: d.Get("app_name").(string), Type: "HTTP", IsVisible: true, IsNotificationEnabled: true}
	newAppHttp.ConnectionSettings.InternalAddress = d.Get("internal_address").(string)
	newAppHttp.ConnectionSettings.CustomRootPath = "/"
	newAppHttp.ConnectionSettings.HealthURL = "/"
	newAppHttp.ConnectionSettings.HealthMethod = "Head"
	ctx := context.Background()
	//p.log.Debug("calling resourceLuminateSiteCreate()")
	HttpApp, _, err := client.CreateApp(ctx, newAppHttp)
	if err != nil {
		return err
	}

	//  -- Bind App to Site
	client.BindAppToSite(ctx, HttpApp.ID, d.Get("site_id").(string))

	d.SetId(HttpApp.ID)

	return nil
}

func resourceLuminateAppHttpRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceLuminateAppHttpUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)

	newAppHttp := goluminate.AppHttpCreateRequest{Name: d.Get("app_name").(string), Type: "HTTP", IsVisible: true, IsNotificationEnabled: true}
	newAppHttp.ConnectionSettings.InternalAddress = d.Get("internal_address").(string)
	newAppHttp.ConnectionSettings.CustomRootPath = "/"
	newAppHttp.ConnectionSettings.HealthURL = "/"
	newAppHttp.ConnectionSettings.HealthMethod = "Head"
	ctx := context.Background()
	HttpApp, _, err := client.UpdateApp(ctx, newAppHttp, d.Id())
	if err != nil {
		return err
	}

	return nil
}

func resourceLuminateAppHttpDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	ctx := context.Background()
	client.DeleteApp(ctx, d.Id())

	return nil
}
