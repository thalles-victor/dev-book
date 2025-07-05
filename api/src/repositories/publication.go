package repositories

import (
	"api/src/models"
	"database/sql"
)

// Represent publication repository
type Publications struct {
	db *sql.DB
}

// Create a new repository of publications
func NewRepositoryPublication(db *sql.DB) *Publications {
	return &Publications{db}
}

// Create a new publication in database
func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO publication (title, content, author_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}
