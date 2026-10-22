package digitalocean

import (
	"fmt"
	"slices"

	"github.com/digitalocean/godo"
)

type Image struct {
	Available bool
	Name string
	ID int
}

func NewImage(image godo.Image) Image {
	return Image{
		Available: true,
		Name: image.Name,
		ID: image.ID,
	}
}

func (c *Client) GetImage() (Image, error) {
	var image Image
	c.mu.Lock()
	defer c.mu.Unlock()
	images, _, err := c.client.Images.ListUser(c.ctx, &godo.ListOptions{})
	if err != nil {
		return image, fmt.Errorf("failed to list images: %v\n", err)
	}
	for i := range images {
		region := slices.Contains(images[i].Regions, c.region)
		if images[i].Name == VALHEIM_IMAGE && region {
			image = NewImage(images[i])
			break
		}
	}
	return image,nil
}
