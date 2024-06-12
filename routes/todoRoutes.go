package routes

import (
	"todo-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    app.Post("/todos", handlers.CreateTodoList)
    app.Get("/todos/:id", handlers.GetTodoList)
    app.Get("/todos", handlers.GetAllTodoLists)
    app.Put("/todos/:id", handlers.UpdateTodoList)
    app.Delete("/todos/:id", handlers.DeleteTodoList)
    app.Patch("/todos/:id/completed", handlers.MarkAsCompleted)
}
