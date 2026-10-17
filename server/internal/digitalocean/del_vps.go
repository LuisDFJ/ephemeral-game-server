package digitalocean

import (
	"fmt"
)

func (c *Client) DelVps() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_,err := c.client.Droplets.DeleteByTag(c.ctx, VALHEIM_DROPLET)
	if err != nil {
		return fmt.Errorf("Could not delete droplet: %v", err)
	}
	return nil
}
