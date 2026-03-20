package main

import (
	"fmt"
	"log"
	"os"

	"mylibary/configs"
	"mylibary/module/createbook/createcontroller"
	"mylibary/module/createbook/createrepository"
	"mylibary/module/createbook/createusecase"
	"mylibary/module/getbook/controller"
	"mylibary/module/getbook/repository"
	"mylibary/module/getbook/usecase"
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
	db := database.ConnecDB(&cgf.PostgresSQL)

	bookRepo := repository.GetBooksRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookController := controller.NewBookController(bookUsecase)

	createRepo := createrepository.GetDbCreateBook(db)
	createUsecase := createusecase.NewCreateBookUsecase(createRepo)
	createController := createcontroller.NewCreateBookController(createUsecase)

	app := fiber.New()
	fiberPort := os.Getenv("FiberPort")

	app.Get("/api/config", configs.GetSecretkey)
	app.Get("/book", bookController.GetBooksHandler)
	app.Post("/book", createController.CreateBooksHandler)
	// booksRepo := adapter.NewPostgrestBookRepo(db)
	// createBookService := createbook.NewBookService(booksRepo)
	// booksHandler := adapter.NewHttpCreateBookHandler(createBookService)

	// app.Post("/createbook", booksHandler.CreateBook)
	app.Listen(":" + fiberPort)
}
