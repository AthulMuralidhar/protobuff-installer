package cmd

import "testing"

func Test_protocInstaller(t *testing.T) {
	type args struct {
		installProtoc bool
		sm            semVar
		installPath   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "happy flow",
			args: args{
				installProtoc: true,
				sm:            semVar{major: 3, minor: 17},
				installPath:   ".",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			protocInstaller(tt.args.installProtoc, tt.args.sm, tt.args.installPath)
		})
	}
}
