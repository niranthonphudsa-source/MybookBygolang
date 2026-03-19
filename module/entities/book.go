package entities

type Books struct {
	Book_id        int    `json:"book_id"`
	Book_name      string `json:"book_name"`
	Book_author    string `json:"book_author"`
	Adminupdate_id string `json:"adminupdate_id"`
}
