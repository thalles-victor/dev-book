package models

import (
	"api/src/security"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (user *User) Prepare(step string) error {

	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if step != "register" && step != "update" {
		log.Fatal("step must be a string or update")
	}

	if user.Name == "" {
		return errors.New("O name é obrigatorio e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("O nick é obrigatorio e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O email é obrigatorio e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O email inserido é inválido")
	}

	if step == "register" && user.Password == "" {
		return errors.New("O password é obrigatorio e não pode estar em branco")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}
