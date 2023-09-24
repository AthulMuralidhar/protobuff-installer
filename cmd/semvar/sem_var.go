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

func (s SemVar) String() string {
	//return "v" + strconv.Itoa(s.major) + "." + strconv.Itoa(s.minor) + "." + strconv.Itoa(s.patch)
	return strconv.Itoa(s.Major) + "." + strconv.Itoa(s.Minor)
}

func (s SemVar) ToProtocURL() string {
	return fmt.Sprintf("https://github.com/protocolbuffers/protobuf/releases/download/v%s/protoc-%s-linux-x86_64.zip", s.String(), s.String())
}
