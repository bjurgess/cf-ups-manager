package utils

import "github.com/bjurgess1/cf-ups-manager/errors"

type Space struct {
	Name string
	UserProvidedServices []UserProvidedService
}

func (space *Space) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw rawSpace

	err := unmarshal(&raw)
	if err != nil {
		return err
	}

	space.Name = raw.Name
	space.UserProvidedServices = raw.UserProvidedServices

	return nil
}
func (space *Space) FindUPS(ups string) (*UserProvidedService, error) {

	for _, element := range space.UserProvidedServices {
		if (element.Name == ups) {
			return &element, nil
		}
	}

	return nil, &errors.UPSNotFound{ups}
}