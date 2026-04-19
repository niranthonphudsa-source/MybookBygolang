package entities

type Users struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Passwords string `json:"passwords"`
	Status    string `json:"status"`
}


type UsersRes struct {
	FirstName string 
	LastName  string 
	Message	  string 
}

func (user *Users) ToRespose() UsersRes {
	return UsersRes{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Message: "REGISTER SUCCESS",
	}
}