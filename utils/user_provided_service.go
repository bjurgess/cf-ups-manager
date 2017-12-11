package utils

import (
	"github.com/bjurgess1/cf-ups-manager/errors"
	"encoding/json"
)

type UserProvidedService struct {
	Name string `yaml:"name,omitempty"`
	Credentials map[string]string `yaml:"credentials"`
}

func (ups *UserProvidedService) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw rawUserProvidedService

	err := unmarshal(&raw)
	if err != nil {
		return err
	}

	if len(raw.Credentials) == 0 {
		return &errors.InvalidUPSError{raw.Name}
	}

	ups.Name = raw.Name
	ups.Credentials = raw.Credentials

	return nil
}

func (ups *UserProvidedService) MarshalCredentials() (string, error) {
	jsonBytes, err := json.Marshal(ups.Credentials)

	if err != nil {
		return "", err
	}

	jsonString := string(jsonBytes[:])

	return jsonString, nil
}