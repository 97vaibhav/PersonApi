package usecase

import (
	"github.com/97vaibhav/PersonApi/internal/domain"
)

type PersonRepository interface {
	GetAll() []domain.Person
	GetById(id string) (domain.Person, error)
	CreatePerson(person domain.Person) error
	UpdatePersonDetails(id string, updatedPerson domain.Person) error
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
