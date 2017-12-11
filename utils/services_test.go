package utils_test

import (
	. "cf-ups-manager/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Services Test Suite", func() {
	Describe("Success cases", func() {
		var filename string

		BeforeEach(func() {
			filename = "../fixtures/ups-list.yml"
		})

		Context("Environments UPS Parser", func() {
			var env *UPSManifest
			var err error

			BeforeEach(func() {
				env, err = ParseFile(filename)
			})

			It("Should Return new Environment", func() {
				Expect(env).ToNot(BeNil())
			})

			It("Should have two keys", func() {
				Expect(len(env.Spaces)).To(Equal(3))
				Expect(env.Spaces[0].Name).To(Equal("dev"))
				Expect(env.Spaces[1].Name).To(Equal("qa"))
				Expect(env.Spaces[2].Name).To(Equal("stage"))
			})

			Context("Dev UPS", func() {
				var UPSes []UserProvidedService

				BeforeEach(func() {
					UPSes = env.Spaces[0].UserProvidedServices
				})

				It("Should have 2 user provided services", func() {
					Expect(len(UPSes)).To(Equal(2))
				})

				It("Should have a UPS named UPS2", func() {
					Expect(UPSes[1].Name).To(Equal("UPS2"))
				})

				Context("UPS1", func() {
					var ups UserProvidedService

					BeforeEach(func() {
						ups = env.Spaces[0].UserProvidedServices[0]
					})

					It("Should have a UPS named named UPS1", func() {
						Expect(ups.Name).To(Equal("UPS1"))
					})

					It("Should have 3 credentials", func() {
						Expect(len(ups.Credentials)).To(Equal(3))
						Expect(ups.Credentials["Credential1"]).To(Equal("1"))
						Expect(ups.Credentials["Credential2"]).To(Equal("2"))
						Expect(ups.Credentials["Credential3"]).To(Equal("3"))
					})
				})
			})

			Context("qa UPS", func() {
				var UPSes []UserProvidedService

				BeforeEach(func() {
					UPSes = env.Spaces[1].UserProvidedServices
				})

				It("Should have 1 UPS", func() {
					Expect(len(UPSes)).To(Equal(1))
				})
			})

			Context("stage UPS", func() {
				var UPSes []UserProvidedService

				BeforeEach(func() {
					UPSes = env.Spaces[2].UserProvidedServices
				})

				It("Should have 0 UPSes", func() {
					Expect(len(UPSes)).To(Equal(0))
				})
			})
		})
	})

	Describe("Failure Cases", func() {
		var filename string
		BeforeEach(func() {
			filename = "../fixtures/ups-no-credentials-list.yml"
		})

		Context("Parse invalid Yaml file", func() {
			var env *UPSManifest
			var err error

			BeforeEach(func() {
				env, err = ParseFile(filename)
			})

			It("Should return error not null", func() {
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("Invalid User Provided Service: UPS1"))
			})
		})
	})
})
