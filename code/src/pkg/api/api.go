package api

type Host struct {
	Address string `json:"host"`
	Port    int    `json:"port"`
}

type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
