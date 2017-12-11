package utils_test

import (
	. "cf-ups-deployer/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Provided Service", func() {
	var ups UserProvidedService

	BeforeEach(func() {
		ups = UserProvidedService{"UPS1", map[string]string{"Credential1":"1", "Credential2":"2"}}
	})

	Context("Marshal Credentials", func() {
		It("Should return valid string in json format", func() {
			Expect(ups.MarshalCredentials()).To(Equal("{\"Credential1\":\"1\",\"Credential2\":\"2\"}"))
		})
	})
})
