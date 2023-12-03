package usecase

import (
	"testing"

	"github.com/97vaibhav/PersonApi/internal/domain"
	"github.com/97vaibhav/PersonApi/internal/errors"
	"github.com/stretchr/testify/assert"
)

type MockPersonRepository struct {
	people []domain.Person
}

func (m *MockPersonRepository) GetAll() []domain.Person {
	return m.people
}

func (m *MockPersonRepository) GetById(id string) (domain.Person, error) {
	for _, person := range m.people {
		if person.ID == id {
			return person, nil
		}
	}
	return domain.Person{}, errors.ErrPersonNotFound
}

func (m *MockPersonRepository) CreatePerson(person domain.Person) error {
	for _, existingPerson := range m.people {
		if existingPerson.Email == person.Email {
			return errors.ErrEmailAlreadyExists
		}
	}
	m.people = append(m.people, person)
	return nil
}

func (m *MockPersonRepository) UpdatePersonDetails(id string, updatedPerson domain.Person) error {
	for i, person := range m.people {
		if person.ID == id {
			m.people[i] = updatedPerson
			return nil
		}
	}
	return errors.ErrPersonNotFound
}

func (m *MockPersonRepository) DeletePerson(id string) (domain.Person, error) {
	for i, person := range m.people {
		if person.ID == id {
			deletedPerson := m.people[i]
			m.people = append(m.people[:i], m.people[i+1:]...)
			return deletedPerson, nil
		}
	}
	return domain.Person{}, errors.ErrPersonNotFound
}

