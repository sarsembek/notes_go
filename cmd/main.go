package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sarsembek/notes_go/internal/api"
	"github.com/sarsembek/notes_go/internal/store"
)

func main() {
	e := echo.New()

	dsn := "host=localhost user=postgres password=postgres dbname=notes sslmode=disable"

	store := store.NewStore(dsn)

	// Инициализация маршрутов
	api.SetupRoutes(e, store)

	e.Logger.Fatal(e.Start(":8080"))
}
