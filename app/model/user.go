package model

type UserRequest struct {
	Name     string `json:"name" validate:"required,max=60"`
	Phone    string `json:"phone_number" validate:"required,max=13,min=10"`
	Password string `json:"password" validate:"required,min=6"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"-" `
}

type Login struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserUpdate struct {
	Name string `json:"name" validate:"required,max=60"`
}

// regexp=^08[1-9][0-9]{10,13}$
// regexp=^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{6,}$
