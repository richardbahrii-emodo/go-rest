package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/richardbahrii-emodo/go-rest/config"
	"github.com/richardbahrii-emodo/go-rest/database"
	"github.com/richardbahrii-emodo/go-rest/routes"
)

func main() {
	err := initApplication()

	if err != nil {
		println("Appliction was crashed. Sorry for that.")
		panic(err)
	}
}

func initApplication() error {
	err := config.LoadENV()
	if err != nil {
		return err
	}

	err = database.InitDB()
	if err != nil {
		return err
	}

	defer database.CloseDb()

	app := fiber.New()

	prefix := app.Group("/api")

	routes.AddTodoGroup(prefix)

	app.Listen(":" + os.Getenv("PORT"))

	defer app.Shutdown()

	return nil
}
