package errors

import "fmt"

type UPSNotFound struct {
	Name string
}


func (e *UPSNotFound) Error() string {
	return fmt.Sprintf("Unable to find User Provided Service: %s", e.Name)
}