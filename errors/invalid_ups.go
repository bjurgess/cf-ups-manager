package errors

import (
	"fmt"
)

type InvalidUPSError struct {
	Name string
}

func (e *InvalidUPSError) Error() string {
	return fmt.Sprintf("Invalid User Provided Service: %s", e.Name)
}