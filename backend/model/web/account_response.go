package web

type AccountResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	NoTelepon int    `json:"no_telepon"`
	Role      string `json:"role"`
}
