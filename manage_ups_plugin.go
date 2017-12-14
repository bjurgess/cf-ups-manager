package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"github.com/bjurgess1/cf-ups-manager/utils"
	"github.com/bjurgess1/cf-ups-manager/errors"
	"fmt"
	"log"
	"os"
)

var PluginVersion string

type DeployUPSCmd struct {
	Deployer UPSDeployer
	ErrorFunc ErrorFunction
}

func (p *DeployUPSCmd) GetMetadata() plugin.PluginMetadata {
	var major, minor, build int
	fmt.Sscanf(PluginVersion, "%d.%d.%d", &major, &minor, &build)

	return plugin.PluginMetadata{
		Name: "user-provided-service-manager",
		Version: plugin.VersionType{
			Major: major,
			Minor: minor,
			Build: build,
		},
		Commands: []plugin.Command{
			{
				Name: "ups-manager",
				Alias: "upsm",
				HelpText: "Manage, create, and update multiple user provided services across multiple environments",
				UsageDetails: plugin.Usage{
					Usage: "deploy-ups [-f UPS_MANIFEST_FILE] -u [USER_PROVIDED_SERVICES]",
					Options: map[string]string {
						"f": "The location of the ups manifest file",
						"u": "Comma separated list of specific User Provided Services to create/update. Leave blank to create/update everything",
					},
				},
			},
		},
	}
}

func (p *DeployUPSCmd) Run(cliConnection plugin.CliConnection, args []string) {
	if len(args) > 0 && args[0] == "CLI-MESSAGE-UNINSTALL" {
		return
	}

	arguments := ParseArguments(args)

	if arguments.UPSManifestPath == "" {
		p.ErrorFunc("", &errors.InvalidArgument{"Manifest Path"})
		return
	}

	manifest, err := utils.ParseFile(arguments.UPSManifestPath)
	if err != nil {
		p.ErrorFunc("", err)
		return
	}

	space, err := cliConnection.GetCurrentSpace()
	if err != nil {
		p.ErrorFunc("", err)
		return
	}

	manifestSpace, err := manifest.FindSpace(space.Name)
	if err != nil {
		p.ErrorFunc("", err)
		return
	}
	p.Deployer.Setup(cliConnection)
	p.Deployer.Deploy(*manifestSpace, arguments.UPSToDeploy)
	return
}

func main() {
	deployer := UPSDeploy{
		Out: os.Stdout,
		ErrorFunc: func(s string, e error) {
			log.Fatalf("%v", e)
		},
	}
	cmd := DeployUPSCmd{
		&deployer,
		func(s string, e error) {
			log.Fatalf("%v", e)
		},
	}
	plugin.Start(&cmd)
}
