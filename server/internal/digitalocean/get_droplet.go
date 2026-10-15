package digitalocean

import (
	"fmt"
	"github.com/digitalocean/godo"
)

type VPS struct {
	Alive bool `json:"alive"`
	Name string `json:"name"`
	Ip string `json:"ip"`
	Specs string `json:"specs"`
}

func NewVPS( droplet godo.Droplet ) VPS {
	ip,_ := droplet.PublicIPv4()
	return VPS {
		Alive: true,
		Name: droplet.Name,
		Ip: ip,
		Specs: droplet.SizeSlug,
	}
}

func (c *Client) GetDroplet() (VPS, error) {
	var vps VPS
	c.mu.Lock()
	defer c.mu.Unlock()
	droplets, _, err := c.client.Droplets.List(c.ctx, &godo.ListOptions{})
	if err != nil {
		return vps, fmt.Errorf("failed to list droplets: %v\n", err)
	}
	for i := range droplets {
		if droplets[i].Name == VALHEIM_DROPLET {
			vps = NewVPS(droplets[i])
			break
		}
	}
	return vps,nil
}
