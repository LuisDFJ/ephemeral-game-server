package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) StatusHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	vps,err := s.doc.GetDroplet()
	if err != nil {
		http.Error(w, "unable to connect with VPS API", http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vps)
}
