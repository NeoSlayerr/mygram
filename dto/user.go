package dto

type NewUserRequest struct {
	Username    string `json:"username" valid:"required~username cannot be empty" example:"user1"`
	Email    string `json:"email" valid:"required~email cannot be empty, email~Invalid email format" example:"user1@gmail.com"`
	Password string `json:"password" valid:"required~password cannot be empty, minstringlength(6)~Password has to have a minimum length of 6 characters" example:"abc123"`
	Age int `json:"age" valid:"required~age cannot be empty, range(9|100)~age must greater than 8" example:"17"`
}

type NewLoginRequest struct {
	Email    string `json:"email" valid:"required~email cannot be empty, email~Invalid email format" example:"user1@gmail.com"`
	Password string `json:"password" valid:"required~password cannot be empty, minstringlength(6)~Password has to have a minimum length of 6 characters" example:"abc123"`
}

type NewUserResponse struct {
	Result     string `json:"result" example:"success"`
	StatusCode int    `json:"statusCode" example:"201"`
	Message    string `json:"message" example:"registered successfully"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Result     string        `json:"result" example:"success"`
	StatusCode int           `json:"statusCode" example:"200"`
	Message    string        `json:"message" example:"logged in successfully"`
	Data       TokenResponse `json:"data"`
}