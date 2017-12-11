package errors

import "fmt"

type SpaceNotFoundError struct {
	Name string
}

func (e *SpaceNotFoundError) Error() string {
	return fmt.Sprintf("Unable to find Space: %s in the manifest", e.Name)
}