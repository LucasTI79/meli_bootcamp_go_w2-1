package apperr

import (
	"errors"
	"fmt"
)

// Resource Not Found
type ResourceNotFound struct {
	message string
}

func (e ResourceNotFound) Error() string {
	return e.message
}

func NewResourceNotFound(message string, args ...interface{}) *ResourceNotFound {
	return &ResourceNotFound{message: fmt.Sprintf(message, args...)}
}

type DependentResourceNotFound struct {
	message string
}

func (e DependentResourceNotFound) Error() string {
	return e.message
}

func NewDependentResourceNotFound(message string, args ...interface{}) *DependentResourceNotFound {
	return &DependentResourceNotFound{message: fmt.Sprintf(message, args...)}
}

// Resource Already Exists
type ResourceAlreadyExists struct {
	message string
}

func (e ResourceAlreadyExists) Error() string {
	return e.message
}

func NewResourceAlreadyExists(message string, args ...interface{}) *ResourceAlreadyExists {
	return &ResourceAlreadyExists{message: fmt.Sprintf(message, args...)}
}

func Is[T error](err error) bool {
	var comparisonErr T
	return errors.As(err, &comparisonErr)
}
