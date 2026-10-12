package main

import (
	"os"
	"fmt"
	"net/http"
	"server/internal/server"
)




func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Succesfully loged in"))
}


const PORT = ":8080"

func main() {
	server := server.New()
	mux := http.NewServeMux()
	mux.Handle("POST /login", http.HandlerFunc(server.LoginHandler))
	mux.Handle("GET /status", server.AuthMiddleware(ProtectedHandler) )

	fmt.Printf("Server Starting on %s\n", PORT)
	if err := http.ListenAndServe(PORT, mux); err != nil {
		fmt.Printf("Server failed to start: %v", err)
		os.Exit(1)
	}
}

