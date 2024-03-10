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
    // defer db.CloseDB()

    // Create a new router
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/pokemon", handlers.CreatePokemon).Methods("POST")
	r.HandleFunc("/pokemon", handlers.GetAllPokemon).Methods("GET")
    r.HandleFunc("/pokemon/{id}", handlers.GetPokemon).Methods("GET")
    r.HandleFunc("/pokemon/{id}", handlers.UpdatePokemon).Methods("PUT")
    r.HandleFunc("/pokemon/{id}", handlers.DeletePokemon).Methods("DELETE")

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", r))
}
