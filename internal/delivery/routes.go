package delivery

import (
	"github.com/97vaibhav/PersonApi/internal/usecase"
	"github.com/gorilla/mux"
)

func RegisterPersonHandlers(router *mux.Router, personUsecase *usecase.PersonUsecase) {
	router.HandleFunc("/people", getPeopleHandler(personUsecase)).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonByID(personUsecase)).Methods("GET")
	router.HandleFunc("/people", createPerson(personUsecase)).Methods("POST")
	router.HandleFunc("/people/{id}", updatePersonDetails(personUsecase)).Methods("PUT")
	router.HandleFunc("/people/{id}", deletePerson(personUsecase)).Methods("DELETE")
	router.HandleFunc("/people", getPeopleByNameHandler(personUsecase)).Methods("GET")
	router.HandleFunc("/people/{id}/age", getPersonAgeHandler(personUsecase)).Methods("GET")
}
