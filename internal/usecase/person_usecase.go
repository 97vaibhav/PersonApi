package usecase

import (
	"github.com/97vaibhav/PersonApi/internal/domain"
	"github.com/97vaibhav/PersonApi/internal/errors"
)

type PersonRepository interface {
	GetAll() []domain.Person
	GetById(id string) (domain.Person, error)
	CreatePerson(person domain.Person) error
	UpdatePersonDetails(id string, updatedPerson domain.Person) error
	DeletePerson(id string) (domain.Person, error)
}

type PersonUsecase struct {
	personRepo PersonRepository
}

func NewPersonUsecase(personRepo PersonRepository) *PersonUsecase {
	return &PersonUsecase{
		personRepo: personRepo,
	}
}
func (uc *PersonUsecase) GetAll() []domain.Person {
	return uc.personRepo.GetAll()
}

func (uc *PersonUsecase) GetById(id string) (domain.Person, error) {
	return uc.personRepo.GetById(id)
}

func (uc *PersonUsecase) CreatePerson(person domain.Person) error {
	// Validating non-empty first name
	if person.FirstName == "" {
		return errors.ErrEmptyFirstName
	}
	// Validating if email is correct format and not empty
	if !isValidEmail(person.Email) || person.Email == "" {
		return errors.ErrInvalidEmail
	}
	// Validating birthday format
	if !isValidBirthday(person.Birthday) {
		return errors.ErrInvalidBirthday
	}

	return uc.personRepo.CreatePerson(person)
}

func (uc *PersonUsecase) UpdatePersonDetails(id string, updatedPerson domain.Person) error {
	// Validate email format and non-empty value
	if !isValidEmail(updatedPerson.Email) || updatedPerson.Email == "" {
		return errors.ErrInvalidEmail
	}

	// Validate birthday format and non-empty value
	if !isValidBirthday(updatedPerson.Birthday) || updatedPerson.Birthday == "" {
		return errors.ErrInvalidBirthday
	}
	// Validate non-empty first name
	if updatedPerson.FirstName == "" {
		return errors.ErrEmptyFirstName
	}
	// Geting the existing person
	existingPerson, err := uc.personRepo.GetById(id)
	if err != nil {
		return errors.ErrPersonNotFound
	}

	// Updating fields
	existingPerson.FirstName = updatedPerson.FirstName
	existingPerson.LastName = updatedPerson.LastName
	existingPerson.Email = updatedPerson.Email
	existingPerson.Birthday = updatedPerson.Birthday

	return uc.personRepo.UpdatePersonDetails(id, existingPerson)
}

func (uc *PersonUsecase) DeletePerson(id string) (domain.Person, error) {
	return uc.personRepo.DeletePerson(id)
}

func (uc *PersonUsecase) GetByName(name string) []domain.Person {
	var matchingPeople []domain.Person
	people := uc.personRepo.GetAll()

	// Iterate through people and check if either first or last name contains the provided name
	for _, person := range people {
		if containsName(person.FirstName, name) || containsName(person.LastName, name) {
			matchingPeople = append(matchingPeople, person)
		}
	}
	return matchingPeople
}

func (uc *PersonUsecase) GetAgeByID(id string) (int, error) {
	person, err := uc.personRepo.GetById(id)
	if err != nil {
		return 0, err
	}
	return calculateAge(person.Birthday), nil
}
