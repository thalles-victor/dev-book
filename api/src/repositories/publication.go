package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repository Publications) Search(userID uint64) ([]models.Publication, error) {
	rows, err := repository.db.Query(`
		SELECT p.*, u.nick FROM publication p
			INNER JOIN users u on u.id = p.author_id
			INNER JOIN subscriptions s ON p.author_id = s.user_id
		WHERE u.id = ? OR s.follower_id = ?;
	`,
		userID, userID,
	)

	fmt.Println(rows.Columns())

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pubs []models.Publication

	for rows.Next() {
		var pub models.Publication

		if err = rows.Scan(
			&pub.ID,
			&pub.Title,
			&pub.Content,
			&pub.AuthorID,
			&pub.Likes,
			&pub.CreatedAt,
			&pub.AuthorNick,
		); err != nil {
			return nil, err
		}

		pubs = append(pubs, pub)
	}

	return pubs, nil
}

// Update the data of a publication
func (repository Publications) Update(pubID uint64, pub models.Publication) error {
	statement, err := repository.db.Prepare("UPDATE publication SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(pub.Title, pub.Content, pub.ID); err != nil {
		return err
	}

	return nil
}

func (repository Publications) DeleteById(pubID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM publication WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(pubID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Publications) SearchByUserID(userID uint64) ([]models.Publication, error) {
	rows, err := repository.db.Query(`
	  SELECT p.* FROM publications p
	  JOIN users u ON u.id = p.author_id
	  WHERE p.author_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Curtir add a linke in a publication
func (repository Publications) Like(pubID uint64) error {
	statement, err := repository.db.Prepare("UPDATE publication SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	statement.Exec(pubID)

	return nil
}

func (repository Publications) Unlike(pubID uint64) error {
	statement, err := repository.db.Prepare("UPDATE publication SET likes = likes - 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	statement.Exec(pubID)

	return nil
}
