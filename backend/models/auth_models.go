package models

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupModel struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
