package server

import (
	"log"
	"sync"
	"crypto/rsa"
	"crypto/rand"
	"server/internal/server/db"
	cf "server/internal/claudflare"
	do "server/internal/digitalocean"
)


func fatal(s string) {
	log.Fatal(s)
}

type World struct {
	mu sync.Mutex
	alive bool
	name string
	password string
}

func (w *World) update(name string, password string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.alive = true
	w.name = name
	w.password = password
}

func (w *World) delete() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.alive = false
	w.name = ""
	w.password = ""
}

type Server struct {
	key *rsa.PrivateKey
	pub *rsa.PublicKey
	doc do.Client
	cfc cf.Client
	db *db.DB
	world World
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
	log.Println("Closing Server Handler")
	s.db.Close()
}
