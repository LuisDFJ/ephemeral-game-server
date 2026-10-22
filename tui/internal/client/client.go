package client

import "os"

type Client struct {
	url string
	jwt []byte
}

func Register(ott string, username string, password string) error {
	c := Client{ url:os.Getenv("VALHEIM_SERVER") }
	_,err := c.post(
		"register",
		RegistrationRequest{
			Username:username,
			Password:password,
			Ott:ott,
		},
		false,
	)
	return err
}

func New(username string, password string) (Client, error) {
	c := Client{ url:os.Getenv("VALHEIM_SERVER") }
	jwt, err := c.post(
		"login",
		CredentialsRequest{
			Username: username,
			Password: password,
		},
		false,
	)
	c.jwt = jwt
	return c,err
}
