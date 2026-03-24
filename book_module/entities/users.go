package entities

type UserReq struct {
	UserId    int    `json:"User_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       string `json:"age"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
