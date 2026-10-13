package server

import (
	"sync"
	"os"
	"fmt"
	"crypto/rsa"
	"crypto/rand"
	"server/internal/do"
	"server/internal/cf"
)


func fatal(s string) {
	fmt.Println(s)
	os.Exit(1)
}

type Server struct {
	mu sync.Mutex
	key *rsa.PrivateKey
	pub *rsa.PublicKey
	ott []string
	doc do.Client
	cfc cf.Client
}

func New() Server {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil { fatal("error generating key") }
	dot,cft,err := load_tokens()
	if err != nil { fatal(err.Error()) }
	return Server{
		key: key,
		pub: &key.PublicKey,
		ott: load_one_time_tokens(),
		doc: do.New(dot),
		cfc: cf.New(cft),
	}
}

func (s *Server) Close() {
	save_one_time_tokens(s.ott)
}
