package digitalocean

import (
	"fmt"
)

func (c *Client) DelVps(id int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_,err := c.client.Droplets.Delete(c.ctx, id)
	if err != nil {
		return fmt.Errorf("Could not delete droplet: %v", err)
	}
	return nil
}
