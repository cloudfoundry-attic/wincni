package main_test

import (
	"path/filepath"

	"github.com/containernetworking/cni/libcni"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("wincni", func() {
	var (
		CNIConfig     libcni.CNI
		networkConfig *libcni.NetworkConfig
		runtimeConf   *libcni.RuntimeConf
		err           error
	)

	BeforeEach(func() {
		CNIConfig = &libcni.CNIConfig{Path: []string{filepath.Dir(wincniBin)}}
		runtimeConf = &libcni.RuntimeConf{}

		confJSON := []byte(`{"cniVersion": "0.3.1", "type": "wincni"}`)
		networkConfig, err = libcni.ConfFromBytes(confJSON)
		Expect(err).To(BeNil())
	})

	Describe("add", func() {
		It("returns successfully", func() {
			_, err := CNIConfig.AddNetwork(networkConfig, runtimeConf)
			Expect(err).To(BeNil())
		})
	})

	Describe("delete", func() {
		It("returns successfully", func() {
			Expect(CNIConfig.DelNetwork(networkConfig, runtimeConf)).To(Succeed())
		})
	})
})
