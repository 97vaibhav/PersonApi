package main

import (
	"fmt"
	"net/http"

	"github.com/97vaibhav/PersonApi/internal/delivery"
	"github.com/97vaibhav/PersonApi/internal/repository"
	"github.com/97vaibhav/PersonApi/internal/usecase"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	personRepo := repository.NewPersonRepository()
	personUsecase := usecase.NewPersonUsecase(personRepo)
	delivery.RegisterPersonHandlers(router, personUsecase)

	http.Handle("/", router)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
