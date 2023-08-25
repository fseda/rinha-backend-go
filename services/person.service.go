package services

import (
	"database/sql"
	"strings"
	"time"

	"github.com/fseda/rinha-backend-go/models"
	dto "github.com/fseda/rinha-backend-go/models/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"

	"github.com/google/uuid"
)

type PersonService struct {
	db *sql.DB
}

func NewPersonService(db *sql.DB) *PersonService {
	return &PersonService{db}
}

func (ps *PersonService) InsertPerson(name string, nickname string, birthdate string, stack pq.StringArray) error {
	id := uuid.New()
	_, err := ps.db.Exec("INSERT INTO people (id, name, nickname, birthdate, stack) VALUES ($1, $2, $3, $4, $5)",
		id, name, nickname, birthdate, stack)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PersonService) GetPersonById(id string) (models.Person, error) {
	var person models.Person

	err := ps.db.QueryRow(`SELECT id, name, nickname, birthdate, stack FROM people WHERE id = ($1)`, id).Scan(
		&person.ID,
		&person.Name,
		&person.Nickname,
		&person.Birthdate,
		&person.Stack,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Person{}, models.ErrPersonNotFound
		} else {
			return models.Person{}, err
		}
	}
	return person, nil
}

func (ps *PersonService) SearchBy(term string) ([]models.Person, error) {
	var people []models.Person

	cleanTerm := strings.ToLower(term)
	query := `
		SELECT * 
		FROM people 
		WHERE name ILIKE '%' || $1 || '%' 
		OR nickname ILIKE '%' || $1 || '%' 
		OR EXISTS (SELECT 1 FROM unnest(stack) AS s WHERE s ILIKE '%' || $1 || '%')
		LIMIT 50
	`

	rows, err := ps.db.Query(query, cleanTerm)
	if err != nil {
		return people, err
	}
	defer rows.Close()

	for rows.Next() {
		var person models.Person
		err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Nickname,
			&person.Birthdate,
			&person.Stack,
		)
		if err != nil {
			return people, err
		}
		people = append(people, person)
	}

	return people, nil
}

func (ps *PersonService) CountPeople() (int, error) {
	var count int

	err := ps.db.QueryRow(`SELECT COUNT(*) FROM people`).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ps *PersonService) getPersonByNickname(nickname string) (string, error) {
	var nicknameInDB string
	err := ps.db.QueryRow(`SELECT nickname FROM people WHERE nickname = ($1)`, nickname).Scan(&nicknameInDB)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // No record with with given nickname
		} else {
			return "", err // Some error occurred in the database
		}
	}

	return nicknameInDB, nil // Record exists
}

func (ps *PersonService) NicknameTaken(nickname string) bool {
	nicknameInDB, err := ps.getPersonByNickname(nickname)
	return nicknameInDB != "" && err == nil
}

func (ps *PersonService) ValidateBody(body dto.CreatePersonRequest) error {
	var err error

	if ps.NicknameTaken(body.Nickname) {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Nickname already taken")
	}

	if !isString(body.Name) || !isString(body.Nickname) {
		return fiber.ErrBadRequest
	}

	date_layout := "2000-01-01"
	_, err = time.Parse(date_layout, body.Birthdate)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid birthdate", err.Error())
	}

	for _, tech := range body.Stack {
		if !isString(tech) {
			return fiber.ErrBadRequest
		}

		if len(tech) > 32 {
			return fiber.ErrBadRequest
		}
	}
	return nil
}


func isString(variable interface{}) bool {
	switch variable.(type) {
	case string:
		return true
	default:
		return false
	}
}
