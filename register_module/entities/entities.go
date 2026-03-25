package entities

type Users struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Passwords string `json:"passwords"`
}
