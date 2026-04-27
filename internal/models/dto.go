package models

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=4,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
}

type LoginRequest struct {
	// Identifier can be either an Email or a Username
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
