

---

# Go ORM SIMPLE Application

This is a simple Go application designed for testing and showcasing basic skills in handling HTTP requests and working with a PostgreSQL database using the Go `database/sql` package. The application provides basic CRUD operations for managing user records in a PostgreSQL database.

## Features

- **Get All Users:** Retrieve a list of all users from the database.
- **Get User by ID:** Retrieve details of a specific user based on their ID.
- **Update User:** Update the details of a user using their ID.
- **Delete User:** Delete a user from the database based on their ID.

## Prerequisites

Before running the application, make sure you have the following installed:

- Go (1.16 or later)
- PostgreSQL
- `go-chi` router: `go get -u github.com/go-chi/chi/v5`

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/darlio/go-orm
   cd go-orm
   ```

2. Create a PostgreSQL database and set the connection details in a `.env` file:

   ```dotenv
   DB_URL=your_postgres_connection_url
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

## Usage

- Access the home route: `http://localhost:5000/`
- Get all users: `http://localhost:5000/users`
- Get a specific user: `http://localhost:5000/users/{id}`
- Update a user: Send a PATCH request to `http://localhost:5000/users/{id}` with the updated user data.
- Delete a user: Send a DELETE request to `http://localhost:5000/users/{id}`.

## Dependencies

- [chi](https://github.com/go-chi/chi) - A lightweight, idiomatic and composable router for building Go HTTP services.

