package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/97vaibhav/PersonApi/internal/domain"
	"github.com/97vaibhav/PersonApi/internal/usecase"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RegisterPersonHandlers(router *mux.Router, personUsecase *usecase.PersonUsecase) {
	router.HandleFunc("/people", getPeopleHandler(personUsecase)).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonByID(personUsecase)).Methods("GET")
	router.HandleFunc("/people", createPerson(personUsecase)).Methods("POST")
}

func getPeopleHandler(personUsecase *usecase.PersonUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		people := personUsecase.GetAll()
		json.NewEncoder(w).Encode(people)
	}
}

func getPersonByID(personUsecase *usecase.PersonUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		personId := params["id"]

		person, err := personUsecase.GetById(personId)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "person with ID %s not found ", personId)
			return
		}
		json.NewEncoder(w).Encode(person)
	}
}

func createPerson(personUsecase *usecase.PersonUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var newPerson domain.Person

		err := json.NewDecoder(r.Body).Decode(&newPerson)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error decoding request body: %v", err)
			return
		}
		newPerson.ID = uuid.New().String()

		err = personUsecase.CreatePerson(newPerson)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error creating person: %v", err)
			return
		}
		json.NewEncoder(w).Encode(newPerson)

	}
}
