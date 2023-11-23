package downloader

import (
	"github.com/AthulMuralidhar/protobuff-installer/pkg/cmd/semvar"
	"go.uber.org/zap"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDownloader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Downloader Suite")
}

//var _ = Describe("DownloadAndCreateFile", func() {
//	logger, _ := zap.NewProduction()
//	defer logger.Sync() // flushes buffer, if any
//	sugar := logger.Sugar()
//	sm := semvar.SemVar{Major: 24, Minor: 2}
//
//	It("check for a protoc.zip to be created in the cwd", func() {
//		got, err := DownloadAndCreateFile(sugar, sm, "https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip", "")
//
//		Expect(err).To(BeNil())
//		Expect(got).ToNot(BeNil())
//	})
//
//	It("check for a non 200 return with a fake url", func() {
//		got, err := DownloadAndCreateFile(sugar, sm, "https://github.com/protocolbuffers/proases/download/v24.2/protoc-24.2-linux-x86_64.zip", "")
//
//		Expect(err).ToNot(BeNil())
//		Expect(got).To(BeNil())
//	})
//
//	It("check perms on the file - it should be accessible to everybody", func() {
//		got, err := DownloadAndCreateFile(sugar, sm, "https://github.com/protocolbuffers/protobuf/releases/download/v24.4/protoc-24.4-linux-x86_64.zip", "")
//		Expect(err).To(BeNil())
//
//		gotData, err := os.ReadFile(got.Name())
//		Expect(got).ToNot(BeNil())
//		Expect(gotData).ToNot(BeNil())
//	})
//})

var _ = Describe("checkIfValid", func() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	//It("check if the given sem var exists on protoc official page", func() {
	//	sm := semvar.SemVar{Major: 25, Minor: 0}
	//	got := checkIfValid(sugar, sm)
	//	Expect(got).To(Equal(sm))
	//})

	It("check if the given sem var exists on protoc official page", func() {
		sm := semvar.SemVar{Major: 30, Minor: 0}
		got := checkIfValid(sugar, sm)
		Expect(got).To(Equal(semvar.SemVar{Major: 25, Minor: 0})) // at the time of writing 25.0 is the max release version

		// notes:
		// the print in the chek function returns a body with the enite html op,
		// most likely will need a parser, which ive come across in the go book
		// so should see how things are done there once again and use it here
	})

})
