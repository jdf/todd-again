package main

import (
	"log"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

func emitField(g *protogen.GeneratedFile, f *protogen.Field) {
	if f.Desc.Kind() == protoreflect.MessageKind {
		g.P("m.", f.GoIdent, ".Render()")
	}
}

func emitColor(g *protogen.GeneratedFile, m *protogen.Message) {

}

func emitMessage(g *protogen.GeneratedFile, m *protogen.Message) {
	if m.Desc.FullName() == "game.Color" {
		emitColor(g, m)
	} else {
		g.P("func (m *", m.GoIdent, ") Render() {")
		for _, f := range m.Fields {
			emitField(g, f)
		}
		g.P("}")
	}
}

type varTypes struct {
	hasFloat32 bool
	hasFloat64 bool
	hasInt32   bool
	hasString  bool
	hasColor   bool
}

func (v *varTypes) gather(m *protogen.Message) {
	for _, f := range m.Fields {
		switch f.Desc.Kind() {
		case protoreflect.FloatKind:
			v.hasFloat32 = true
		case protoreflect.DoubleKind:
			v.hasFloat64 = true
		case protoreflect.Int32Kind:
			v.hasInt32 = true
		case protoreflect.StringKind:
			v.hasString = true
		case protoreflect.MessageKind:
			if f.Desc.FullName() == "game.Color" {
				v.hasColor = true
			} else {
				v.gather(f.Message)
			}
		default:
			log.Fatalf("unsupported field type: %v", f.Desc.Kind())
		}
	}
}

// generateFile generates go source to render a UI for editing the source proto.
func generateFile(gen *protogen.Plugin, file *protogen.File) {
	filename := file.GeneratedFilenamePrefix + "_dear_imgui.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-imgui. DO NOT EDIT.")
	g.Import("github.com/inkyblackness/imgui-go/v4")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	v := varTypes{}
	for _, m := range file.Messages {
		v.gather(m)
	}
	g.P("var (")
	if v.hasFloat32 {
		g.P("tmpFloat32 float32")
	}
	if v.hasFloat64 {
		g.P("tmpFloat64 float64")
	}
	if v.hasInt32 {
		g.P("tmpInt32 int32")
	}
	if v.hasString {
		g.P("tmpString string")
	}
	if v.hasColor {
		g.P("tmpColor [3]float32")
	}
	g.P(")")
	g.P()

	for _, m := range file.Messages {
		emitMessage(g, m)
	}
	g.Content()
}
