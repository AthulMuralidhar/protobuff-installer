package semvar_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSemvar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Semvar Suite")
}

//func Test_semVar_toProtocURL(t *testing.T) {
//	type fields struct {
//		major int
//		minor int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		{
//			name: "happy flow",
//			fields: fields{
//				major: 24,
//				minor: 2,
//			},
//			want: "https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := SemVar{
//				Major: tt.fields.major,
//				Minor: tt.fields.minor,
//			}
//			if got := s.ToProtocURL(); got != tt.want {
//				t.Errorf("toProtocURL() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
