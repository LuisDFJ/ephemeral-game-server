package server

import "net/http"

func (s *Server) DestroyVPSHandler( w http.ResponseWriter, r *http.Request ) {
	vps,err := s.doc.GetDroplet()
	if err != nil {
		http.Error(w, "Unable to connect with VPS API", http.StatusInternalServerError)
		return 
	}

	if !vps.Alive {
		http.Error(w, "VPS is already destroyed", http.StatusConflict)
		return
	}

	err = s.doc.DelVps()
	if err != nil {
		http.Error(w, "Unable to destroy VPS, consider retrying", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
