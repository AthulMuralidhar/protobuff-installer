package unzip_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUnzip(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unzip Suite")
}

//func Test_unzipToDir(t *testing.T) {
//	g := gomega.NewWithT(t)
//	type args struct {
//		source string
//		dest   string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: make sure that the unzipping happens in the protoc dir
//		{
//			name:    "happy flow",
//			args:    args{
//				source: "",
//				dest:   "",
//			},
//			wantErr: false,
//		},
//	}
//
//	err := globalSetup()
//	g.Expect(err).ToNot(g.)
//
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := UnzipToDir(tt.args.source, tt.args.dest); (err != nil) != tt.wantErr {
//				t.Errorf("unzipToDir() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
