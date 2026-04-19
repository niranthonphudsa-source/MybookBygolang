package entities

type BrrowBookRepo struct{
	Borrow_Id	int 	`json:"brrow_id"`
	Book_id		int 	`json:"book_id"`
	User_id 	int 	`json:"user_id"`
}

type BrrowBookResponse struct{
	Borrow_Id	int 	
	Book_id		int 	
	User_id 	int 
	Message 	string
}


func (br *BrrowBookRepo) ToRespose() BrrowBookResponse{
	return BrrowBookResponse{
		Borrow_Id: br.Borrow_Id,
		Book_id: br.Book_id,
		User_id: br.User_id,
		Message: "BORROWBOOK SUCCESS!!",
	}
}
