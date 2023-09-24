package cmd_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCmd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cmd Suite")
}

//func Test_protocInstaller(t *testing.T) {
//	g := gomega.NewWithT(t)
//
//	type args struct {
//		sm          semvar.semVar
//		installPath string
//		want        func()
//	}
//	tests := []struct {
//		name string
//		args args
//		want func()
//	}{
//		{
//			name: "happy flow",
//			args: args{
//				sm:          semvar.semVar{major: 3, minor: 17},
//				installPath: ".",
//			},
//			want: func() {
//
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			err := main.protocInstaller(tt.args.sm, tt.args.installPath)
//			g.Expect(err).To(gomega.BeNil())
//
//		})
//	}
//}
