package semvar

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSemvar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Semvar Suite")
}

var _ = Describe("Semvar", func() {
	It("happy flow", func() {
		sm := newSemVar(24, 2)

		got := sm.ToProtocURL()

		Expect(got).To(Equal("https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip"))
	})
})
