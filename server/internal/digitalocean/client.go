package digitalocean

import (
	"sync"
	"context"
	"github.com/digitalocean/godo"
)

const VALHEIM_IMAGE string = "valheim-arch-do"
const VALHEIM_DROPLET string = "valheim-server"
const VALHEIM_VOLUME string = "volume-nyc1-valheim"
const REGION string = "nyc1"
const MAX_PRICE float64 = 75.0

type Client struct {
	mu sync.Mutex
	ctx context.Context
	client *godo.Client
}

func New( pat string ) Client {
	return Client{
		ctx: context.Background(),
		client: godo.NewFromToken(pat),
	}
}
