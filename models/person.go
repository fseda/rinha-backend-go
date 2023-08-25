package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

var (
	ErrNicknameTaken  = errors.New("nickname taken")
	ErrPersonNotFound = errors.New("person not found")
)

type Person struct {
	ID        uuid.UUID      `json:"id"`
	Nickname  string         `json:"apelido"`
	Name      string         `json:"nome"`
	Birthdate string         `json:"nascimento"`
	Stack     pq.StringArray `json:"stack"`
}
