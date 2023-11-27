package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/97vaibhav/PersonApi/internal/usecase"
	"github.com/gorilla/mux"
)

func RegisterPersonHandlers(router *mux.Router, personUsecase *usecase.PersonUsecase) {
	router.HandleFunc("/people", getPeopleHandler(personUsecase)).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonByID(personUsecase)).Methods("GET")
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
