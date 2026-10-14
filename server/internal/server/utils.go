package server

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

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
