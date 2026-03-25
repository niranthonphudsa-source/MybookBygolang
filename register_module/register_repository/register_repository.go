package registerrepository

import (
	"database/sql"
	"fmt"
	"mylibary/register_module/entities"

	_ "github.com/lib/pq"
)

type RegisterRepository interface {
	RegisterRepository(data *entities.Users) error
}

type SQLconnectDB struct {
	db *sql.DB
}

func NewRegisterRepo(db *sql.DB) *SQLconnectDB {
	return &SQLconnectDB{db: db}
}

func (conn *SQLconnectDB) RegisterRepository(data *entities.Users) error {

	fmt.Print(data.FirstName, data.LastName, data.Age, data.Phone, data.Email, data.Passwords)
	_, err := conn.db.Exec("INSERT INTO public.register_db(first_name, last_name, age, phone, email, passwords) "+
		" VALUES ($1, $2, $3, $4, $5, $6)", data.FirstName, data.LastName, data.Age, data.Phone, data.Email, data.Passwords)

	if err != nil {
		return err
	}
	return nil
}
