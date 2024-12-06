package models

import (
	"database/sql"
	"errors"
	"time"
)

// Snippet type with data for individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expired time.Time
}

// SnippetModel type which wraps sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert new snippet into DB
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId() // Gets the ID of newly inserted record
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Return specific snippet based on ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	s := &Snippet{}

	// Get the specific row by passing in id and use Scan to point them to location to copy to
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expired)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// Return the 10 most recently created snippet
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
