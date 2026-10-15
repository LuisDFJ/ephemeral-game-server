package digitalocean

import (
	"fmt"
	userdata "server/internal/user_data"
	"github.com/digitalocean/godo"
)


func (c *Client) PostVps( ud userdata.Params, size Size, image Image, volume Volume, key Key ) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_,_,err := c.client.Droplets.Create(c.ctx, &godo.DropletCreateRequest{
		Name: VALHEIM_DROPLET,
		Region: REGION,
		Size: size.Slug,
		Image: godo.DropletCreateImage{ID: image.ID},
		Volumes: []godo.DropletCreateVolume{
			godo.DropletCreateVolume{ID: volume.ID},
		},
		SSHKeys: []godo.DropletCreateSSHKey{
			godo.DropletCreateSSHKey{ID: key.ID},
		},
		Backups: false,
		IPv6: false,
		Monitoring: false,
		UserData: ud.Gen(),
	})
	if err != nil {
		return fmt.Errorf("failed to create droplet:  %v\n", err)
	}
	return nil
}
