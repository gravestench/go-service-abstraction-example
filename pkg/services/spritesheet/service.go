package spritesheet

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

type Service struct {
	render  abstract.RenderService
	texture abstract.TextureService
	sheets  map[string]*sheet
}

func (s *Service) Init(possibleDependencies *[]interface{}) {
	s.sheets = make(map[string]*sheet)
	s.populateDependencies(possibleDependencies)
}

func (s *Service) Name() string {
	return "Spritesheet Loader"
}

func (s *Service) LoadImage(key string, imgAbstract interface{}, w, h int) {
	if _, found := s.sheets[key]; found {
		return
	}

	var img image.Image

	switch v := imgAbstract.(type) {
	case string:
		// Open the image file
		file, err := os.Open(v)
		if err != nil {
			panic(err)
		}

		// Decode the image file into an image.Image object
		img, _, err = image.Decode(file)
		if err != nil {
			panic(err)
		}
		file.Close()
	case image.Image:
		img = v
	}

	sh := &sheet{
		rows:      h,
		columns:   w,
		sequences: make(map[string][]int),
		frames:    make([]rl.Rectangle, w*h),
	}

	mainthread.Call(func() {
		tx := rl.LoadTextureFromImage(rl.NewImageFromImage(&imageBugHack{img: img}))
		sh.texture = &tx
	})

	s.sheets[key] = sh

	numFrames := sh.columns * sh.rows

	// Define the dimensions of each sprite frame
	frameW := int(sh.texture.Width) / sh.columns
	frameH := int(sh.texture.Height) / sh.rows

	// Create an array of rectangles that define the location of each sprite in the sprite sheet
	sh.frames = make([]rl.Rectangle, sh.columns*sh.rows)
	for i := 0; i < numFrames; i++ {
		x := float32((i % sh.columns) * frameW)
		y := float32((i / sh.rows) * frameH)
		sh.frames[i] = rl.NewRectangle(x, y, float32(frameW), float32(frameH))
	}
}

func (s *Service) GetSheet(name string) abstract.SpriteSheet {
	return s.sheets[name]
}
