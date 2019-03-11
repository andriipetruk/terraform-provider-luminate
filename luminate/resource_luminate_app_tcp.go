package luminate

import (
	"context"
	"github.com/andriipetruk/go-luminate/luminate"
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

func resourceLuminateAppTcp() *schema.Resource {
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
			"tcp_port": {
				Type:        schema.TypeString,
				Description: "Tcp port number",
				Required:    true,
			},
			"site_id": {
				Type:        schema.TypeString,
				Description: "Site id",
				Required:    true,
			},
		},

		Create: resourceLuminateAppTcpCreate,
		Read:   resourceLuminateAppTcpRead,
		Update: resourceLuminateAppTcpUpdate,
		Delete: resourceLuminateAppTcpDelete,
	}
}

func resourceLuminateAppTcpCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	// -- Create tcp application
	newAppTCP := goluminate.AppTcpCreateRequest{Name: d.Get("app_name").(string), Type: "TCP", IsVisible: true, IsNotificationEnabled: true}
	var TcpAppPortList []string
	var subdomain string
	subdomain = strings.Replace(d.Get("app_name").(string), " ", "", -1)
	newAppTCP.ConnectionSettings.Subdomain = strings.ToLower(subdomain)
	TcpAppPortList = append(TcpAppPortList, d.Get("tcp_port").(string))
	newAppTCP.TcpTunnelSettings = append(newAppTCP.TcpTunnelSettings, goluminate.TcpTunnelSettings{Target: d.Get("internal_address").(string), Ports: TcpAppPortList})
	ctx := context.Background()
	//p.log.Debug("calling resourceLuminateSiteCreate()")
	TCPApp, _, err := client.CreateApp(ctx, newAppTCP)
	if err != nil {
		return err
	}

	//  -- Bind App to Site
	client.BindAppToSite(ctx, TCPApp.ID, d.Get("site_id").(string))

	d.SetId(TCPApp.ID)

	return nil
}

func resourceLuminateAppTcpRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceLuminateAppTcpUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	newAppTCP := goluminate.AppTcpCreateRequest{Name: d.Get("app_name").(string), Type: "TCP", IsVisible: true, IsNotificationEnabled: true}
	var TcpAppPortList []string
	var subdomain string
	subdomain = strings.Replace(d.Get("app_name").(string), " ", "", -1)
	newAppTCP.ConnectionSettings.Subdomain = strings.ToLower(subdomain)
	TcpAppPortList = append(TcpAppPortList, d.Get("tcp_port").(string))
	newAppTCP.TcpTunnelSettings = append(newAppTCP.TcpTunnelSettings, goluminate.TcpTunnelSettings{Target: d.Get("internal_address").(string), Ports: TcpAppPortList})
	ctx := context.Background()
	_, _, err := client.UpdateApp(ctx, newAppTCP, d.Id())
	if err != nil {
		return err
	}

	return nil
}

func resourceLuminateAppTcpDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	ctx := context.Background()
	client.DeleteApp(ctx, d.Id())

	return nil
}
