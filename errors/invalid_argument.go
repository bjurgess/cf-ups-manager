package errors

import "fmt"

type InvalidArgument struct {
	Name string
}


func (e *InvalidArgument) Error() string {
	return fmt.Sprintf("Invalid Argument: %s", e.Name)
}