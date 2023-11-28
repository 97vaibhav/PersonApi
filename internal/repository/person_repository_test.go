package repository

import (
	"testing"

	"github.com/97vaibhav/PersonApi/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := NewPersonRepository()

	people := repo.GetAll()

	assert.Equal(t, 4, len(people), "Expected number of people not returned corrcectly")
}

func TestGetById(t *testing.T) {
	repo := NewPersonRepository()
	id := "bf552a1c-fd73-4bd0-b64a-d3f69a9ff9de"

	person, err := repo.GetById(id)

	assert.NoError(t, err, "Unexpected error returned ")
	assert.Equal(t, id, person.ID, "Incorrect person ID check id once agin")
}

func TestGetByIdNotFound(t *testing.T) {
	repo := NewPersonRepository()
	id := "nonexistent-id"

	person, err := repo.GetById(id)

	assert.Error(t, err, "Expected error for nonexistent ID")
	assert.Equal(t, domain.Person{}, person, "Expected empty person for nonexistent ID")
}

func TestCreatePerson(t *testing.T) {
	repo := NewPersonRepository()
	newPerson := domain.Person{
		ID:        "12345678",
		FirstName: "New",
		LastName:  "Person",
		Email:     "newperson@example.com",
		Birthday:  "1998-10-03",
	}

	err := repo.CreatePerson(newPerson)

	assert.NoError(t, err, "Unexpected error while creating person")

	people := repo.GetAll()
	assert.Equal(t, 5, len(people), "Expected number of people not updated after creation of new person")
}

func TestCreatePersonDuplicateEmail(t *testing.T) {
	repo := NewPersonRepository()
	existingPerson := repo.GetAll()[0]

	duplicatePerson := domain.Person{
		ID:        "12345678",
		FirstName: "Duplicate",
		LastName:  "Person",
		Email:     existingPerson.Email,
		Birthday:  "1997-10-01",
	}

	err := repo.CreatePerson(duplicatePerson)

	assert.Error(t, err, "Expected error for duplicate email")
}

func TestUpdatePersonDetails(t *testing.T) {
	repo := NewPersonRepository()
	id := "d82fc695-5ac2-4fed-9387-a7d9c0fb0c4f"
	updatePerson := domain.Person{
		ID:        id,
		FirstName: "UpdatedFirstName",
		LastName:  "UpdatedLastname",
		Email:     "updatedperson@example.com",
		Birthday:  "1998-09-01",
	}

	err := repo.UpdatePersonDetails(id, updatePerson)

	assert.NoError(t, err, "Unexpected error while updating person details")

	updatedPerson, _ := repo.GetById(id)
	assert.Equal(t, updatePerson, updatedPerson, "Person details not updated correctly")
}

func TestUpdatePersonDetailsNotFound(t *testing.T) {
	repo := NewPersonRepository()
	id := "nonexistent-id"
	updatePerson := domain.Person{
		ID:        id,
		FirstName: "UpdatedFirst",
		LastName:  "UpdatedLastName",
		Email:     "updatedperson@example.com",
		Birthday:  "1998-09-01",
	}

	err := repo.UpdatePersonDetails(id, updatePerson)

	assert.Error(t, err, "Expected error for updating nonexistent person details")
}

func TestDeletePerson(t *testing.T) {
	repo := NewPersonRepository()
	id := "bf552a1c-fd73-4bd0-b64a-d3f69a9ff9de"

	deletedPerson, err := repo.DeletePerson(id)

	assert.NoError(t, err, "Unexpected error while deleting person")

	people := repo.GetAll()
	assert.Equal(t, 3, len(people), "Expected number of people not updated after deletion")
	assert.Equal(t, "bf552a1c-fd73-4bd0-b64a-d3f69a9ff9de", deletedPerson.ID, "Incorrect ID for deleted person")
}

func TestDeletePersonNotFound(t *testing.T) {
	repo := NewPersonRepository()
	id := "nonexistent-id"

	deletedPerson, err := repo.DeletePerson(id)

	assert.Error(t, err, "Expected error for deleting nonexistent person")
	assert.Equal(t, domain.Person{}, deletedPerson, "Expected empty person for deleting nonexistent person")
}
