package request

type CreateUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshToken struct {
	AccessToken string `json:"access_token" validate:"required"`
}
