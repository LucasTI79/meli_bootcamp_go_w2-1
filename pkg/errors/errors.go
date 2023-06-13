package apperr

import "fmt"

// Internal Server Error
type InternalServerError struct {
	message string
}

func (e InternalServerError) Error() string {
	return e.message
}

func NewInternalServerError() error {
	return &InternalServerError{message: "an internal error ocurred"}
}

// Resource Not Found
type ResourceNotFound struct {
	message string
}

func (e ResourceNotFound) Error() string {
	return e.message
}

func NewResourceNotFound(message string, args ...interface{}) error {
	return &ResourceNotFound{message: fmt.Sprintf(message, args...)}
}

// Resource Already Exists
type ResourceAlreadyExists struct {
	message string
}

func (e ResourceAlreadyExists) Error() string {
	return e.message
}

func NewResourceAlreadyExists(message string, args ...interface{}) error {
	return &ResourceAlreadyExists{message: fmt.Sprintf(message, args...)}
}

// Resource Already Exists
type RequiredField struct {
	message string
}

func (e RequiredField) Error() string {
	return e.message
}

func NewRequiredField(fieldName string) error {
	return &RequiredField{message: fmt.Sprint(fieldName, " is required")}
}
