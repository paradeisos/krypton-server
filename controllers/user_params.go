package controllers

type UserRegisterParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
