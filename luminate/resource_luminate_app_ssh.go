package luminate

import (
	"context"
	"github.com/andriipetruk/go-luminate/luminate"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLuminateAppSsh() *schema.Resource {
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
			"ssh_login": {
				Type:        schema.TypeString,
				Description: "ssh login",
				Required:    true,
			},
			"site_id": {
				Type:        schema.TypeString,
				Description: "Site id",
				Required:    true,
			},
		},

		Create: resourceLuminateAppSshCreate,
		Read:   resourceLuminateAppSshRead,
		Update: resourceLuminateAppSshUpdate,
		Delete: resourceLuminateAppSshDelete,
	}
}

func resourceLuminateAppSshCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	// -- Create ssh application
	newAppSSH := goluminate.AppSshCreateRequest{Name: d.Get("app_name").(string), Type: "SSH", IsVisible: true, IsNotificationEnabled: true}
	newAppSSH.ConnectionSettings.InternalAddress = d.Get("internal_address").(string)
	newAppSSH.SSHSettings.UserAccounts = append(newAppSSH.SSHSettings.UserAccounts, goluminate.SshUserAccounts{Name: d.Get("ssh_login").(string)})
	ctx := context.Background()
	//p.log.Debug("calling resourceLuminateSiteCreate()")
	SSHApp, _, err := client.CreateApp(ctx, newAppSSH)
	if err != nil {
		return err
	}

	//  -- Bind App to Site
	client.BindAppToSite(ctx, SSHApp.ID, d.Get("site_id").(string))

	d.SetId(SSHApp.ID)

	return nil
}

func resourceLuminateAppSshRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceLuminateAppSshUpdate(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceLuminateAppSshDelete(d *schema.ResourceData, meta interface{}) error {

	return nil
}
