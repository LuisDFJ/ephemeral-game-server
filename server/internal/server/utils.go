package server

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func load_one_time_tokens() []string {
	tokens := []string{}
	content,err := os.ReadFile("token.env")
	if err == nil {
		for _,s := range strings.Split(string(content), "\n") {
			tokens = append(tokens, s)
		}
	}
	return tokens
}

func save_one_time_tokens(tokens []string) {
	err := os.WriteFile("tokens.env", []byte(strings.Join(tokens,"\n")), 0644)
	if err != nil {
		fmt.Printf("Error saving tokens.env: %s", err.Error())
	}
}

func load_tokens() (string,string,error) {
	var dot,cft string
	err := godotenv.Load(".env")
	if err != nil { return dot,cft,err }
	dot = os.Getenv("DO_TOKEN")
	cft = os.Getenv("CF_TOKEN")
	if dot == "" || cft == "" {
		return dot,cft,fmt.Errorf("Unable to retreive API tokens")
	}
	return dot,cft,nil
}
