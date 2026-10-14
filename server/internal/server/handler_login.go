package server

import (
	"time"
	"net/http"
	"encoding/json"
	jwt "github.com/golang-jwt/jwt/v5"
)

func (s *Server) createToken(cred Credentials) (string, error) {
	expiration := time.Now().Add(5*time.Minute)
	claims := &Claims{
		Username: cred.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(s.key)
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := s.db.AuthUser(cred.Username, cred.Password); err != nil {
		http.Error(w, "Unauthorized: Invalid username or password", http.StatusUnauthorized)
		return
	}

	token,err := s.createToken(cred)
	if err != nil {
		http.Error(w, "Invalid server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/jwt")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

