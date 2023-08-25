package handlers

import (
	"errors"

	"github.com/fseda/rinha-backend-go/database"
	"github.com/fseda/rinha-backend-go/models"
	dto "github.com/fseda/rinha-backend-go/models/dto"
	"github.com/fseda/rinha-backend-go/services"

	"github.com/gofiber/fiber/v2"
)

var InvalidDtoErr = errors.New("invalid dto")

func HandleCreatePerson(c *fiber.Ctx) error {
	var err error
	var body dto.CreatePersonRequest
	ps := services.NewPersonService(database.Conn)

	if err = c.BodyParser(&body); err != nil {
		return err
	}

	err = ps.ValidateBody(body)
	if err != nil {
		return err
	}

	id, err := ps.InsertPerson(body.Name, body.Nickname, body.Birthdate, body.Stack)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create person", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func HandleGetPersonById(c *fiber.Ctx) error {
	var err error
	var person models.Person
	id := c.Params("id")
	ps := services.NewPersonService(database.Conn)

	person, err = ps.GetPersonById(id)
	if err != nil {
		if err != models.ErrPersonNotFound {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return fiber.NewError(fiber.StatusNotFound, "Person not found")
	}

	return c.Status(fiber.StatusOK).JSON(person)
}

func HandleSearchPeople(c *fiber.Ctx) error {
	var err error
	var people []models.Person
	term := c.Queries()["t"]

	ps := services.NewPersonService(database.Conn)

	people, err = ps.SearchBy(term)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if len(people) == 0 {
		return c.Status(fiber.StatusOK).JSON([]interface{}{})
	}

	return c.Status(fiber.StatusOK).JSON(people)
}

func HandleCountPeople(c *fiber.Ctx) error {
	ps := services.NewPersonService(database.Conn)

	count, err := ps.CountPeople()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not count people")
	}

	return c.Status(fiber.StatusOK).JSON(count)
}
