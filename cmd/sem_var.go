package cmd

import "strconv"

type semVar struct {
	major int
	minor int
	//patch int // FIXME
}

func (s semVar) String() string {
	//return "v" + strconv.Itoa(s.major) + "." + strconv.Itoa(s.minor) + "." + strconv.Itoa(s.patch)
	return "v" + strconv.Itoa(s.major) + "." + strconv.Itoa(s.minor)
}

func (s semVar) toProtocURL() string {
		url := "https://github.com/protocolbuffers/protobuf/releases/download/v24.2/protoc-24.2-linux-x86_64.zip"
}