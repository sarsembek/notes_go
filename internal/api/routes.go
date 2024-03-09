package api

import (
	"github.com/labstack/echo/v4"
	"github.com/yourusername/yourprojectname/internal/store"
)

func SetupRoutes(e *echo.Echo, store *store.Store) {
	handler := NewHandler(store)

	// Маршруты для работы с пользователями
	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)

	// Маршруты для работы с заметками
	e.GET("/notes", handler.GetNotes)

	// Добавьте другие маршруты согласно вашему API
}
