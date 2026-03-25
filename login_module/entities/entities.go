package entities

type UserLogin struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Passwords string `json:"passwords"`
}
