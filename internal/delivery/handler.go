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

func updatePersonDetails(personUsecase *usecase.PersonUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		personID := params["id"]
		var updatedPerson domain.Person
		err := json.NewDecoder(r.Body).Decode(&updatedPerson)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error decoding request body: %v", err)
			return
		}
		err = personUsecase.UpdatePersonDetails(personID, updatedPerson)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Person with ID %s not found", personID)
			return
		}
		json.NewEncoder(w).Encode(updatedPerson)
	}
}

func deletePerson(personUsecase *usecase.PersonUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		personID := params["id"]
		deletedPerson, err := personUsecase.DeletePerson(personID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Person with ID %s not found", personID)
			return
		}
		json.NewEncoder(w).Encode(deletedPerson)
	}
}

func getPeopleByNameHandler(personUsecase *usecase.PersonUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		personName := r.URL.Query().Get("name")
		people := personUsecase.GetByName(personName)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(people)
	}
}
func getPersonAgeHandler(personUsecase *usecase.PersonUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		personID := params["id"]
		age, err := personUsecase.GetAgeByID(personID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Person with ID %s not found", personID)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]int{"age": age})
	}
}
