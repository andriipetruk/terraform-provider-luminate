package luminate

import (
	"context"
	"github.com/andriipetruk/go-luminate/luminate"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLuminateSite() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"site_name": {
				Type:        schema.TypeString,
				Description: "Site name",
				Required:    true,
				//ForceNew:    true,
			},
		},

		Create: resourceLuminateSiteCreate,
		Read:   resourceLuminateSiteRead,
		Update: resourceLuminateSiteUpdate,
		Delete: resourceLuminateSiteDelete,
	}
}

func resourceLuminateSiteCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	site := goluminate.NewSiteRequest{Name: d.Get("site_name").(string)}
	ctx := context.Background()
	//p.log.Debug("calling resourceLuminateSiteCreate()")
	newSite, _, err := client.CreateSite(ctx, site)
	if err != nil {
		return err
	}
	d.SetId(newSite.ID)

	return nil
}

func resourceLuminateSiteRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceLuminateSiteUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	site := goluminate.NewSiteRequest{Name: d.Get("site_name").(string)}
	ctx := context.Background()
	//p.log.Debug("calling resourceLuminateSiteCreate()")
	newSite, _, err := client.UpdateSite(ctx, site, d.Id())
	if err != nil {
		return err
	}

	return nil
}

func resourceLuminateSiteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*goluminate.Client)
	ctx := context.Background()
	client.DeleteSite(ctx, d.Id())
	

	return nil
}
