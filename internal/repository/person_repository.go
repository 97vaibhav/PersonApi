package repository

import (
	"errors"
	"sync"

	"github.com/97vaibhav/PersonApi/internal/domain"
)

type PersonRepository struct {
	mutex  sync.RWMutex
	people []domain.Person
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{
		people: []domain.Person{
			{
				ID:        "bf552a1c-fd73-4bd0-b64a-d3f69a9ff9de",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "johndoe@example.com",
				Birthday:  "1997-01-01",
			},
			{
				ID:        "d5356358-b39f-4c6e-9690-2c965a607702",
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "janedoe@example.com",
				Birthday:  "1991-07-28",
			},
			{
				ID:        "cb2bfa60-e2ae-46ec-ad77-60cf7e8979fd",
				FirstName: "Brian",
				LastName:  "Smith",
				Email:     "briansmith@example.com",
				Birthday:  "2000-05-10",
			},
			{
				ID:        "d82fc695-5ac2-4fed-9387-a7d9c0fb0c4f",
				FirstName: "Ashley",
				LastName:  "Yu",
				Email:     "ashleyyu@example.com",
				Birthday:  "2003-12-24",
			},
		},
	}
}

func (r *PersonRepository) GetAll() []domain.Person {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.people
}

func (r *PersonRepository) GetById(id string) (domain.Person, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, p := range r.people {
		if p.ID == id {
			return p, nil
		}
	}
	return domain.Person{}, errors.New("person Not found in database")
}
