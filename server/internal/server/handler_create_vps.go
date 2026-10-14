package server

import (
	"net/http"
	"encoding/json"
)

type CreateVPS struct {
	Slug string `json:"slug"`
}

func (s *Server) CreateVPSHandler( w http.ResponseWriter, r *http.Request ) {
	var params CreateVPS
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Incorrect Payload", http.StatusBadRequest)
		return
	}

	vps,err_1 := s.doc.GetDroplet()
	sizes,err_2 := s.doc.GetSizes()
	if err_1 != nil || err_2 != nil {
		http.Error(w, "unable to connect with VPS API", http.StatusInternalServerError)
		return 
	}
	if vps.Alive {
		http.Error(w, "VPS already exists", http.StatusConflict)
		return
	}
	var available bool = false
	for i := range sizes {
		if sizes[i].Slug == params.Slug {
			available = true
			break
		}
	}
	if !available {
		http.Error(w, "Slug not available", http.StatusConflict)
		return
	}

	// TODO: Validate that User Image and Volume Exists, Create Image, Retreive IP and Update DNS

	w.WriteHeader(http.StatusOK)
}
