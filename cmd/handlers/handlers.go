package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"pokemon_project/pkg/model"
	"pokemon_project/pkg/db"
)

func CreatePokemon(w http.ResponseWriter, r *http.Request) {
    var p model.Pokemon
    _ = json.NewDecoder(r.Body).Decode(&p)
    err := db.CreatePokemon(p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func GetAllPokemon(w http.ResponseWriter, r *http.Request) {
    pokemons, err := db.GetAllPokemon()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(pokemons)
}


func GetPokemon(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    pokemon, err := db.GetPokemonByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(pokemon)
}

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var p model.Pokemon
    _ = json.NewDecoder(r.Body).Decode(&p)
    err := db.UpdatePokemonByID(id, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func DeletePokemon(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    err := db.DeletePokemonByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}