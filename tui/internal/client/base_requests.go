package client

import (
	"io"
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
)

func (c *Client) api_url(s string) string {
	return fmt.Sprintf("%s/%s", c.url, s)
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
	var e error
	if response.StatusCode != http.StatusOK {
		e = fmt.Errorf("%v: %s", response.Status, body)
	}
	return body,e
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
		e = fmt.Errorf("%v: %s", response.Status, body)
	}
	return body,e
}
