package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "pokemon_project/cmd/handlers"
    "pokemon_project/pkg/db"
)

func main() {
    // Initialize the database
    db.InitDB()

    // Create a new router
    r := mux.NewRouter()
    api := r.PathPrefix("/api").Subrouter()
    api.Use(handlers.Authenticate)

    // Define routes
    api.HandleFunc("/pokemon", handlers.CreatePokemon).Methods("POST")
	api.HandleFunc("/pokemon", handlers.GetAllPokemon).Methods("GET")
    api.HandleFunc("/pokemon/{id}", handlers.GetPokemon).Methods("GET")
    api.HandleFunc("/pokemon/{id}", handlers.UpdatePokemon).Methods("PUT")
    api.HandleFunc("/pokemon/{id}", handlers.DeletePokemon).Methods("DELETE")

    r.HandleFunc("/user/register", handlers.RegisterUser).Methods("POST")
    r.HandleFunc("/user/login", handlers.LoginUser).Methods("POST")

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", r))
}
