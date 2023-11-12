package semvar

import (
	"fmt"
	"strconv"
)

type SemVar struct {
	Major int
	Minor int
	//patch int // FIXME
}

func newSemVar(major, minor int) SemVar {
	return SemVar{
		Major: major,
		Minor: minor,
	}
}

func (s SemVar) String() string {
	return strconv.Itoa(s.Major) + "." + strconv.Itoa(s.Minor)
}

func (s SemVar) ToProtocURL() string {
	return fmt.Sprintf("https://github.com/protocolbuffers/protobuf/releases/download/v%s/protoc-%s-linux-x86_64.zip", s.String(), s.String())
}
