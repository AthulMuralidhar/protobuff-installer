package downloader

import (
	"go.uber.org/zap"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDownloader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Downloader Suite")
}

var _ = Describe("Downloader", func() {

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	It("check for a protoc.zip to be created in the cwd", func() {
		got, err := DownloadAndCreateFile(sugar, "https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip", "")

		Expect(err).To(BeNil())
		Expect(got).ToNot(BeNil())
	})

	It("check for a non 200 return with a fake url", func() {
		got, err := DownloadAndCreateFile(sugar, "https://github.com/protocolbuffers/proases/download/v24.2/protoc-24.2-linux-x86_64.zip", "")

		Expect(err).ToNot(BeNil())
		Expect(got).To(BeNil())
	})

	//		// TODO: check perms on the file - it should be accessible to everybody
	//		// TODO: check if the file is really a zip
	It("check perms on the file - it should be accessible to everybody", func() {
		got, err := DownloadAndCreateFile(sugar, "https://github.com/protocolbuffers/proases/download/v24.2/protoc-24.2-linux-x86_64.zip", "")

		Expect(err).ToNot(BeNil())
		Expect(got).To(BeNil())
	})

})
