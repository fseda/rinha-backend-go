package handlers

import (
	"errors"

	"github.com/fseda/rinha-backend-go/services"
	"github.com/fseda/rinha-backend-go/database"

	"github.com/gofiber/fiber/v2"
)

var InvalidDtoErr = errors.New("invalid dto")

type CreatePersonRequest struct {
	Name      string   `json:"nome" validate:"required,max=100"`
	Nickname  string   `json:"apelido" validate:"required,max=32"`
	Birthdate string   `json:"nascimento" validate:"required,datetime=2000-01-01"`
	Stack     []string `json:"stack" validate:"dive,max=32"`
}

func isString(variable interface{}) bool {
    switch variable.(type) {
    case string:
        return true
    default:
        return false
    }
}

func HandleCreatePerson(c *fiber.Ctx) error {
	var err error
	var body CreatePersonRequest
	if err = c.BodyParser(&body); err != nil {
		return err
	}
	
	ps := services.NewPersonService(database.Conn)

	if !isString(body.Name) || !isString(body.Nickname) {
		return fiber.ErrBadRequest
	}

	if ps.NicknameTaken(body.Nickname) {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Nickname already taken")
	}

	for _, tech := range body.Stack {
		if !isString(tech) {
			return fiber.ErrBadRequest
		}

		if len(tech) > 32 {
			return fiber.ErrBadRequest
		}
	}

	err = ps.InsertPerson(body.Name, body.Nickname, body.Birthdate, body.Stack)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to insert person")
	}

	return nil
}
