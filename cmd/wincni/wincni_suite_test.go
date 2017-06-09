package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var wincniBin string

func TestWincni(t *testing.T) {
	BeforeSuite(func() {
		var err error
		wincniBin, err = gexec.Build("code.cloudfoundry.org/wincni/cmd/wincni")
		Expect(err).To(BeNil())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Wincni Suite")
}
