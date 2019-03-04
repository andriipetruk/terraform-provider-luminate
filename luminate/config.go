package luminate

import (
	"context"
	"github.com/andriipetruk/go-luminate/luminate"
	//"github.com/hashicorp/terraform/helper/logging"
)

type Config struct {
	TenantName   string
	ClientID     string
	ClientSecret string
}

// Client() returns a new client for accessing dyn.
func (c *Config) Client() (*goluminate.Client, error) {
	ctx := context.Background()
	client := goluminate.NewClient(ctx, c.ClientID, c.ClientSecret, c.TenantName)
	/*    if logging.IsDebugOrHigher() {
	      client.Verbose(true)
	  }*/

	return client, nil
}
