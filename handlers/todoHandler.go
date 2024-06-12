package handlers

import (
	"todo-app/database"
	"todo-app/models"

	// "todo-app/types"

	"github.com/gofiber/fiber/v2"
)

// CreateTodoList godoc
//	@Summary		Create a new todo list
//	@Description	Create a new todo list with tasks
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		types.Payload	true	"Todo list"
//	@Success		201		{object}	models.Todo
//	@Failure		400		{object}	types.ErrorResponse
//	@Failure		500		{object}	types.ErrorResponse
//	@Router			/todos [post]
func CreateTodoList(c *fiber.Ctx) error {
    todo := new(models.Todo)
    if err := c.BodyParser(todo); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := database.DB.Create(&todo).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not create todo"})
    }

    return c.Status(201).JSON(todo)
}

// GetTodoList godoc
//	@Summary		Get a todo list by ID
//	@Description	Get a todo list by ID
//	@Tags			todos
//	@Produce		json
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	models.Todo
//	@Failure		404	{object}	types.ErrorResponse
//	@Router			/todos/{id} [get]
func GetTodoList(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo models.Todo

    if err := database.DB.First(&todo, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
    }

    return c.JSON(todo)
}

// GetAllTodoLists godoc
//	@Summary		Get all todo lists
//	@Description	Get all todo lists
//	@Tags			todos
//	@Produce		json
//	@Success		200	{array}	models.Todo
//	@Router			/todos [get]
func GetAllTodoLists(c *fiber.Ctx) error {
    var todos []models.Todo
    database.DB.Find(&todos)
    return c.JSON(todos)
}

// UpdateTodoList godoc
//	@Summary		Update a todo list by ID
//	@Description	Update a todo list by ID
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Todo ID"
//	@Param			todo	body		types.Payload	true	"Todo list"
//	@Success		200		{object}	models.Todo
//	@Failure		400		{object}	types.ErrorResponse
//	@Failure		403		{object}	types.ErrorResponse
//	@Failure		404		{object}	types.ErrorResponse
//	@Router			/todos/{id} [put]
func UpdateTodoList(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo models.Todo

    if err := database.DB.First(&todo, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
    }

    if todo.Deleted {
        return c.Status(403).JSON(fiber.Map{"error": "Cannot update deleted todo"})
    }

    if todo.Completed {
        return c.Status(403).JSON(fiber.Map{"error": "Cannot update completed todo"})
    }

    if err := c.BodyParser(&todo); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    database.DB.Save(&todo)
    return c.JSON(todo)
}

// DeleteTodoList godoc
//	@Summary		Soft delete a todo list by ID
//	@Description	Soft delete a todo list by ID
//	@Tags			todos
//	@Param			id	path	int	true	"Todo ID"
//	@Success		204
//	@Failure		404	{object}	types.ErrorResponse
//	@Router			/todos/{id} [delete]
func DeleteTodoList(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo models.Todo

    if err := database.DB.First(&todo, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
    }

    todo.Deleted = true
    database.DB.Save(&todo)
    return c.SendStatus(204)
}

// MarkAsCompleted godoc
//	@Summary		Mark a todo list as completed by ID
//	@Description	Mark a todo list as completed by ID
//	@Tags			todos
//	@Param			id	path		int	true	"Todo ID"
//	@Success		200	{object}	models.Todo
//	@Failure		404	{object}	types.ErrorResponse
//	@Router			/todos/{id}/completed [patch]
func MarkAsCompleted(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo models.Todo

    if err := database.DB.First(&todo, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
    }

    todo.Completed = true
    database.DB.Save(&todo)
    return c.JSON(todo)
}
