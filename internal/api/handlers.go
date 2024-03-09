package api

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/yourusername/yourprojectname/internal/model"
	"github.com/yourusername/yourprojectname/internal/store"
)

// Handler структура для инъекции зависимостей
type Handler struct {
	store *store.Store
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{store: store}
}

// Login обработчик для аутентификации пользователя
func (h *Handler) Login(c echo.Context) error {
	// Реализация логики входа
	return c.JSON(http.StatusOK, "Logged in")
}

// Register обработчик для регистрации пользователя
func (h *Handler) Register(c echo.Context) error {
	// Реализация логики регистрации
	return c.JSON(http.StatusOK, "Registered")
}

// GetNotes обработчик для получения списка заметок
func (h *Handler) GetNotes(c echo.Context) error {
	notes, err := h.store.GetNotes(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, notes)
}
