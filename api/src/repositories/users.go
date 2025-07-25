package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
)

type Users struct {
	db *sql.DB
}

// Create a new user repository
func NewRepositoryUser(db *sql.DB) *Users {
	return &Users{db}
}

// Create and insert a new use in the database
func (repository Users) Create(user models.User) (uint64, error) {
	statment, err := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil

}

// Search all users with filters
func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	rows, err := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FORM from users WHERE name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Get one user with user ID
func (repository Users) SearchByID(userID uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, name, nick, email, createdAt from users where id = ?", userID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update user informations
func (repository Users) Update(userId uint64, user models.User) error {
	statement, err := repository.db.Prepare(`UPDATE users SET name = ?, nick = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	fmt.Printf("user id is %d", userId)
	if _, err = statement.Exec(user.Name, user.Nick, userId); err != nil {
		return err
	}

	return nil
}

func (repository Users) DeleteUser(userId uint64) error {
	statement, err := repository.db.Prepare(`DELETE FROM users WHERE id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId); err != nil {
		return err
	}

	return nil
}

// Search user by email and return id and password with hash
func (repository Users) SearchUserByEmail(email string) (models.User, error) {
	var user models.User

	err := repository.db.QueryRow("SELECT id, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("user not found")
		}

		return models.User{}, err
	}

	return user, nil
}

// Follow allows a user to follow another
func (repository Users) Follow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare("INSERT IGNORE INTO subscriptions (user_id, follower_id) VALUES (?, ?);")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

// Unfollow user
func (repository Users) Unfollow(userId, followId uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM subscriptions WHERE user_id = ? AND follower_id = ?")
	fmt.Println(statement)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, followId); err != nil {
		return err
	}

	return nil
}

// Search all follower from user
func (repository Users) SearchFollowersFromUser(userId uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT
			u.id,
			u.name,
			u.nick,
			u.email
		FROM users u
		INNER JOIN subscriptions s
			ON u.id = s.follower_id WHERE s.user_id = ?
	`, userId)

	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Nick,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) SearchFollowingFromUser(userId uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email
		FROM users u INNER JOIN subscriptions s ON u.id = s.user_id WHERE s.follower_id = ?`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) GetPasswordByUserId(userId uint64) (string, error) {
	rows, err := repository.db.Query(`SELECT password FROM user WHERE id = ?`, userId)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	rows.Scan()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository Users) UpdatePassword(userID uint64, password string) error {
	statement, err := repository.db.Prepare(`UPDATE user SET password = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}
