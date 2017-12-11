package utils_test

import (
	. "cf-ups-manager/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Provided Service", func() {
	var userProvidecServices []UserProvidedService
	var spaces []Space
	var manifest UPSManifest

	BeforeEach(func() {
		spaces = []Space{
			{
				Name: "TEST",
				UserProvidedServices: userProvidecServices,
			},
			{
				Name: "TEST2",
				UserProvidedServices: userProvidecServices,
			},
		}

		manifest = UPSManifest{spaces}
	})

	Context("FindUPS", func() {
		It("Should return proper UPS", func() {
			space, _ := manifest.FindSpace("TEST")
			Expect(space.Name).To(Equal("TEST"))
		})

		It("Should return error on non-existent UPS", func() {
			_, err := manifest.FindSpace("NonExistentSpace")
			Expect(err.Error()).To(Equal("Unable to find Space: NonExistentSpace in the manifest"))
		})
	})
})
