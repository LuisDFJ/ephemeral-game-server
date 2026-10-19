package client

type CredentialsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationRequest struct {
	Ott string `json:"ott"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateVPSRequest struct {
	Slug string `json:"slug"`
	World WorldParams `json:"world"`
}
type WorldParams struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type StatusResponse struct {
	Vps Vps `json:"vps"`
	World WorldParams `json:"world"`

}
type Vps struct {
	Id int `json:"id"`
	Alive bool `json:"alive"`
	Name string `json:"name"`
	Ip string `json:"ip"`
	Specs string `json:"specs"`
}

type JWTResponse []byte

type SizesResponse []Size
type Size struct {
	Slug string `json:"slug"`
	Disk int `json:"disk"`
	Mem int `json:"mem"`
	VCpu int `json:"vcpu"`
	PriceMonth float64 `json:"pricemonth"`
	PriceHour float64 `json:"pricehour"`
}
