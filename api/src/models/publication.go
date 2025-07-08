package models

import (
	"errors"
	"strings"
	"time"
)

// Represent a publication maked from user
type Publication struct {
	ID         uint64    `json:"id"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (pub *Publication) Prepare() error {
	if err := pub.validate(); err != nil {
		return err
	}

	pub.format()
	return nil
}

func (pub *Publication) validate() error {
	if pub.Title == "" {
		return errors.New("O título é obrigatorio")
	}

	if pub.Content == "" {
		return errors.New("O conteúdo é obrigatorio e não podeestar em branco")
	}

	return nil
}

func (pub *Publication) format() {
	pub.Title = strings.TrimSpace(pub.Title)
	pub.Content = strings.TrimSpace(pub.Content)
}
