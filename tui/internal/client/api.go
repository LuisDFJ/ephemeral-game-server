package client

import (
	"bytes"
	"encoding/json"
)


func (c *Client) PostCreateVps( slug string, world_name string, world_pswd string ) error {
	_,err := c.post(
		"create-vps",
		CreateVPSRequest{
			Slug: slug,
			World: WorldParams{
				Name: world_name,
				Password: world_pswd,
			},
		},
		true)
	return err
}

func (c *Client) PostDestroyVps() error {
	_,err := c.post("destroy-vps",nil,true)
	return err
}

func (c *Client) GetStatus() (StatusResponse, error) {
	res, err := c.get("status")
	var status StatusResponse
	if err != nil { return status,err }
	err = json.NewDecoder(bytes.NewBuffer(res)).Decode(&status)
	return status,err
}

func (c *Client) GetSizes() (SizesResponse, error) {
	res, err := c.get("available-sizes")
	if err != nil { return nil,err }
	var sizes SizesResponse
	err = json.NewDecoder(bytes.NewReader(res)).Decode(&sizes)
	return sizes,err
}
