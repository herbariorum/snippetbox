package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippedModel struct {
	DB *sql.DB
}

func (m *SnippedModel) Insert(title string, content string, expires int) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES($1, $2, (now() at time zone 'utc'), (now() at time zone 'utc' + interval '1 day' * $3))  RETURNING ID;`

	// _, err := m.DB.Query(stmt, title, content, expires)
	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippedModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippedModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
