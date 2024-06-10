package web

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,max=20,min=1"`
	Password string `json:"password" validate:"required,max=20,min=8"`
}
