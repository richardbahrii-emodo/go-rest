package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/richardbahrii-emodo/go-rest/config"
	"github.com/richardbahrii-emodo/go-rest/controllers"
	"github.com/richardbahrii-emodo/go-rest/database"
	"github.com/richardbahrii-emodo/go-rest/routes"
)

func main() {
	err := initApplication()

	if err != nil {
		fmt.Println("Appliction was crashed. Sorry for that.")
		panic(err)
	}
}

func initApplication() error {
	err := config.LoadENV()
	if err != nil {
		return err
	}

	db := database.InitDatabase("mongo")
	app := fiber.New()

	prefix := app.Group("/api")

	mainHandler := controllers.NewHandlerWithDb(db)

	routes.AddTodoGroup(prefix, mainHandler)

	gracefullyShutdown(app, db)

	if err = app.Listen(":" + os.Getenv("PORT")); err != nil {
		return err
	}

	return nil
}

func gracefullyShutdown(app *fiber.App, db database.Database) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("Gracefully shutdown.")

		defer app.Shutdown()
		defer db.Close()

		fmt.Println("Server stopped.")
	}()
}
