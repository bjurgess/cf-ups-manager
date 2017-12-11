package main_test

import (
	"bytes"
	. "cf-ups-deployer"
	"os/exec"

	"code.cloudfoundry.org/cli/plugin/models"
	"code.cloudfoundry.org/cli/plugin/pluginfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UPS Deployer", func() {
	var (
		cmd DeployUPSCmd
		deployer UPSDeploy
		cliConnection *pluginfakes.FakeCliConnection
		expectedException []error
		errorFunc func(message string, err error)
	)

	BeforeEach(func() {
		expectedException = []error{}
		errorFunc = func(message string, err error) {
			expectedException = append(expectedException, err)
		}
		cliConnection = &pluginfakes.FakeCliConnection{}
		deployer = UPSDeploy{
			Connection: nil,
			Out: &bytes.Buffer{},
			ErrorFunc: errorFunc,
		}
		cmd = DeployUPSCmd{
			Deployer: &deployer,
			ErrorFunc: errorFunc,
		}
	})

	Context("Run", func() {
		It("Should throw exception with invalid manifest argument", func() {
			cmd.Run(cliConnection, []string{"asdf", "-u", "args, one"})
			Expect(expectedException[0]).To(HaveOccurred())
		})

		It("Should throw exception with invalid manifest yaml", func() {
			cmd.Run(cliConnection, []string{"asdf", "-f", "fixtures/ups-invalid-yaml.yml"})
			Expect(expectedException[0]).To(HaveOccurred())
		})

		It("Should throw exception with invalid credentials", func() {
			cmd.Run(cliConnection, []string{"asdf", "-f", "fixtures/ups-no-credentials-list.yml"})
			Expect(expectedException[0]).To(HaveOccurred())
		})

		It("Should throw exception on invalid space", func() {
			cliConnection.GetCurrentSpaceStub = func() (plugin_models.Space, error) {
				return plugin_models.Space{plugin_models.SpaceFields{"GUID", "NAME"}}, nil
			}
			cmd.Run(cliConnection, []string{"asdf", "-f", "fixtures/ups-list.yml"})
			Expect(expectedException[0]).To(HaveOccurred())
			Expect(expectedException[0].Error()).To(Equal("Unable to find Space: NAME in the manifest"))
		})

		It("Should throw exception on get current space error", func() {
			cliConnection.GetCurrentSpaceStub = func() (plugin_models.Space, error) {
				var space plugin_models.Space
				return space, &exec.Error{"Not logged In", nil}
			}
			cmd.Run(cliConnection, []string{"asdf", "-f", "fixtures/ups-list.yml"})
			Expect(expectedException[0]).To(HaveOccurred())
		})

		It("Should call CUPS", func() {
			cliConnection.GetCurrentSpaceStub = func() (plugin_models.Space, error) {
				return plugin_models.Space{plugin_models.SpaceFields{"GUID", "dev"}}, nil
			}
			cmd.Run(cliConnection, []string{"asdf", "-f", "fixtures/ups-list.yml"})
			cfCommands := getAllCfCommands(cliConnection)
			Expect(cfCommands).To(Equal([]string{
				"cups UPS1 -p '{\"Credential1\":\"1\",\"Credential2\":\"2\",\"Credential3\":\"3\"}'",
				"cups UPS2 -p '{\"Credential1\":\"1\",\"Credential2\":\"2\",\"Credential3\":\"3\"}'",
			}))
		})

		It("Should call CUPS w/ only UPS1", func() {
			cliConnection.GetCurrentSpaceStub = func() (plugin_models.Space, error) {
				return plugin_models.Space{plugin_models.SpaceFields{"GUID", "dev"}}, nil
			}
			cmd.Run(cliConnection, []string{"asdf", "-f", "fixtures/ups-list.yml", "-u", "UPS1"})
			cfCommands := getAllCfCommands(cliConnection)
			Expect(cfCommands).To(Equal([]string{
				"cups UPS1 -p '{\"Credential1\":\"1\",\"Credential2\":\"2\",\"Credential3\":\"3\"}'",
			}))
		})

		It("Should not call cups on non-existent ups", func() {
			cliConnection.GetCurrentSpaceStub = func() (plugin_models.Space, error) {
				return plugin_models.Space{plugin_models.SpaceFields{"GUID", "dev"}}, nil
			}
			cmd.Run(cliConnection, []string{"asdf", "-f", "fixtures/ups-list.yml", "-u", "NonExistentUPS"})
			cfCommands := getAllCfCommands(cliConnection)
			Expect(cfCommands).To(Equal([]string{
			}))
			Expect(expectedException[0]).To(HaveOccurred())
		})
	})
})
