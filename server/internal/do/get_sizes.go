package do

import (
	"fmt"
	"slices"
	"github.com/digitalocean/godo"
)

type Size struct {
	Slug string `json:"slug"`
	Disk int `json:"disk"`
	Mem int `json:"mem"`
	VCpu int `json:"vcpu"`
	PriceMonth float64 `json:"pricemonth"`
	PriceHour float64 `json:"pricehour"`
}

func NewSize( size godo.Size ) Size {
	return Size {
		Slug: size.Slug,
		Disk: size.Disk,
		Mem: size.Memory,
		VCpu: size.Vcpus,
		PriceMonth: size.PriceMonthly,
		PriceHour: size.PriceHourly,
	}
}

func (c *Client) GetSizes() ([]Size, error) {
	var res = []Size{}
	c.mu.Lock()
	defer c.mu.Unlock()
	sizes, _, err := c.client.Sizes.List(c.ctx, &godo.ListOptions{PerPage: 100, Page: 1})
	if err != nil {
		return res, fmt.Errorf("failed to list droplets: %v\n", err)
	}
	for i := range sizes {
		if isValidSize(&sizes[i]) {
			res = append(res, NewSize(sizes[i]))
		}
	}
	return res,nil
}

func isValidSize(size *godo.Size) bool {
	region := slices.Contains(size.Regions, REGION)
	return size.Available && size.Memory >= 4000 && size.Vcpus >= 2 && region && size.PriceMonthly <= MAX_PRICE
}
