package models

type CreatePersonRequest struct {
	Name      string   `json:"nome" validate:"required,max=100"`
	Nickname  string   `json:"apelido" validate:"required,max=32"`
	Birthdate string   `json:"nascimento" validate:"required,datetime=2000-01-01"`
	Stack     []string `json:"stack" validate:"dive,max=32"`
}
