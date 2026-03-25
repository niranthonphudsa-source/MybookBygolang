package main

import (
	"fmt"
	"log"
	"os"

	bookcontroller "mylibary/book_module/book_controller"
	bookrepository "mylibary/book_module/book_repository"
	bookusecase "mylibary/book_module/book_usecase"

	"mylibary/configs"
	"mylibary/pkg/database"

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

	bookRepo := bookrepository.NewBooksRepository(db)
	bookusecase := bookusecase.NewBookUsecase(bookRepo)
	bookcontroller := bookcontroller.NewBookController(bookusecase)

	app := fiber.New()
	fiberPort := os.Getenv("FiberPort")

	// app.Get("/api/config", configs.GetSecretkey)

	app.Get("/book", bookcontroller.GetBookAllController)
	app.Get("/book/:book_id", bookcontroller.GetBookIdController)
	app.Post("/book", bookcontroller.CreateBookController)
	app.Delete("/book/:book_id", bookcontroller.DeleteBookIdController)
	app.Put("/book/:book_id", bookcontroller.UpdateBookIdController)
	app.Listen(":" + fiberPort)
}
