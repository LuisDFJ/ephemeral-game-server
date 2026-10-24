package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/internal/cli"
	"server/internal/server"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

const PORT = ":8080"

func init_server() func() {
	handler := server.New()

	mux := http.NewServeMux()
	mux.Handle("POST /login", http.HandlerFunc(handler.LoginHandler))
	mux.Handle("POST /register", http.HandlerFunc(handler.RegisterHandler))
	mux.Handle("GET /status", handler.AuthMiddleware(handler.StatusHandler))
	mux.Handle("GET /available-sizes", handler.AuthMiddleware(handler.AvailableSizesHandler))
	mux.Handle("POST /create-vps", handler.AuthMiddleware(handler.CreateVPSHandler))
	mux.Handle("POST /destroy-vps", handler.AuthMiddleware(handler.DestroyVPSHandler))
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	srv := http.Server{
		Addr: PORT,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server Starting on %s\n", PORT)
	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
		log.Println("Shuting Down Server")
	}()

	return func() {
		ctx, close := context.WithTimeout(context.Background(), time.Second * 10)
		defer close()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server failed to shut-down: %v", err)
		}
		log.Println("Server Shut Down")
		handler.Close()
	}
}

func _main() {
	shut_down := init_server()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGABRT)

	<-sig

	shut_down()
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	if len(os.Args) != 1 {
		err := godotenv.Load(".env")
		fatal(err)
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

