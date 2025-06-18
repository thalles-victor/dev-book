package models

import (
	"errors"
	"strings"
	"time"
)

// User represent a user in social network
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare call validate and  format
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("O nome é obrigatorio e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("O nome é obrigatorio e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O nome é obrigatorio e não pode estar em branco")
	}

	if user.Password == "" {
		return errors.New("O nome é obrigatorio e não pode estar em branco")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
