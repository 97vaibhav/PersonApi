package repository

import (
	"errors"
	"strings"
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

func (r *PersonRepository) CreatePerson(person domain.Person) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, existingPerson := range r.people {
		if existingPerson.Email == person.Email {
			return errors.New("email already exist so cant create person")
		}
	}
	r.people = append(r.people, person)
	return nil
}

func (r *PersonRepository) UpdatePersonDetails(id string, updatePerson domain.Person) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, p := range r.people {
		if p.ID == id {
			r.people[i] = updatePerson
		}
		return nil
	}
	return errors.New("person not found ")
}

func (r *PersonRepository) DeletePerson(id string) (domain.Person, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, p := range r.people {
		if p.ID == id {
			// Remove the person from the slice
			deletedPerson := r.people[i]
			r.people = append(r.people[:i], r.people[i+1:]...)
			return deletedPerson, nil
		}
	}

	return domain.Person{}, errors.New("person not found")
}

func (r *PersonRepository) GetByName(name string) []domain.Person {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var matchingPeople []domain.Person
	for _, p := range r.people {
		if strings.Contains(strings.ToLower(p.FirstName), strings.ToLower(name)) ||
			strings.Contains(strings.ToLower(p.LastName), strings.ToLower(name)) {
			matchingPeople = append(matchingPeople, p)
		}
	}

	return matchingPeople
}
