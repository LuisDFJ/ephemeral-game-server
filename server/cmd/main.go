package main

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"server/internal/server"
	"server/internal/client"
	"server/internal/cli"
)

const PORT = ":8080"

func init_server() func() {
	server := server.New()

	mux := http.NewServeMux()
	mux.Handle("POST /login", http.HandlerFunc(server.LoginHandler))
	mux.Handle("POST /register", http.HandlerFunc(server.RegisterHandler))
	mux.Handle("GET /status", server.AuthMiddleware(server.StatusHandler))
	mux.Handle("GET /available-sizes", server.AuthMiddleware(server.AvailableSizesHandler))
	mux.Handle("POST /create-vps", server.AuthMiddleware(server.CreateVPSHandler))
	mux.Handle("POST /destroy-vps", server.AuthMiddleware(server.DestroyVPSHandler))

	fmt.Printf("Server Starting on %s\n", PORT)
	go func() {
		if err := http.ListenAndServe(PORT, mux); err != nil {
			fmt.Printf("Server failed to start: %v", err)
			os.Exit(1)
		}
	}()

	return func() { server.Close() }
}

func _main() {
	//ott,err := cli.CreateOTT()
	//fatal(err)

	cleanup := init_server()
	defer cleanup()

	//client.Register(ott,"user","password",PORT)

	c := client.New("user", "password", PORT)
	res,_ := c.GetStatus()
	fmt.Println(string(res))
	res,_ = c.GetSizes()
	fmt.Println(string(res))
	res,_ = c.PostCreateVps("x-2vcpu-4gb", "World", "password.1234")
	fmt.Println(string(res))
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	if len(os.Args) != 1 {
		switch os.Args[1] {
		case "gen-ott":
			token,err := cli.CreateOTT()
			fatal(err)
			fmt.Print(token)
		case "create-user":
			if len(os.Args) == 4 {
				user,pass := os.Args[2], os.Args[3]
				err := cli.CreateUser(user,pass)
				fatal(err)
			} else {
				fmt.Println("not enought arguments")
			}
		case "auth-user":
			if len(os.Args) == 4 {
				user,pass := os.Args[2], os.Args[3]
				err := cli.AuthUser(user,pass)
				fatal(err)
			} else {
				fmt.Println("not enought arguments")
			}
		}
		return
	}
	_main()
}

