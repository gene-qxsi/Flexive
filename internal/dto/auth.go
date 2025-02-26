package dto

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password_hash" binding:"required,min=8"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignUpRequest struct {
	Username     string `json:"username" binding:"required,max=64"`
	Email        string `json:"email" binding:"required,email,max=128"`
	PasswordHash string `json:"password_hash" binding:"required,max=255"`
}
