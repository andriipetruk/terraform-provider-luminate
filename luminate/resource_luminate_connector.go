package luminate

import (
	"context"
	"fmt"
	"github.com/andriipetruk/go-luminate/luminate"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLuminateConnector() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connector_name": {
				Type:        schema.TypeString,
				Description: "Connector name",
				Required:    true,
			},
			"site_id": {
				Type:        schema.TypeString,
				Description: "Site id",
				Required:    true,
			},
			"install": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},

		Create: resourceLuminateConnectorCreate,
		Read:   resourceLuminateConnectorRead,
		Update: resourceLuminateConnectorUpdate,
		Delete: resourceLuminateConnectorDelete,
	}
}

func resourceLuminateConnectorCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	connector := goluminate.NewConnectorRequest{Name: d.Get("connector_name").(string), Version: "1.0"}
	ctx := context.Background()
	//p.log.Debug("calling resourceLuminateSiteCreate()")
	newConnector, _, err := client.CreateConnector(ctx, connector, d.Get("site_id").(string))
	if err != nil {
		return err
	}

	d.SetId(newConnector.ID)

	return resourceLuminateConnectorRead(d, meta)
}

func resourceLuminateConnectorRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	ctx := context.Background()
	// Get K8S install command
	connectorGetCommand := goluminate.ConnectorCommandRequest{ConnectorName: d.Get("connector_name").(string)}
	ConnectorInstall, _, err := client.GetConnectorCommand(ctx, connectorGetCommand, d.Id())
	if err != nil {
		return fmt.Errorf("Connector not exist: %s", err)
	}
	d.Set("install", ConnectorInstall.K8S)

	return nil
}

func resourceLuminateConnectorUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	connector := goluminate.NewConnectorRequest{Name: d.Get("connector_name").(string), Version: "1.0"}
	ctx := context.Background()
	newConnector, _, err := client.UpdateConnector(ctx, connector,  d.Id())
	if err != nil {
		return err
	}

	return nil
}

func resourceLuminateConnectorDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*goluminate.Client)
	ctx := context.Background()
	client.DeleteConnector(ctx, d.Id())

	return nil
}
