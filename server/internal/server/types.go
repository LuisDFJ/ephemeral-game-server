package server

import (
	do "server/internal/digitalocean"
	jwt "github.com/golang-jwt/jwt/v5"
)

type Registration struct {
	Ott string `json:"ott"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type CreateVPS struct {
	Slug string `json:"slug"`
	World WorldParams `json:"world"`
}

type Status struct {
	Vps do.VPS `json:"vps"`
	World WorldParams `json:"world"`
}

type WorldParams struct {
	Name string `json:"name"`
	Password string `json:"password"`
}
