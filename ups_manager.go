package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"cf-ups-deployer/utils"
	"fmt"
	"io"
)

type ErrorFunction func(string, error)

type UPSDeployer interface {
	Setup(plugin.CliConnection)
	CreateUserProvidedService(utils.UserProvidedService) error
	UpdateUserProvidedService(utils.UserProvidedService) error
	Deploy(utils.Space, []string)
}

type UPSDeploy struct {
	Connection plugin.CliConnection
	Out io.Writer
	ErrorFunc ErrorFunction
}

func (p *UPSDeploy) Setup(connection plugin.CliConnection) {
	p.Connection = connection
}

func (p *UPSDeploy) Deploy(space utils.Space, upses []string) {
	deployUserProvidedServices := []utils.UserProvidedService{}
	if upses != nil {
		for _, upsToDeploy := range upses {
			ups, err := space.FindUPS(upsToDeploy)
			if err != nil {
				p.ErrorFunc("UPS Not found", err)
				return
			}
			deployUserProvidedServices = append(deployUserProvidedServices, *ups)
		}
	} else {
		deployUserProvidedServices = space.UserProvidedServices
	}

	for _, element := range deployUserProvidedServices {
		if err := p.CreateUserProvidedService(element); err != nil {
			err = p.UpdateUserProvidedService(element)

			if err != nil {
				p.ErrorFunc("Error Creating/Updating UPS", err)
			}
		}
	}
}

func (p *UPSDeploy) CreateUserProvidedService(ups utils.UserProvidedService) error {
	jsonString, err := ups.MarshalCredentials()

	if err != nil {
		fmt.Errorf("%s\n", err.Error())
		return err
	}

	_, err = p.Connection.CliCommand("cups", ups.Name, "-p", fmt.Sprintf("'%s'", jsonString))
	return err
}

func (p *UPSDeploy) UpdateUserProvidedService(ups utils.UserProvidedService) error {
	jsonString, err := ups.MarshalCredentials()

	if err != nil {
		fmt.Errorf("%s\n", err.Error())
		return err
	}

	_, err = p.Connection.CliCommand("uups", ups.Name, "-p", fmt.Sprintf("'%s'", jsonString))
	return err
}