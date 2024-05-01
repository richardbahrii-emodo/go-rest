package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/richardbahrii-emodo/go-rest/controllers"
)

func AddTodoGroup(prefix fiber.Router, mainHandler *controllers.HandlerWithDb) fiber.Router {
	todos := prefix.Group("/todos")

	todos.Get("/", mainHandler.GetAllTodo)
	todos.Post("/", mainHandler.CreateTodo)
	todos.Put("/:id", mainHandler.UpdateTodo)
	todos.Delete("/:id", mainHandler.DeleteTodo)

	return todos
}
