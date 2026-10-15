package digitalocean

import (
	"fmt"
	"github.com/digitalocean/godo"
)

type Volume struct {
	Available bool
	Name string
	ID string
}

func NewVolume( volume godo.Volume ) Volume {
	return Volume {
		Available: true,
		Name: volume.Name,
		ID: volume.ID,
	}
}

func (c *Client) GetVolume() (Volume, error) {
	var volume Volume
	c.mu.Lock()
	defer c.mu.Unlock()
	volumes, _, err := c.client.Storage.ListVolumes(c.ctx, &godo.ListVolumeParams{Region: REGION})
	if err != nil {
		return volume, fmt.Errorf("failed to list volumes: %v\n", err)
	}
	for i := range volumes {
		if volumes[i].Name == VALHEIM_VOLUME {
			volume = NewVolume(volumes[i])
			break
		}
	}
	return volume,nil
}
