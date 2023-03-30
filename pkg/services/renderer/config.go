package renderer

import rl "github.com/gen2brain/raylib-go/raylib"

type config struct {
	window
	fps int32
}

type window struct {
	Title  string
	Width  int32
	Height int32
}

func (*config) Default() *config {
	const fps60 = 60

	ww, wh := getScreenWidthAndHeight()

	return &config{
		window: window{
			Title:  "example",
			Width:  ww,
			Height: wh,
		},
		fps: fps60,
	}
}

func getScreenWidthAndHeight() (int32, int32) {
	width := int32(rl.GetMonitorWidth(0))
	height := int32(rl.GetMonitorHeight(0))
	return width, height
}
