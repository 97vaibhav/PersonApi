package errors

import "errors"

var (
	ErrPersonNotFound     = errors.New("person Not found in database")
	ErrEmailAlreadyExists = errors.New("email already exist so cant create person")
)
