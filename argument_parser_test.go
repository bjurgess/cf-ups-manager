package main_test

import (
	. "cf-ups-manager"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Argument Parser", func() {
	It("Should return argument struct", func() {
		args := ParseArguments([]string{"dups", "-f", "manifest", "-u", "ups1"})
		Expect(args.UPSManifestPath).To(Equal("manifest"))
		Expect(args.UPSToDeploy).To(Equal([]string {
			"ups1",
		}))
	})

	It("Should return argument struct w/ multiple upses", func() {
		args := ParseArguments([]string{"dups", "-f", "manifest", "-u", "ups1,ups2"})
		Expect(args.UPSManifestPath).To(Equal("manifest"))
		Expect(args.UPSToDeploy).To(Equal([]string {
			"ups1",
			"ups2",
		}))
	})

	It("Should return argument struct w/ no upses", func() {
		args := ParseArguments([]string{"dups", "-f", "manifest"})
		Expect(args.UPSManifestPath).To(Equal("manifest"))
		Expect(args.UPSToDeploy).To(BeNil())
	})

	It("Should return argument struct w/ multiple upses ignoring whitespace", func() {
		args := ParseArguments([]string{"dups", "-f", "manifest", "-u", "ups1, ups2"})
		Expect(args.UPSManifestPath).To(Equal("manifest"))
		Expect(args.UPSToDeploy).To(Equal([]string {
			"ups1",
			"ups2",
		}))
	})

	It("Should should return argument struct w/ semi-colon separated", func() {
		args := ParseArguments([]string{"dups", "-f", "manifest", "-u", "ups1; ups2"})
		Expect(args.UPSManifestPath).To(Equal("manifest"))
		Expect(args.UPSToDeploy).To(Equal([]string {
			"ups1",
			"ups2",
		}))
	})
})