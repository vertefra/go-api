package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/vertefra/go-fiber-tutorial/book"
	"github.com/vertefra/go-fiber-tutorial/database"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World")
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		fmt.Println(err)
		panic("failed connection with database ")
	}

	fmt.Println("Database connected")

	// automigration
	// takes the sbook struct and create a table

	db := database.DBConn.AutoMigrate(&book.Book{})

	fmt.Println("Database migrated ", db)

}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	defer database.DBConn.Close()

	app.Listen(3000)
}
