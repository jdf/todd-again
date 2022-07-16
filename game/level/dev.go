package level

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/jdf/todd-again/game/proto"
	"github.com/tanema/gween/ease"
	"google.golang.org/protobuf/encoding/prototext"
)

var (
	Instance *proto.Tuning

	levelInPath = flag.String(
		"level_in_path", "/tmp/level.textproto", "Path to input proto file")
	levelOutPath = flag.String(
		"level_out_path", "/tmp/level-out.textproto", "Path to output proto file")

	// TODO: #4 Move camera and speed steps into level proto.
	CameraTiltEasing = ease.Linear
	Step1            float32
	Step2            float32
)

func init() {
	Load()
}

func Load() {
	log.Printf("Loading level from %s", *levelInPath)
	Instance = &proto.Tuning{}
	buf, err := ioutil.ReadFile(*levelInPath)
	if err != nil {
		log.Printf("Failed to read level from %s: %v", *levelInPath, err)
		log.Printf("Creating default level.")
	} else if err = prototext.Unmarshal(buf, Instance); err != nil {
		panic(err)
	}
	if Instance.Todd == nil {
		Instance.Todd = &proto.Todd{
			Blink: &proto.Blink{},
			Color: &proto.Color{
				C: []float32{233 / 255.0, 180 / 255.0, 30 / 255.0},
			},
		}
	}
	if Instance.World == nil {
		Instance.World = &proto.World{}
		Instance.World.Bg = &proto.Color{
			C: []float32{0, 0, 0},
		}
	}
	if Instance.Camera == nil {
		Instance.Camera = &proto.Camera{}
	}
	Step1 = Instance.Todd.GetMaxVelocity() * .333
	Step2 = Instance.Todd.GetMaxVelocity() * .666
}

func Save() error {
	buf, err := prototext.MarshalOptions{
		Multiline: true,
		Indent:    "    ",
	}.Marshal(Instance)
	if err != nil {
		return err
	}
	log.Printf("Saving level to %s", *levelOutPath)
	if err = ioutil.WriteFile(*levelOutPath, buf, 0644); err != nil {
		// TODO: #3 #2 handle error by trying to save to some generated path
		return err
	}
	return nil
}
