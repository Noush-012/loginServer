package main

import (
	"fmt"
	"net/http"

	"github.com/Noush-012/Login-Page-Server/controls"
	"github.com/Noush-012/Login-Page-Server/db"
	"github.com/gorilla/mux"
)

var port = ":8080"

func main() {
	db.DataBase["admin@gmail.com"] = db.UserDetails{
		Name:  "Noush.",
		Email: "admin@gmail.com",
		Pass:  "123123",
	}

	r := mux.NewRouter()
	r.HandleFunc("/", controls.LoginPage).Methods("GET")
	r.HandleFunc("/", controls.LoginSubmit).Methods("POST")
	r.HandleFunc("/home", controls.HomePage)
	r.HandleFunc("/register", controls.RegisterPage).Methods("GET")
	r.HandleFunc("/register", controls.RegisterSubmit).Methods("POST")
	r.HandleFunc("/logout", controls.Logout)

	// if no handler found then use these handler
	r.NotFoundHandler = http.HandlerFunc(controls.ErrorHandleFunc)
	fmt.Println("Server running at ", port)

	http.ListenAndServe(port, r)
}
