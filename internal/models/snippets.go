package models

import (
	"database/sql"
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
func (m *SnippetModel) Insert(title string, content string, expires string) (int, error) {
	return 0, nil
}

// Return specific snippet based on ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Return the 10 most recently created snippet
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
