package abstract

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpriteSheet interface {
	NumFrames() int
	NumSequences() int
	Sequences() map[string][]int // name -> frame indices
	Texture() *rl.Texture2D
	Frame(int) rl.Rectangle
	SetSequence(name string, frames []int)
	FrameDimensions() (w, h int)
}

type SpriteSheetService interface {
	GetSheet(name string) SpriteSheet
	LoadImage(key string, img interface{}, framesWide, framesTall int)
}
