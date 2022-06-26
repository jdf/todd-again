package tuning

import (
	"flag"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/encoding/prototext"
)

var (
	tuningProtoPath = flag.String("tuning_proto_path", "/tmp/tuning.textproto", "Path to proto file")
)

func Save() error {
	buf, err := prototext.MarshalOptions{
		Multiline: true,
		Indent:    "    ",
	}.Marshal(Instance)
	if err != nil {
		return err
	}
	log.Printf("Saving tuning to %s", *tuningProtoPath)
	if err = ioutil.WriteFile(*tuningProtoPath, buf, 0644); err != nil {
		return err
	}
	return nil
}
