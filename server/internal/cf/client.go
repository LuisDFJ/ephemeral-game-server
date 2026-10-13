package cf

type Client struct {
	pat string
}

func New( pat string ) Client {
	c := Client{
		pat: pat,
	}
	return c
}

