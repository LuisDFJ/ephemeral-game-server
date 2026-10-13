package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"server/internal/server"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Succesfully loged in"))
}

const PORT = ":8080"

func init() {
	server := server.New()
	//defer server.Close()
	mux := http.NewServeMux()
	mux.Handle("POST /login", http.HandlerFunc(server.LoginHandler))
	mux.Handle("GET /status", server.AuthMiddleware(server.StatusHandler))
	mux.Handle("GET /test", server.AuthMiddleware(ProtectedHandler))

	fmt.Printf("Server Starting on %s\n", PORT)
	go func() {
		if err := http.ListenAndServe(PORT, mux); err != nil {
			fmt.Printf("Server failed to start: %v", err)
			os.Exit(1)
		}
	}()
}


func fatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type Client struct {
	jwt []byte
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(username string, password string) Client {
	c := Client{}
	jwt, err := c.post("login",Credentials{Username: "admin", Password: "password"})
	fatal(err)
	c.jwt = jwt
	return c
}

func api_url(s string) string {
	return fmt.Sprintf("http://localhost%s/%s", PORT, s)
}


func (c *Client) get(slug string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, api_url(slug), nil)
	if err != nil { return nil,err }
	req.Header.Add("Authorization", "Bearer " + string(c.jwt))
	response, err := http.DefaultClient.Do(req)
	if err != nil { return nil,err }
	defer response.Body.Close()
	body,err := io.ReadAll(response.Body)
	if err != nil { return nil,err }
	return body,nil
}

func (c *Client) post(slug string, b any) ([]byte, error) {
	jsonBody,err := json.Marshal(b)
	if err != nil { return nil,err }

	req, err := http.NewRequest(http.MethodPost, api_url(slug), bytes.NewBuffer(jsonBody))
	if err != nil { return nil,err }
	req.Header.Add("Content-Type", "application/json")
	if c.jwt != nil {
		req.Header.Add("Authorization", "Bearer " + string(c.jwt))
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil { return nil,err }
	defer response.Body.Close()
	body,err := io.ReadAll(response.Body)
	if err != nil { return nil,err }
	return body,nil
}

func (c *Client) GetStatus() ([]byte, error) {
	return c.get("status")
}

func main() {
	c := New("admin", "password")
	res,_ := c.GetStatus()
	fmt.Println(string(res))
}

