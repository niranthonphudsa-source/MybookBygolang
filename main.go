package main

import (
	"fmt"
	"log"
	"os"

	bookcontroller "mylibary/book_module/book_controller"
	bookrepository "mylibary/book_module/book_repository"
	bookusecase "mylibary/book_module/book_usecase"
	registercontroller "mylibary/register_module/register_controller"
	registerrepository "mylibary/register_module/register_repository"
	registerusecase "mylibary/register_module/register_usecase"

	"golang.org/x/crypto"
	"github.com/golang-jwt/jwt/v4"

	"mylibary/configs"
	"mylibary/pkg/database"
	"mylibary/middleware/checkmiddleware"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Print("Hello Mybooks")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cgf := configs.GetConnectServer()
	db := database.ConnecDB(cgf)

	registerRepo := registerrepository.NewRegisterRepo(db)
	registerUsecase := registerusecase.NewRegisterUsecaseImplrepo(registerRepo)
	registerController := registercontroller.NewRegisterController(registerUsecase)

	bookRepo := bookrepository.NewBooksRepository(db)
	bookusecase := bookusecase.NewBookUsecase(bookRepo)
	bookcontroller := bookcontroller.NewBookController(bookusecase)

	app := fiber.New()
	app.Use("/book", checkmiddleware.A)
	fiberPort := os.Getenv("FiberPort")

	// app.Get("/api/config", configs.GetSecretkey)
	app.Post("/register", registerController.RegisterController)

	app.Get("/book", bookcontroller.GetBookAllController)
	app.Get("/book/:book_id", bookcontroller.GetBookIdController)

	//
	app.Post("/book", bookcontroller.CreateBookController)
	app.Delete("/book/:book_id", bookcontroller.DeleteBookIdController)
	app.Put("/book/:book_id", bookcontroller.UpdateBookIdController)

	app.Listen(":" + fiberPort)
}
