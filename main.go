package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-app/database"
	_ "todo-app/docs"
	"todo-app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

//	@title			Todo API
//	@version		1.0
//	@description	This is a simple todo list API.
//	@host			localhost:3000
//	@BasePath		/
func main() {
    app := fiber.New()

    // Initialize the database connection
    database.Connect()

    // Setup the routes
    routes.Setup(app)

    // Swagger endpoint
    app.Get("/swagger/*", swagger.HandlerDefault)

    // app.Listen(":3000")
    // Start server in a goroutine so that it doesn't block
    go func() {
        if err := app.Listen(":3000"); err != nil {
            log.Panicf("Error starting server: %v", err)
        }
    }()

    // Channel to listen for interrupt signals
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    // Block until an interrupt signal is received
    <-quit

    log.Println("Shutting down server...")

    // Create a deadline to wait for
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Shut down the server with a deadline
    if err := app.ShutdownWithContext(ctx); err != nil {
        log.Fatalf("Server forced to shut down: %v", err)
    }
}
