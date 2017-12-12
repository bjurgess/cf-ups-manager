package utils

import (
	"github.com/bjurgess1/cf-ups-manager/errors"
	"encoding/json"
)

type UserProvidedService struct {
	Name string `yaml:"name,omitempty"`
	Syslog string `yaml:"syslog,omitempty"`
	RouteService string `yaml:"route-service,omitempty"`
	Credentials map[string]string `yaml:"credentials"`
}

func (ups *UserProvidedService) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw rawUserProvidedService

	err := unmarshal(&raw)
	if err != nil {
		return err
	}

	if len(raw.Credentials) == 0 && raw.Syslog == "" && raw.RouteService == "" {
		return &errors.InvalidUPSError{raw.Name}
	}

	if (len(raw.Credentials) > 0 && raw.Syslog != "") ||
		(len(raw.Credentials) >0 && raw.RouteService != "") ||
		(raw.RouteService != "" && raw.Syslog != "") {
				return &errors.InvalidUPSError{raw.Name}
	}

	ups.Name = raw.Name
	ups.Credentials = raw.Credentials
	ups.Syslog = raw.Syslog
	ups.RouteService = raw.RouteService

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