package store

import (
	"context"
	"github.com/yourusername/yourprojectname/internal/model"
)

// CreateNote добавляет новую заметку в базу данных
func (s *Store) CreateNote(ctx context.Context, note *model.Note) error {
	query := `INSERT INTO notes (user_id, title, content) VALUES ($1, $2, $3) RETURNING id`
	return s.db.QueryRowContext(ctx, query, note.UserID, note.Title, note.Content).Scan(&note.ID)
}

// GetNotes возвращает список всех заметок
func (s *Store) GetNotes(ctx context.Context) ([]model.Note, error) {
	notes := []model.Note{}

	query := `SELECT id, user_id, title, content FROM notes`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}
