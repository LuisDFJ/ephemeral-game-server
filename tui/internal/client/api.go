package client

import (
	"bytes"
	"encoding/json"
	"os"
)


func (c *Client) PostCreateVps( slug string ) error {
	_,err := c.post(
		"create-vps",
		CreateVPSRequest{
			Slug: slug,
			World: WorldParams{
				Name: os.Getenv("VALHEIM_WORLD"),
				Password: os.Getenv("VALHEIM_PASSWORD"),
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
