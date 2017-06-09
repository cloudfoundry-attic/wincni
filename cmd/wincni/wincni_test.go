package main_test

import (
	"fmt"
	"path/filepath"

	"github.com/containernetworking/cni/libcni"
	"github.com/containernetworking/cni/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("wincni", func() {
	var (
		CNIConfig     libcni.CNI
		networkConfig *libcni.NetworkConfig
		runtimeConf   *libcni.RuntimeConf
	)

	BeforeEach(func() {
		CNIConfig = &libcni.CNIConfig{Path: []string{filepath.Dir(wincniBin)}}

		networkConfig = &libcni.NetworkConfig{
			Network: &types.NetConf{
				CNIVersion: "0.3.1",
				Type:       "wincni",
			},
		}

		runtimeConf = &libcni.RuntimeConf{}
	})

	It("returns successfully on ADD", func() {
		output, err := CNIConfig.AddNetwork(networkConfig, runtimeConf)
		fmt.Println(output)
		Expect(err).To(BeNil())
	})
})
