package loginrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"mylibary/login_module/entities"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	var storedHash string
	var Status string
	query := `
        SELECT 
            r.passwords, 
            CASE WHEN a.email IS NOT NULL THEN 'admin' ELSE 'user' END as status
        FROM public.register_db r
        LEFT JOIN public.admin_db a ON r.email = a.email
        WHERE r.email = $1
    `
	err := conn.db.QueryRow(query, data.Email).Scan(&storedHash, &Status)
	fmt.Print("Email: ", data.Email, "Password ", storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return err
	}

	errCheckHash := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(data.Passwords))
	if errCheckHash != nil {
		return errors.New("Invalid Password")
	}
	data.Status = Status
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

	claims := jwt.MapClaims{
		"email":  data.Email,
		"status": data.Status,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err
	}
	return errors.New("Login Success" + t)
}
