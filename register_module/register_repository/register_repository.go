package registerrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"mylibary/register_module/entities"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRepository interface {
	HashPasswords(data *entities.Users) error
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

func (conn *SQLconnectDB) HashPasswords(data *entities.Users) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Passwords), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	data.Passwords = string(hash)
	return nil
}

func (conn *SQLconnectDB) RegisterRepository(data *entities.Users) error {
	err := conn.CheckUserRepository(data)
	fmt.Print("Check:", err)
	if err == nil {
		if err := conn.HashPasswords(data); err != nil {
			return err
		}

		fmt.Print(data.FirstName, data.LastName, data.Age, data.Phone, data.Email, data.Passwords)

		_, err := conn.db.Exec("INSERT INTO public.register_db(first_name, last_name, age, phone, email, passwords) "+
			" VALUES ($1, $2, $3, $4, $5, $6)", data.FirstName, data.LastName, data.Age, data.Phone, data.Email, data.Passwords)

		if data.Email == "adminrmuti@gmail.com" {
			_, err := conn.db.Exec("INSERT INTO public.admin_db(first_name, last_name, age, phone, email, passwords) "+
				" VALUES ($1, $2, $3, $4, $5, $6)", data.FirstName, data.LastName, data.Age, data.Phone, data.Email, data.Passwords)
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
		return fmt.Errorf("Register Success")
	}
	return errors.New("Email already exits1")

}
