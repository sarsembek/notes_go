package api

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/sarsembek/notes_go/internal/store"
	"github.com/sarsembek/notes_go/internal/model"
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

// CreateNote обработчик для создания новой заметки
func (h *Handler) CreateNote(c echo.Context) error {
	var note model.Note
	if err := c.Bind(&note); err != nil {
		return err
	}

	err := h.store.CreateNote(c.Request().Context(), &note)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, note)
}
// GetNote обработчик для получения заметки по ID
func (h *Handler) GetNote(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	note, err := h.store.GetNoteByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, note)
}

func (h *Handler) UpdateNote(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id")) // Получаем ID заметки из URL
    var note model.Note
    if err := c.Bind(&note); err != nil { // Привязываем тело запроса к структуре заметки
        return err
    }

    // Вызываем метод UpdateNote из store, передавая контекст, ID и заметку
    err := h.store.UpdateNote(c.Request().Context(), id, &note)
    if err != nil {
        // Обработка ошибки, например, если заметка не найдена
        return c.JSON(http.StatusNotFound, map[string]string{"message": "Note not found"})
    }

    // Возвращаем обновленную заметку с кодом 200 OK
    return c.JSON(http.StatusOK, note)
}


// DeleteNote обработчик для удаления заметки по ID
func (h *Handler) DeleteNote(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.store.DeleteNote(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}