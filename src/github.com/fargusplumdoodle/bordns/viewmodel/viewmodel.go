package viewmodel

type Arecord struct {
	IP   string `json:"ip"`
	FQDN string `json:"fqdn"`
}

type Zone struct {
	Name    string    `json:"zone"`
	Domains []Arecord `json:"domains"`
}
