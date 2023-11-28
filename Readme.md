# Hello To PersonApi

## This is a simple Golang Api made with idea of clean architecture in mind which is doing CRUD operation on a Person struct .

## Functionalities Provided :
The API is doing the following :

- Respond with JSON output.
- Implement route `GET /people` with a 200 response that contains all people in the system
- Implement route `GET /people/:id` with a 200 response containing the requested person or a 404 response if the person does not exist.
- Implement route `GET /people/:id/age` with a 200 response that calculates and returns the age of the person on the current date.
- Implement route `Get /people?name=:name` with a 200 response that contains the people whose first or last name meets the search criteria. If there are no results, it should return an empty array. Ex. If there are two people, Jack and Jill, the following request `/people?name=j` should return both Jack and Jill, but `/people?name=ja` should only return Jack.
- Implement route `POST /people` that creates 1 or more persons. If the email address of the person provided already exists in the data set, it should return an error explaining why the request failed. If creation is successful, it should return a 200 response and the created person.
- Implement route `PUT /people/:id` that updates the person with the provided id. It should return a 200 response and the updated person if it successfully updates or a 404 response if the person does not exist.
- Implement route `DELETE /people/:id` that deletes the person with the provided id. It should return a 200 response and the deleted person if it successfully deletes or a 404 response if the person does not exists.

> [!NOTE]
> The application is using in memeory data to get the things done .


### The defination of Person struct is given as
```
type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
}
```

# Project structure
```
├── Readme.md
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── internal
    ├── delivery
    │   ├── handler.go //handling routes
    │   └── routes.go // declaring routes
    ├── domain
    │   └── person.go //person struct defined here
    ├── errors
    │   └── errors.go
    ├── repository
    │   ├── person_repository.go // all database operation goes here (in memory for this project)
    │   └── person_repository_test.go
    └── usecase
        ├── helpers.go
        ├── person_usecase.go // business logic
        └── person_usecase_test.go
```
## To run the application
```
cd cmd
run main.go
```

## To test the application code

```
go test ./...
```

## Sample Input Output in JSON

- Run the program
- Go to postman
- Hit the endpoint at localhost:8080

### For `GET /people`

Request
```
localhost:8080/people
```

Response
```
[
    {
        "id": "bf552a1c-fd73-4bd0-b64a-d3f69a9ff9de",
        "firstName": "John",
        "lastName": "Doe",
        "email": "johndoe@example.com",
        "birthday": "1997-01-01"
    },
    {
        "id": "d5356358-b39f-4c6e-9690-2c965a607702",
        "firstName": "Jane",
        "lastName": "Doe",
        "email": "janedoe@example.com",
        "birthday": "1991-07-28"
    },
    {
        "id": "cb2bfa60-e2ae-46ec-ad77-60cf7e8979fd",
        "firstName": "Brian",
        "lastName": "Smith",
        "email": "briansmith@example.com",
        "birthday": "2000-05-10"
    },
    {
        "id": "d82fc695-5ac2-4fed-9387-a7d9c0fb0c4f",
        "firstName": "Ashley",
        "lastName": "Yu",
        "email": "ashleyyu@example.com",
        "birthday": "2003-12-24"
    }
]
```
### For  `GET /people/:id/age`

Request:
```
localhost:8080/people/bf552a1c-fd73-4bd0-b64a-d3f69a9ff9de/age
```
Response:
```
{
    "age": 26
}
```
> **_NOTE:_**
>Try Hitting different endpoint and experience Yourself.
