package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/97vaibhav/PersonApi/internal/usecase"
	"github.com/gorilla/mux"
)

func RegisterPersonHandlers(router *mux.Router, personUsecase *usecase.PersonUsecase) {
	router.HandleFunc("/people", getPeopleHandler(personUsecase)).Methods("GET")
}

func getPeopleHandler(personUsecase *usecase.PersonUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		people := personUsecase.GetAll()
		json.NewEncoder(w).Encode(people)
	}
}
