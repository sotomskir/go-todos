package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sotomskir/go-todos/app"
	"github.com/sotomskir/go-todos/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/status", controllers.Status).Methods("GET")
	router.HandleFunc("/api/register", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos", controllers.GetTodosFor).Methods("GET")
	router.HandleFunc("/api/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}
