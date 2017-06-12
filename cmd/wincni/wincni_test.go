package main_test

import (
	"path/filepath"

	"github.com/containernetworking/cni/libcni"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/020"
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

		It("is convertable to 0.2.0 and includes the container IP address", func() {
			result, err = CNIConfig.AddNetwork(networkConfig, runtimeConf)
			result020, err := result.GetAsVersion("0.2.0")
			Expect(err).To(BeNil())

			containerIP := result020.(*types020.Result).IP4.IP.IP
			Expect(containerIP).NotTo(BeNil())
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
