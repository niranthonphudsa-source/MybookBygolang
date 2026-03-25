package loginrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"mylibary/login_module/entities"

	_ "github.com/lib/pq"
)

type UserLoginRepository interface {
	CheckLoginRepository(email string, passwords string) error
	LoginRepo(data *entities.UserLogin) error
}

type SQLconnectDB struct {
	db *sql.DB
}

func NewLoginRepo(db *sql.DB) *SQLconnectDB {
	return &SQLconnectDB{db: db}
}

func (conn *SQLconnectDB) CheckLoginRepository(Email string, Passwords string) error {
	var b entities.UserLogin

	err := conn.db.QueryRow("SELECT email, passwords FROM public.register_db WHERE email = $1 AND passwords = $2;", Email, Passwords).Scan(&b.Email, &b.Passwords)
	fmt.Print(Email, Passwords, " . ", err)
	if err == sql.ErrNoRows {
		return errors.New("Have to register, No user")
	}
	return nil
}

func (conn *SQLconnectDB) LoginRepo(data *entities.UserLogin) error {
	errCheck := conn.CheckLoginRepository(data.Email, data.Passwords)
	fmt.Println("Error is", errCheck)
	if errCheck == nil {
		_, err := conn.db.Exec("INSERT INTO public.login_db "+
			"(email, password) VALUES ($1, $2)", data.Email, data.Passwords)
		fmt.Println("Error is", data.Email, data.Passwords)
		if err == nil {
			return errors.New("Login Success!!!")
		}
		return errors.New("Have to register, No user")
	}
	return errors.New("Have to register, No user")
}
