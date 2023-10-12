package unzip

import (
	"github.com/AthulMuralidhar/protobuff-installer/pkg/cmd/downloader"
	"go.uber.org/zap"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUnzip(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unzip Suite")
}

var _ = Describe("Unzip", func() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	var f *os.File
	var err error

	BeforeAll(func() {
		f, err = downloader.DownloadAndCreateFile(sugar, "https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip", "")
		Expect(err).ToNot(HaveOccurred())
		Expect(f).ToNot(BeNil())
	})

	AfterAll(func() {
		// TODO: this is not working, dunno y
		err := os.Remove(f.Name())
		Expect(err).ToNot(HaveOccurred())
	})

	It("make sure that the unziped file is not empty", func() {
		//		// TODO: make sure that the unziped file is not empty
		//		// TODO: make sure that the namr of the file is protoc

	})

	//		// TODO: make sure that the unzipping happens in the protoc dir
	//		// TODO: make sure that the unziped file is not empty
})
