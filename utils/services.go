package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ErrorHandler func(string, error)


func ParseFile(filename string) (*UPSManifest, error) {
	dat, err := ioutil.ReadFile(filename)

	if (err != nil) {
		return nil, err
	}

	environments := UPSManifest{}

	err = yaml.Unmarshal(dat, &environments)

	if (err != nil) {
		return nil, err
	}

	return &environments, nil
}