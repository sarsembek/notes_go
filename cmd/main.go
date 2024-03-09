package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sarsembek/notes_go/internal/api"
)

func main() {
	e := echo.New()

	// Инициализация маршрутов
	api.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
