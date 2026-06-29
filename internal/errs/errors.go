package errs

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
	ErrInvalid  = errors.New("invalid")
	ErrRequired = errors.New("required")
)

func NotFound(resource string) error {
	return fmt.Errorf("%s: %w", resource, ErrNotFound)
}

func Invalid(field string) error {
	return fmt.Errorf("%s: %w", field, ErrInvalid)
}

func Required(field string) error {
	return fmt.Errorf("%s: %w", field, ErrRequired)
}

func Wrap(operation string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", operation, err)
}
