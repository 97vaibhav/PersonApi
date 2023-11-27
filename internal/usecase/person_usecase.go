package usecase

import (
	"github.com/97vaibhav/PersonApi/internal/domain"
)

type PersonRepository interface {
	GetAll() []domain.Person
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
