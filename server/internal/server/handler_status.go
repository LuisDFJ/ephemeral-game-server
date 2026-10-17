package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) StatusHandler(w http.ResponseWriter, r *http.Request) {
	vps,err := s.doc.GetDroplet()

	if err != nil {
		http.Error(w, "unable to connect with VPS API", http.StatusInternalServerError)
		return 
	}

	s.world.mu.Lock()
	defer s.world.mu.Unlock()
	
	if s.world.alive {
		http.Error(w, "Vps is Offline", http.StatusInternalServerError)
		return 
	}

	status := Status{
		Vps: vps,
		World: WorldParams{
			Name: s.world.name,
			Password: s.world.password,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
