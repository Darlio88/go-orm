package main

import (
	"log"
	"net/http"

	"github.com/darlio88/go-orm/handlers"
	"github.com/darlio88/go-orm/internals"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	//instance of the database
	db := internals.DatabaseInstance()
	// close connection
	defer db.Close()

	//ping the database
	err := db.Ping()
	if err != nil {
		log.Println("err connecting to database", err)
	}
	log.Println("connected to database")

	//new instance of the server
	r := chi.NewRouter()

	//middlewares
	r.Use(middleware.Logger)

	//testing server
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.Get("/users/{id}", handlers.GetUser)
	r.Patch("/users/{id}", handlers.UpdateUser)
	r.Delete("/users/{id}", handlers.DeleteUser)
	r.Get("/users", handlers.GetAllUsers)
	http.ListenAndServe(":5000", r)

}
