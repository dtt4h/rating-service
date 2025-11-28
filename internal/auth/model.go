package auth

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Hash  string `json:"-"`
}

type RegisterRequest struct {
	Login    string `json:"login" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