func TestPersonUsecase(t *testing.T) {
	mockRepo := &MockPersonRepository{
		people: []domain.Person{
			{ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Birthday: "1997-05-05"},
			{ID: "2", FirstName: "Jane", LastName: "Doe", Email: "jane@example.com", Birthday: "1998-05-15"},
			{ID: "3", FirstName: "Jay", LastName: "gupta", Email: "jay@example.com", Birthday: "1996-05-15"},
		},
	}
	usecase := NewPersonUsecase(mockRepo)

	t.Run("GetAll", func(t *testing.T) {
		people := usecase.GetAll()
		assert.Equal(t, 3, len(people), "Expected number of people not returned")
	})

	t.Run("GetById", func(t *testing.T) {
		person, err := usecase.GetById("1")
		assert.NoError(t, err, "Unexpected error")
		assert.Equal(t, "John", person.FirstName, "Incorrect person returned")
	})

	t.Run("GetByIdNotFound", func(t *testing.T) {
		person, err := usecase.GetById("nonexistent-id")
		assert.Error(t, err, "Expected error for nonexistent ID")
		assert.Equal(t, domain.Person{}, person, "Expected empty person for nonexistent ID")
	})

	t.Run("CreatePerson", func(t *testing.T) {
		newPerson := domain.Person{
			ID:        "4",
			FirstName: "NewFirst",
			LastName:  "PersonLast",
			Email:     "newperson@example.com",
			Birthday:  "2000-01-01",
		}
		err := usecase.CreatePerson(newPerson)
		assert.NoError(t, err, "Unexpected error while creating person")
		people := usecase.GetAll()
		assert.Equal(t, 4, len(people), "Expected number of people not updated after creation")
	})

	t.Run("CreatePersonDuplicateEmail", func(t *testing.T) {
		existingPerson := mockRepo.people[0]
		duplicatePerson := domain.Person{
			ID:        "4",
			FirstName: "Duplicate",
			LastName:  "Person",
			Email:     existingPerson.Email,
			Birthday:  "2000-01-01",
		}
		err := usecase.CreatePerson(duplicatePerson)
		assert.Error(t, err, "Expected error for duplicate email")
	})

	t.Run("Create Person failed due to invalid Email", func(t *testing.T) {
		newPerson := domain.Person{
			ID:        "4",
			FirstName: "NewFirst",
			LastName:  "PersonLast",
			Email:     "newpersonexample.com",
			Birthday:  "2000-01-01",
		}
		err := usecase.CreatePerson(newPerson)
		assert.Error(t, err, "error because Email is not valid")
	})
	t.Run("Create Person failed due to invalid birthday format", func(t *testing.T) {
		newPerson := domain.Person{
			ID:        "4",
			FirstName: "NewFirst",
			LastName:  "PersonLast",
			Email:     "newperson@example.com",
			Birthday:  "01-01-1998",
		}
		err := usecase.CreatePerson(newPerson)
		assert.Error(t, err, "error because birthday is not in correct format")
	})
	t.Run("UpdatePersonDetails", func(t *testing.T) {
		id := "1"
		updatePerson := domain.Person{
			ID:        id,
			FirstName: "Updated",
			LastName:  "Person",
			Email:     "updatedperson@example.com",
			Birthday:  "1995-01-01",
		}
		err := usecase.UpdatePersonDetails(id, updatePerson)
		assert.NoError(t, err, "Unexpected error while updating person details")
		updatedPerson, _ := usecase.GetById(id)
		assert.Equal(t, updatePerson, updatedPerson, "Person details not updated correctly")
	})

	t.Run("UpdatePersonDetailsForInvalidEmailID", func(t *testing.T) {
		id := "1"
		updatePerson := domain.Person{
			ID:        id,
			FirstName: "Updated",
			LastName:  "Person",
			Email:     "invalidemail",
			Birthday:  "1995-01-01",
		}
		err := usecase.UpdatePersonDetails(id, updatePerson)
		assert.Error(t, err, "Expected error because Email is not Valid")
	})
	t.Run("UpdatePersonDetails Failed Because Of wrong Birthday Format", func(t *testing.T) {
		id := "1"
		updatePerson := domain.Person{
			ID:        id,
			FirstName: "Updated",
			LastName:  "Person",
			Email:     "goodemail@gamil.com",
			Birthday:  "01-01-1997",
		}
		err := usecase.UpdatePersonDetails(id, updatePerson)
		assert.Error(t, err, "Expected error because Birthday is not in corrrect format")
	})
	t.Run("UpdatePersonDetailsNotFound", func(t *testing.T) {
		id := "nonexistent-id"
		updatePerson := domain.Person{
			ID:        id,
			FirstName: "Updated",
			LastName:  "Person",
			Email:     "updatedperson@example.com",
			Birthday:  "1995-01-01",
		}
		err := usecase.UpdatePersonDetails(id, updatePerson)
		assert.Error(t, err, "Expected error for updating nonexistent person details")
	})

	t.Run("DeletePerson", func(t *testing.T) {
		id := "1"
		deletedPerson, err := usecase.DeletePerson(id)
		assert.NoError(t, err, "Unexpected error while deleting person")
		people := usecase.GetAll()
		assert.Equal(t, 3, len(people), "Expected number of people not updated after deletion")
		assert.Equal(t, "1", deletedPerson.ID, "Incorrect ID for deleted person")
	})

	t.Run("DeletePersonNotFound", func(t *testing.T) {
		id := "nonexistent-id"
		deletedPerson, err := usecase.DeletePerson(id)
		assert.Error(t, err, "Expected error for deleting nonexistent person")
		assert.Equal(t, domain.Person{}, deletedPerson, "Expected empty person for deleting nonexistent person")
	})

	t.Run("GetByName", func(t *testing.T) {
		matchingPeople := usecase.GetByName("Ja")
		assert.Equal(t, 2, len(matchingPeople), "Expected number of matching people not returned")
	})

	t.Run("GetByName", func(t *testing.T) {
		matchingPeople := usecase.GetByName("x")
		assert.Equal(t, 0, len(matchingPeople), "Expected number of matching people not returned")
	})

	t.Run("GetByNameForLastName", func(t *testing.T) {
		matchingPeople := usecase.GetByName("gup")
		assert.Equal(t, 1, len(matchingPeople), "Expected number of matching people not returned")
	})

	t.Run("GetAgeByID", func(t *testing.T) {
		id := "2"
		age, err := usecase.GetAgeByID(id)
		assert.NoError(t, err, "Unexpected error while calculating age")
		assert.Equal(t, calculateAge("1998-05-15"), age, "Incorrect age calculated")
	})

	t.Run("GetAgeByIDNotFound", func(t *testing.T) {
		id := "nonexistent-id"
		age, err := usecase.GetAgeByID(id)
		assert.Error(t, err, "Expected error for calculating age of nonexistent person")
		assert.Equal(t, 0, age, "Expected age to be 0 for nonexistent person")
	})
}
