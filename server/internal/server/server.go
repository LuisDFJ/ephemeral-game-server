package server

import (
	"os"
	"fmt"
	"crypto/rand"
	"crypto/rsa"
	cf "server/internal/claudflare"
	do "server/internal/digitalocean"
	"server/internal/server/db"
)


func fatal(s string) {
	fmt.Println(s)
	os.Exit(1)
}

type Server struct {
	key *rsa.PrivateKey
	pub *rsa.PublicKey
	doc do.Client
	cfc cf.Client
	db *db.DB
}

func New() Server {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil { fatal("error generating key") }
	dot,cft,err := load_tokens()
	if err != nil { fatal(err.Error()) }
	db,err := db.New()
	if err != nil { fatal(err.Error()) }
	return Server{
		key: key,
		pub: &key.PublicKey,
		doc: do.New(dot),
		cfc: cf.New(cft),
		db: db,
	}
}

func (s *Server) Close() {
	fmt.Println("Closing Server")
	s.db.Close()
}
