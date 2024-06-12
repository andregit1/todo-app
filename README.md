# Overview

This is a simple to-do list app that allows users to create, read, update, and delete tasks (CRUD).

## Features

- Create todo list entries
- Update todo list entries
- Delete todo list entries
- Get a todo list entry by ID
- Get all todo list entries
- Mark a todo list entry as completed

## Technologies

- Go
- Fiber
- GORM
- MySQL
- Swaggo

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/username/todo-app.git
   cd todo-app
   ```

2. Configure the database:
   Rename `.env.example` to `.env` and configure your database settings.

3. Build and start the application using Docker Compose:

   ```sh
   docker compose build
   docker compose up
   ```

4. Access Swagger documentation:
   Open your browser and navigate to `http://localhost:3000/swagger/index.html` to view the API documentation.

## Endpoints

### Todo List

- `POST /todos` - Create a new todo list
- `GET /todos/{id}` - Get a todo list by ID
- `GET /todos` - Get all todo lists
- `PUT /todos/{id}` - Update a todo list by ID
- `DELETE /todos/{id}` - Delete a todo list by ID
- `PATCH /todos/{id}/completed` - Mark a todo list as completed

## Project Structure

- `main.go` - The application's entry point
- `models/` - Data structures for todo lists
- `handlers/` - Handlers for the endpoints
- `database/` - Database configuration
- `types/` - Data structures for error responses and others
- `docs/` - Swagger documentation

### Example Database Schema

```sql
CREATE DATABASE IF NOT EXISTS todoapp;

USE todoapp;

CREATE TABLE IF NOT EXISTS todos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    title VARCHAR(255) NOT NULL UNIQUE,
    tasks TEXT,
    completed BOOLEAN DEFAULT FALSE,
    deleted BOOLEAN DEFAULT FALSE
);
```

### Project Directory Structure

```
todo-app/
│
├── database/
│   └── database.go
│
├── docs/
│   └── swagger files
│
├── handlers/
│   └── todoHandler.go
│
├── models/
│   └── todo.go
|
├── routes/
│   └── todoRoutes.go
│
├── types/
│   └── error_response.go
│   └── payload.go
│
├── .env
├── .env.example
├── .air.toml
├── .gitignore
├── go.mod
├── go.sum
├── main.go
├── README.md
└── schema.sql
```
