package domain

import (
	"errors"
	"fmt"
)

// user errors
var (
	ErrUserNotFound                    = errors.New("user doesn't exists")
	ErrUserNotFoundOrSessionWasExpired = errors.New("user doesn't exists or session was expired")
	ErrUserAlreadyExists               = errors.New("user with such login already exists")
	ErrUserBadPassword                 = errors.New("bad password")

	ErrSessionNotFound      = errors.New("session was not found")
	ErrSessionAlreadyExists = errors.New("session is already exist")
)

// ErrDataNotFound materials error
var (
	ErrDataNotFound = errors.New("data was not found")
)

// ErrInternal common error
var ErrInternal = errors.New("internal error")

// Error structs
type (
	NotFoundError struct {
		Err error
	}
	AlreadyExistsError struct {
		Err error
	}
	StatementPSQLError struct {
		Err error
	}
	ExecutionPSQLError struct {
		Err error
	}
)

// Method of NotFoundError struct
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: not found in storage\n", e.Err.Error())
}

// Method of AlreadyExistsError struct
func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("%s: already exists in storage\n", e.Err.Error())
}

// Method of StatementPSQLError struct
func (e *StatementPSQLError) Error() string {
	return fmt.Sprintf("%s: could not compile statement\n", e.Err.Error())
}

// Method of ExecutionPSQLError struct
func (e *ExecutionPSQLError) Error() string {
	return fmt.Sprintf("%s: could not query", e.Err.Error())
}

// Unwrap Method of NotFoundError struct
func (e *NotFoundError) Unwrap() error {
	return e.Err
}

// Unwrap Method of AlreadyExistsError struct
func (e *AlreadyExistsError) Unwrap() error {
	return e.Err
}

// Unwrap Method of StatementPSQLError struct
func (e *StatementPSQLError) Unwrap() error {
	return e.Err
}

// Unwrap Method of ExecutionPSQLError struct
func (e *ExecutionPSQLError) Unwrap() error {
	return e.Err
}
