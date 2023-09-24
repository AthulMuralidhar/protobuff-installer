package downloader_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDownloader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Downloader Suite")
}

var _ = Describe("Downloader", func() {

})

//
//import (
//"go.uber.org/zap"
//"os"
//"reflect"
//"testing"
//)
//
//func Test_downloadAndCreateFile(t *testing.T) {
//	type args struct {
//		sugar *zap.SugaredLogger
//		url   string
//		cwd   string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *os.File
//		wantErr bool
//	}{
//		// TODO: check for a protoc.zip to be created in the cwd
//		// TODO: check for a non 200 return with a fake url
//		// TODO: check if the file is not empty
//		// TODO: check perms on the file - it should be accessible to everybody
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := DownloadAndCreateFile(tt.args.sugar, tt.args.url, tt.args.cwd)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("downloadAndCreateFile() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("downloadAndCreateFile() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
