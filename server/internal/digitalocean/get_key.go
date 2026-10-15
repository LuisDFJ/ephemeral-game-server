package digitalocean

import (
	"fmt"
	"github.com/digitalocean/godo"
)

type Key struct {
	Available bool
	Name string
	ID int
	Fingerprint string
	PublicKey string
}

func NewKey( key godo.Key ) Key {
	return Key {
		Available: true,
		Name: key.Name,
		ID: key.ID,
		Fingerprint: key.Fingerprint,
		PublicKey: key.PublicKey,
	}
}

func (c *Client) GetKey() (Key, error) {
	var key Key
	c.mu.Lock()
	defer c.mu.Unlock()
	keys, _, err := c.client.Keys.List(c.ctx, &godo.ListOptions{})
	if err != nil {
		return key, fmt.Errorf("failed to list keys: %v\n", err)
	}
	for i := range keys {
		key = NewKey(keys[i])
		break
	}
	return key,nil
}
