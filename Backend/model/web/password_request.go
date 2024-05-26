package web

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,max=200,min=1"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email" validate:"required,max=200,min=1"`
	OTP      string `json:"otp" validate:"required,max=200,min=1"`
	Password string `json:"new_password" validate:"required,max=200,min=8"`
}
