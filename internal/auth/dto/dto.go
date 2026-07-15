package dto
//DTO stands for Data Transfer Object.
//it as a box that carries data between the client and your server.
//without it gin doesn't know where to store data from the request body, so we need to define a struct that matches the expected JSON structure of the request body.

type RegisterRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type LoginResponse struct {
	Token string `json:"token"`
	User UserResponse `json:"user"`
}
type UserResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}