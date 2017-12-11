package utils

import "github.com/bjurgess1/cf-ups-manager/errors"

type UPSManifest struct {
	Spaces []Space `yaml:"spaces"`
}

func (upsManifest *UPSManifest) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw rawUPSManifest

	err := unmarshal(&raw)
	if err != nil {
		return err
	}

	upsManifest.Spaces = raw.Spaces
	return nil
}

func (upsManifest *UPSManifest) FindSpace(cliSpace string) (*Space, error) {
	for _, space := range upsManifest.Spaces {
		if space.Name == cliSpace {
			return &space, nil
		}
	}

	return nil, &errors.SpaceNotFoundError{cliSpace}
}