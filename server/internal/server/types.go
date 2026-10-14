package server

import (
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
