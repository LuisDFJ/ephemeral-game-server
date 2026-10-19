package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) AvailableSizesHandler(w http.ResponseWriter, r *http.Request) {
	sizes,err := s.doc.GetSizes()
	if err != nil {
		http.Error(w, "unable to connect with VPS API", http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sizes)
	log.Printf("VPS sizes requested by %s", r.Context().Value("user"))
}
