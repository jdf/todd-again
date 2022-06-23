package tuning

import (
	"flag"
	"io/ioutil"

	"google.golang.org/protobuf/encoding/prototext"
)

var (
	tuningProtoPath = flag.String("tuning_proto_path", "/tmp/", "Path to proto file")
)

func SaveTuning() error {
	buf, err := prototext.MarshalOptions{
		Multiline: true,
		Indent:    "    ",
	}.Marshal(Tuning)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(*tuningProtoPath, buf, 0644); err != nil {
		return err
	}
	return nil
}
