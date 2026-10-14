package client

import (
	"io"
	"fmt"
	"log"
	"bytes"
	"net/http"
	"encoding/json"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

type Client struct {
	port string
	jwt []byte
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Registration struct {
	Ott string `json:"ott"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(ott string, username string, password string, port string) error {
	c := Client{ port:port }
	res, err := c.post("register",Registration{Username:username, Password:password, Ott:ott}, false)
	fatal(err)
	fmt.Println(string(res))
	return nil
}

func New(username string, password string, port string) Client {
	c := Client{ port:port }
	jwt, err := c.post("login",Credentials{Username: username, Password: password}, false)
	fatal(err)
	c.jwt = jwt
	return c
}

func (c *Client) api_url(s string) string {
	return fmt.Sprintf("http://localhost%s/%s", c.port, s)
}


func (c *Client) get(slug string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.api_url(slug), nil)
	if err != nil { return nil,err }
	req.Header.Add("Authorization", "Bearer " + string(c.jwt))
	response, err := http.DefaultClient.Do(req)
	if err != nil { return nil,err }
	defer response.Body.Close()
	body,err := io.ReadAll(response.Body)
	if err != nil { return nil,err }
	return body,nil
}

func (c *Client) post(slug string, b any, jwt bool) ([]byte, error) {
	jsonBody,err := json.Marshal(b)
	if err != nil { return nil,err }

	req, err := http.NewRequest(http.MethodPost, c.api_url(slug), bytes.NewBuffer(jsonBody))
	if err != nil { return nil,err }
	req.Header.Add("Content-Type", "application/json")
	if jwt {
		req.Header.Add("Authorization", "Bearer " + string(c.jwt))
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil { return nil,err }
	defer response.Body.Close()
	body,err := io.ReadAll(response.Body)
	if err != nil { return nil,err }
	var e error
	if response.StatusCode != http.StatusOK {
		e = fmt.Errorf(response.Status)
	}
	return body,e
}

func (c *Client) GetStatus() ([]byte, error) {
	return c.get("status")
}

func (c *Client) GetSizes() ([]byte, error) {
	return c.get("available-sizes")
}

