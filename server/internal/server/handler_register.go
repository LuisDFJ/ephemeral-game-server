package server

import (
	"net/http"
	"encoding/json"
)

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var reg Registration
	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
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
