package client

type Client struct {
	url string
	jwt []byte
}

func Register(ott string, username string, password string, url string) error {
	c := Client{ url:url }
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

func New(username string, password string, url string) (Client, error) {
	c := Client{ url:url }
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
