package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	UserID  int    `json:"userId"`
	IsAdmin bool   `json:"isAdmin"`
}
