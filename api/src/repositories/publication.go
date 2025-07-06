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

func (repository Publications) GetById(pubID uint64) (models.Publication, error) {
	row, err := repository.db.Query(`
		SELECT p.*, u.nick FROM
		publication p INNER JOIN users u
		ON u.id = p.author_id WHERE p.id = ?
	`, pubID)
	defer row.Close()

	if err != nil {
		return models.Publication{}, err
	}

	var pub models.Publication
	if row.Next() {
		if err = row.Scan(
			&pub.ID,
			&pub.Title,
			&pub.Content,
			&pub.AuthorID,
			&pub.Likes,
			&pub.CreatedAt,
			&pub.AuthorNick,
		); err != nil {
			return models.Publication{}, err
		}
	}
	return pub, nil
}
