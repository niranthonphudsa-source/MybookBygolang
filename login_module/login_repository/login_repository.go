package loginrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"mylibary/login_module/entities"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginRepository interface {
	CheckLoginRepository(data *entities.UserLogin) error
	LoginRepo(data *entities.UserLogin) error
}

type SQLconnectDB struct {
	db *sql.DB
}

func NewLoginRepo(db *sql.DB) *SQLconnectDB {
	return &SQLconnectDB{db: db}
}

func (conn *SQLconnectDB) CheckLoginRepository(data *entities.UserLogin) error {
	var storeHash string

	err := conn.db.QueryRow("SELECT passwords FROM public.register_db WHERE email = $1", data.Email).Scan(&storeHash)
	fmt.Print("Email: ", data.Email, "Password ", storeHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return err
	}

	errCheckHash := bcrypt.CompareHashAndPassword([]byte(storeHash), []byte(data.Passwords))
	if errCheckHash != nil {
		return errors.New("Invalid Password")
	}

	return nil
}

func (conn *SQLconnectDB) LoginRepo(data *entities.UserLogin) error {

	errCheck := conn.CheckLoginRepository(data)
	if errCheck != nil {
		return errCheck

	}
	fmt.Println("Error is", errCheck)

	_, err := conn.db.Exec("INSERT INTO public.login_db "+
		"(email, password) VALUES ($1, $2)", data.Email, data.Passwords)
	fmt.Println("Error is", data.Email, data.Passwords)
	if err != nil {
		return errors.New("Login Failed")
	}

	return errors.New("Login Success")
}
