package utils_test

import (
	. "github.com/bjurgess1/cf-ups-manager/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Provided Service", func() {
	var userProvidecServices []UserProvidedService
	var space Space

	BeforeEach(func() {
		userProvidecServices = []UserProvidedService{
			{
				"UPS1",
				map[string]string{"Credential1":"1", "Credential2":"2"},
			},
			{
				"UPS2",
				map[string]string{"Credential3":"1", "Credential4":"2"},
			},
		}

		space = Space{
			Name: "TEST",
			UserProvidedServices: userProvidecServices,
		}
	})

	Context("FindUPS", func() {
		It("Should return proper UPS", func() {
			ups, _ := space.FindUPS("UPS2")
			Expect(ups.Name).To(Equal("UPS2"))
		})

		It("Should return error on non-existent UPS", func() {
			_, err := space.FindUPS("NonExistentUPS")
			Expect(err.Error()).To(Equal("Unable to find User Provided Service: NonExistentUPS"))
		})
	})
})
