package store

import (
	"context"
	"github.com/sarsembek/notes_go/internal/model"
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

// GetNoteByID возвращает заметку по её ID
func (s *Store) GetNoteByID(ctx context.Context, id int) (*model.Note, error) {
	note := &model.Note{}

	query := `SELECT id, user_id, title, content FROM notes WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&note.ID, &note.UserID, &note.Title, &note.Content); err != nil {
		return nil, err
	}

	return note, nil
}

// UpdateNote обновляет заметку с заданным ID
func (s *Store) UpdateNote(ctx context.Context, id int, note *model.Note) error {
    query := `UPDATE notes SET title = $1, content = $2, updated_at = NOW() WHERE id = $3`
    _, err := s.db.ExecContext(ctx, query, note.Title, note.Content, id)
    return err
}

// DeleteNote удаляет заметку по её ID
func (s *Store) DeleteNote(ctx context.Context, id int) error {
	query := `DELETE FROM notes WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	return err
}
