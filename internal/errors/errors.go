package errors

import "errors"

var (
	ErrPersonNotFound     = errors.New("person Not found in database")
	ErrEmailAlreadyExists = errors.New("email already exist so cant create person")
	ErrInvalidEmail       = errors.New("invalid email format or empty email !! Please check again")
	ErrInvalidBirthday    = errors.New("invalid birthday format")
	ErrEmptyFirstName     = errors.New("first name cannot be empty ! Please provide first name")
)
