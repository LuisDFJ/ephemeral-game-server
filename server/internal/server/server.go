package server

import (
	"sync"
	"os"
	"fmt"
	"crypto/rsa"
	"crypto/rand"
)


func fatal(s string) {
	fmt.Println(s)
	os.Exit(1)
}

type Server struct {
	mu sync.Mutex
	key *rsa.PrivateKey
	pub *rsa.PublicKey
}

func New() Server {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil { fatal("error generating key") }
	return Server{
		key: key,
		pub: &key.PublicKey,
	}
}

