package models

import "strings"

const (
	// ErrNotFound is  returned  when a resource cannot be found
	// in the database
	ErrNotFound     modelError = "models: resource not found"
	ErrInvalidEmail modelError = "models: invalid email address provided"
	// ErrPasswordIncorrect is returned when an invaid password is used  when attempting to authenticate a user
	ErrPasswordIncorrect modelError = "models: incorrect password provided"
	// ErrEmailRequired is returned when an email address is not provided when creating a user
	ErrEmailRequired modelError = "models: Email address is required"
	//ErrEmailInvalid is returned when an email address provided does not match Regex requirments
	ErrEmailInvalid modelError = "models: Email address is not valid"
	// ErrEmailTaken is returned when an update or create is attempted with an email adress that is already in use
	ErrEmailTaken modelError = "models: Email address is already taken"
	// ErrPasswordRequired is returned when a create is attempted without a user password provided
	ErrPasswordRequired modelError = "Models: Password is required"
	// ErrPasswordTooShort is returned when an update or create is attempted with a user password that is less than 8 characters
	ErrPasswordTooShort modelError = "models: Password must be 8 characters long"
	ErrTitleRequired    modelError = "models: title is required"
	// ErrIDInvalid is returned when an invalid ID is provided to a method like Delete.
	ErrIDInvalid privateError = "models: ID provided was invalid"
	// ErrRememberRequired  is returned when a create or update is attempted without a user remember token hash
	ErrRememberRequired privateError = "Models: Remember token is required"
	// ErrRememberTooShort is returned when a remember token is not at least 32 bytes
	ErrRememberTooShort privateError = "models: Remember token must be at least 32 bytes"
	ErrUserIDRequired   privateError = "models: user ID is required"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	split := strings.Split(s, " ")
	split[0] = strings.Title(split[0])
	return strings.Join(split, " ")
}

type privateError string

func (e privateError) Error() string {
	return string(e)
}
