package main

import (
	"fmt"
	"net/http"
	"os"
	"simple-rest-api/app"
	"simple-rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/auth/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/auth/login", controllers.Authenticate).Methods("POST")

	router.Use(app.UserAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
