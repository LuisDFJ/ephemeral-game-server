package do

import (
	"context"
	"github.com/digitalocean/godo"
)

const VALHEIM_IMAGE string = "valheim-arch-do"
//const VALHEIM_DROPLET string = "valheim-server"
const VALHEIM_DROPLET string = "arch-s-1vcpu-512mb-10gb-nyc1"

type Client struct {
	ctx context.Context
	client *godo.Client
}

func New( pat string ) Client {
	c := Client{
		ctx: context.Background(),
		client: godo.NewFromToken(pat),
	}
	return c
}
