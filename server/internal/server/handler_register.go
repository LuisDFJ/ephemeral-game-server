package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unicode"
)

func isValidEntry(user string) bool {
	for _,c := range user {
		if !(unicode.IsDigit(c) || unicode.IsLetter(c) || unicode.IsPunct(c) || unicode.IsSymbol(c)) {
			return false
		}
	}
	return true
}

func isValidPassword(pass string) bool {
	f_digit := false
	f_letter := false
	f_upper := false
	f_symbol := false
	for _,c := range pass {
		if unicode.IsDigit(c) { f_digit = true }
		if unicode.IsLetter(c) { f_letter = true }
		if unicode.IsPunct(c) || unicode.IsSymbol(c) { f_symbol = true }
		if unicode.IsUpper(c) { f_upper = true }
	}

	return f_digit && f_letter && f_upper && f_symbol
}

func (r *Registration) isValid() error {
	if len(r.Username) < 6 {
		return fmt.Errorf("Username too short (min length: 6)")
	}
	if len(r.Password) < 10 {
		return fmt.Errorf("Password too short (min length: 10)")
	}
	if !isValidEntry(r.Username) || !isValidEntry(r.Password) {
		return fmt.Errorf("Invalid characters in username or password")
	}
	if !isValidPassword(r.Password) {
		return fmt.Errorf("Password must contain at least a digit, a letter, a symbol, and an uppercase")
	}
	return nil
}

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var reg Registration
	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := reg.isValid(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ott_id := s.db.ValidateOTT(reg.Ott)
	if ott_id == 0 {
		http.Error(w, "Invalid One Time Token", http.StatusBadRequest)
		return
	}

	if err := s.db.CreateUser(reg.Username, reg.Password); err != nil {
		http.Error(w, "Invalid username or registration", http.StatusBadRequest)
		return
	}

	s.db.DeleteOTT(ott_id)
	w.WriteHeader(http.StatusOK)
}
