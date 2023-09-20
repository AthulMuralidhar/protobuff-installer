package cmd

import (
	"github.com/onsi/gomega"
	"testing"
)

func Test_protocInstaller(t *testing.T) {
	g := gomega.NewWithT(t)

	type args struct {
		sm          semVar
		installPath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "happy flow",
			args: args{
				sm:          semVar{major: 3, minor: 17},
				installPath: ".",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := protocInstaller(tt.args.sm, tt.args.installPath)
			g.Expect(err).To(gomega.BeNil())

		})
	}
}
