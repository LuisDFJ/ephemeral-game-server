package digitalocean

import (
	"context"
	"os"
	"sync"

	"github.com/digitalocean/godo"
)

const VALHEIM_IMAGE string = "valheim-arch-do"
const VALHEIM_DROPLET string = "valheim-server"
const VALHEIM_VOLUME string = "volume-nyc1-valheim"
const MAX_PRICE float64 = 75.0

type Client struct {
	mu sync.Mutex
	ctx context.Context
	client *godo.Client
	region string
}

func New( pat string ) Client {
	region := os.Getenv("VPS_REGION")
	if region == "" { region = "nyc1" }
	return Client{
		ctx: context.Background(),
		client: godo.NewFromToken(pat),
		region: region,
	}
}
