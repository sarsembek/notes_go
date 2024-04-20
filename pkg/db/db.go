package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
    "pokemon_project/pkg/model"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/pokemon?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// CreateUser inserts a new user into the database
func CreateUser(user model.User) error {
	_, err := DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	return err
}

// GetUserByUsername retrieves a user by username from the database
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// AuthenticateUser checks if the given login credentials are valid
func AuthenticateUser(username, password string) (bool, *model.User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return false, nil, err
	}

	isAuthenticated, err := user.Authenticate(password)
	if err != nil {
		return false, nil, err
	}

	return isAuthenticated, user, nil
}

func CreatePokemon(p model.Pokemon) error {
	_, err := DB.Exec("INSERT INTO Pokemon (name, species, type1, type2, height, weight, base_experience, capture_rate, hp, attack, defense, special_attack, special_defense, speed) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		p.Name,
		p.Species,
		p.Type1,
		p.Type2,
		p.Height,
		p.Weight,
		p.BaseExp,
		p.CaptureRate,
		p.HP,
		p.Attack,
		p.Defense,
		p.SpAttack,
		p.SpDefense,
		p.Speed,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPokemon() ([]model.Pokemon, error) {
    var pokemons []model.Pokemon
    rows, err := DB.Query("SELECT id, name, species, type1, type2, height, weight, base_experience, capture_rate, hp, attack, defense, special_attack, special_defense, speed FROM Pokemon")
    if err != nil {
        return pokemons, err
    }
    defer rows.Close()

    for rows.Next() {
        var p model.Pokemon
        if err := rows.Scan(
            &p.ID,
            &p.Name,
            &p.Species,
            &p.Type1,
            &p.Type2,
            &p.Height,
            &p.Weight,
            &p.BaseExp,
            &p.CaptureRate,
            &p.HP,
            &p.Attack,
            &p.Defense,
            &p.SpAttack,
            &p.SpDefense,
            &p.Speed,
        ); err != nil {
            return pokemons, err
        }
        pokemons = append(pokemons, p)
    }
    if err := rows.Err(); err != nil {
        return pokemons, err
    }
    return pokemons, nil
}


func GetPokemonByID(id int) (model.Pokemon, error) {
	var pokemon model.Pokemon
	err := DB.QueryRow("SELECT id, name, species, type1, type2, height, weight, base_experience, capture_rate, hp, attack, defense, special_attack, special_defense, speed FROM Pokemon WHERE id = $1", id).Scan(
		&pokemon.ID,
		&pokemon.Name,
		&pokemon.Species,
		&pokemon.Type1,
		&pokemon.Type2,
		&pokemon.Height,
		&pokemon.Weight,
		&pokemon.BaseExp,
		&pokemon.CaptureRate,
		&pokemon.HP,
		&pokemon.Attack,
		&pokemon.Defense,
		&pokemon.SpAttack,
		&pokemon.SpDefense,
		&pokemon.Speed,
	)
	if err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func GetPokemonWithPagination(page, limit int, sortBy, filterBy string) ([]model.Pokemon, error) {
    var pokemons []model.Pokemon

    query := "SELECT id, name, species, type1, type2, height, weight, base_experience, capture_rate, hp, attack, defense, special_attack, special_defense, speed FROM Pokemon"

    if filterBy != "" {
        query += " WHERE name LIKE '%" + filterBy + "%'"
    }

    if sortBy != "" {
        query += " ORDER BY " + sortBy
    }

    query += " LIMIT $1 OFFSET $2"
    offset := (page - 1) * limit

    rows, err := DB.Query(query, limit, offset)
    if err != nil {
        return pokemons, err
    }
    defer rows.Close()

    for rows.Next() {
        var p model.Pokemon
        if err := rows.Scan(
            &p.ID,
            &p.Name,
            &p.Species,
            &p.Type1,
            &p.Type2,
            &p.Height,
            &p.Weight,
            &p.BaseExp,
            &p.CaptureRate,
            &p.HP,
            &p.Attack,
            &p.Defense,
            &p.SpAttack,
            &p.SpDefense,
            &p.Speed,
        ); err != nil {
            return pokemons, err
        }
        pokemons = append(pokemons, p)
    }
    if err := rows.Err(); err != nil {
        return pokemons, err
    }
    return pokemons, nil
}

func UpdatePokemonByID(id int, p model.Pokemon) error {
	_, err := DB.Exec("UPDATE Pokemon SET name = $1, species = $2, type1 = $3, type2 = $4, height = $5, weight = $6, base_experience = $7, capture_rate = $8, hp = $9, attack = $10, defense = $11, special_attack = $12, special_defense = $13, speed = $14 WHERE id = $15",
		p.Name,
		p.Species,
		p.Type1,
		p.Type2,
		p.Height,
		p.Weight,
		p.BaseExp,
		p.CaptureRate,
		p.HP,
		p.Attack,
		p.Defense,
		p.SpAttack,
		p.SpDefense,
		p.Speed,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeletePokemonByID(id int) error {
	_, err := DB.Exec("DELETE FROM Pokemon WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
