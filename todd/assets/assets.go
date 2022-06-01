package assets

import (
	"embed"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

//go:embed MoonFlowerBold.ttf
//go:embed InstructionBold.ttf
var assets embed.FS

// GetOrDie returns the named asset, or panics if it doesn't exist.
func GetOrDie(name string) []byte {
	b, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return b
}

// GetFontOrDie returns the named font, or panics if it doesn't exist.
func GetFontOrDie(name string) *truetype.Font {
	f, err := freetype.ParseFont(GetOrDie(name))
	if err != nil {
		panic(err)
	}
	return f
}
