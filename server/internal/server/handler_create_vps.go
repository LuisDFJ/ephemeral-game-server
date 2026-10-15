package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	do "server/internal/digitalocean"
	userdata "server/internal/user_data"
	"time"
)

type CreateVPS struct {
	Slug string `json:"slug"`
	WorldName string `json:"worldname"`
	WorldPswd string `json:"worldpswd"`
}

func (s *Server) CreateVPSHandler( w http.ResponseWriter, r *http.Request ) {
	var params CreateVPS
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Incorrect Payload", http.StatusBadRequest)
		return
	}

	vps,err_a := s.doc.GetDroplet()
	sizes,err_b := s.doc.GetSizes()
	image,err_c := s.doc.GetImage()
	volume,err_d := s.doc.GetVolume()
	key,err_e := s.doc.GetKey()

	if err_a != nil || err_b != nil || err_c != nil || err_d != nil || err_e != nil {
		http.Error(w, "Unable to connect with VPS API", http.StatusInternalServerError)
		return 
	}

	if vps.Alive {
		http.Error(w, "VPS already exists", http.StatusConflict)
		return
	}

	size,ok := is_size_available(params.Slug, sizes)
	if !ok {
		http.Error(w, "Slug not available", http.StatusConflict)
		return
	}

	if !image.Available {
		http.Error(w, "Image not available", http.StatusConflict)
		return
	}

	if !volume.Available {
		http.Error(w, "Volume not available", http.StatusConflict)
		return
	}

	if !key.Available {
		http.Error(w, "Key not available", http.StatusConflict)
		return
	}

	if len(params.WorldName) == 0 || len(params.WorldPswd) < 8 {
		http.Error(w, "World Name is empty or Password is too weak", http.StatusConflict)
		return
	}

	ud := userdata.Params{
		User: "admin",
		SSHKey: key.PublicKey,
		Volume: volume.Name,
		WorldName: params.WorldName,
		WorldPswd: params.WorldPswd,
	}

	fmt.Println(ud.Gen())

	err := s.doc.PostVps(ud, size, image, volume, key)
	if err != nil {
		http.Error(w, "Could not create VPS", http.StatusConflict)
		return
	}

	// Waiting for IP
	ip,err := wait_for_ip(time.Second*4, 15, s)
	if err != nil {
		http.Error(w, "Could not get VPS IP", http.StatusConflict)
		return
	}
	fmt.Println(ip)

	// TODO: Validate that User Image and Volume Exists, Create VPS, Retreive IP and Update DNS

	w.WriteHeader(http.StatusOK)
}

func wait_for_ip(t time.Duration, retries int, s *Server) (string,error) {
	for _ = range retries {
		time.Sleep(t)
		droplet,err := s.doc.GetDroplet()
		if err != nil { return "", err }
		if droplet.Ip == "" { continue }
		return droplet.Ip, nil
	}
	return "", fmt.Errorf("Couldn't get VPS' IP")
}

func is_size_available(slug string, sizes []do.Size) (do.Size, bool) {
	var available bool = false
	var size do.Size
	for i := range sizes {
		if sizes[i].Slug == slug {
			available = true
			size = sizes[i]
			break
		}
	}
	return size, available
}
