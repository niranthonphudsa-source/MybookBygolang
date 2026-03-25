package registerrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"mylibary/register_module/entities"

	_ "github.com/lib/pq"
)

type RegisterRepository interface {
	CheckUserRepository(data *entities.Users) error
	RegisterRepository(data *entities.Users) error
}

type SQLconnectDB struct {
	db *sql.DB
}

func NewRegisterRepo(db *sql.DB) *SQLconnectDB {
	return &SQLconnectDB{db: db}
}

func (conn *SQLconnectDB) CheckUserRepository(data *entities.Users) error {
	var b entities.Users
	err := conn.db.QueryRow("SELECT email FROM public.register_db WHERE email = $1", data.Email).Scan(b.Email)
	if err == sql.ErrNoRows {
		return nil
	}
	return errors.New("Email already exits")
}

func (conn *SQLconnectDB) RegisterRepository(data *entities.Users) error {
	err := conn.CheckUserRepository(data)
	fmt.Print("Check:", err)
	if err == nil {

		fmt.Print(data.FirstName, data.LastName, data.Age, data.Phone, data.Email, data.Passwords)

		_, err := conn.db.Exec("INSERT INTO public.register_db(first_name, last_name, age, phone, email, passwords) "+
			" VALUES ($1, $2, $3, $4, $5, $6)", data.FirstName, data.LastName, data.Age, data.Phone, data.Email, data.Passwords)

		if data.Email == "adminrmuti@gmail.com" && data.Passwords == "adminrmuti1234" {
			_, err := conn.db.Exec("INSERT INTO public.admin_db(first_name, last_name, age, phone, email, passwords) "+
				" VALUES ($1, $2, $3, $4, $5, $6)", data.FirstName, data.LastName, data.Age, data.Phone, data.Email, data.Passwords)
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
	return errors.New("Email already exits")

}
