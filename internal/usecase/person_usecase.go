package usecase

import (
	"github.com/97vaibhav/PersonApi/internal/domain"
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
	return uc.personRepo.CreatePerson(person)
}

func (uc *PersonUsecase) UpdatePersonDetails(id string, updatedPerson domain.Person) error {
	return uc.personRepo.UpdatePersonDetails(id, updatedPerson)
}

func (uc *PersonUsecase) DeletePerson(id string) (domain.Person, error) {
	return uc.personRepo.DeletePerson(id)
}

func (uc *PersonUsecase) GetByName(name string) []domain.Person {
	var matchingPeople []domain.Person
	people := uc.personRepo.GetAll()

	for _, person := range people {
		if ContainsName(person.FirstName, name) || ContainsName(person.LastName, name) {
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

	return CalculateAge(person.Birthday), nil
}
