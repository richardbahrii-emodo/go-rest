package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/richardbahrii-emodo/go-rest/controllers"
)

func AddTodoGroup(prefix fiber.Router) fiber.Router {
	todos := prefix.Group("/todos")

	todos.Get("/", controllers.GetAllTodo)
	todos.Post("/", controllers.CreateTodo)
	todos.Put("/:id", controllers.UpdateTodo)
	todos.Delete("/:id", controllers.DeleteTodo)

	return todos
}
