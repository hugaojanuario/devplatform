package errs

import "fmt"

func NotFound(entity string) error {
	return fmt.Errorf("%s not found", entity)
}
