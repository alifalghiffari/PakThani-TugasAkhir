package web

type UserResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	NoTelepon int    `json:"no_telepon"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}
