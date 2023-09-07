package cmd

import (
	"fmt"
	"strconv"
)

type semVar struct {
	major int
	minor int
	//patch int // FIXME
}

func (s semVar) String() string {
	//return "v" + strconv.Itoa(s.major) + "." + strconv.Itoa(s.minor) + "." + strconv.Itoa(s.patch)
	return strconv.Itoa(s.major) + "." + strconv.Itoa(s.minor)
}

func (s semVar) toProtocURL() string {
		return fmt.Sprintf("https://github.com/protocolbuffers/protobuf/releases/download/v%s/protoc-%s-linux-x86_64.zip", s.String(), s.String())
}