package main

import (
	"net/http"

	InfraController "auth/infra/controller"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", EmptyHandler)

	authController := InfraController.NewAuthController(mux)
	authController.RegisterRoutes()

	http.ListenAndServe(":8080", mux)
}

func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Auth Service"))
}
