package downloader

import (
	"go.uber.org/zap"
	"os"
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
	//
	//It("check for a protoc.zip to be created in the cwd", func() {
	//	got, err := DownloadAndCreateFile(sugar, "https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip", "")
	//
	//	Expect(err).To(BeNil())
	//	Expect(got).ToNot(BeNil())
	//})
	//
	//It("check for a non 200 return with a fake url", func() {
	//	got, err := DownloadAndCreateFile(sugar, "https://github.com/protocolbuffers/proases/download/v24.2/protoc-24.2-linux-x86_64.zip", "")
	//
	//	Expect(err).ToNot(BeNil())
	//	Expect(got).To(BeNil())
	//})

	//		// TODO: check if the file is really a zip
	It("check perms on the file - it should be accessible to everybody", func() {
		got, err := DownloadAndCreateFile(sugar, "https://github.com/protocolbuffers/protobuf/releases/download/v24.4/protoc-24.4-linux-x86_64.zip", "")

		// TODO: the got still hits the nil check for some reason - the readfile is not helping
		Expect(err).ToNot(BeNil())
		gotData, err := os.ReadFile(got.Name())
		Expect(got).To(BeNil())
		Expect(gotData).To(BeNil())
		//		// TODO: check perms on the file - it should be accessible to everybody

		fileInfo, err := got.Stat()
		Expect(err).ToNot(HaveOccurred())
		Expect(fileInfo.Mode().Perm()).To(Equal(os.FileMode(0666)))
	})

})
