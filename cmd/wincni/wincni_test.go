package main_test

import (
	"path/filepath"

	"github.com/containernetworking/cni/libcni"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("wincni", func() {
	var (
		CNIConfig     libcni.CNI
		networkConfig *libcni.NetworkConfig
		runtimeConf   *libcni.RuntimeConf
		err           error
		confJSON      string
		result        types.Result
	)

	JustBeforeEach(func() {
		CNIConfig = &libcni.CNIConfig{Path: []string{filepath.Dir(wincniBin)}}
		runtimeConf = &libcni.RuntimeConf{
			ContainerID: "handle123",
			NetNS:       "path_to_namespace",
			IfName:      "eth123",
		}

		networkConfig, err = libcni.ConfFromBytes([]byte(confJSON))
		Expect(err).To(BeNil())
	})

	Describe("add", func() {
		BeforeEach(func() {
			confJSON = `
{
	"cniVersion": "0.3.1",
	"type": "wincni",
	"name": "container_net",
	"runtimeConfig": {
		"portMappings": [ { "container_port": 8080, "host_port": 1234 } ]
	}
}`
		})

		It("writes the CNI version", func() {
			result, err = CNIConfig.AddNetwork(networkConfig, runtimeConf)
			Expect(err).To(BeNil())
			Expect(result.Version()).To(Equal("0.3.1"))
		})

		It("has at least one IP address", func() {
			result, err = CNIConfig.AddNetwork(networkConfig, runtimeConf)
			Expect(err).To(BeNil())

			resultCurrent := result.(*current.Result)
			Expect(len(resultCurrent.IPs)).To(BeNumerically(">", 0))
		})
	})

	Describe("delete", func() {
		BeforeEach(func() {
			confJSON = `
{
	"cniVersion": "0.3.1",
	"type": "wincni",
	"name": "container_net"
}`
		})
		It("returns successfully", func() {
			Expect(CNIConfig.DelNetwork(networkConfig, runtimeConf)).To(Succeed())
		})
	})
})
