package sprite_test

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (s *Scene) Render() {
	if !s.Ready() {
		return
	}

	if s.singleton.sheet == nil {
		return
	}

	tx := s.singleton.sheet.Texture()
	if tx == nil {
		return
	}

	fw, fh := s.singleton.sheet.FrameDimensions()
	w, h := float32(fw*4), float32(fh*4)
	x, y := s.singleton.position.X-float32(w/2), s.singleton.position.Y-float32(h/2)
	rl.DrawTexturePro(
		*tx,
		s.singleton.sheet.Frame(s.getFrameIndex()),
		rl.Rectangle{
			X:      x,
			Y:      y,
			Width:  w,
			Height: h,
		},
		rl.Vector2Scale(rl.Vector2One(), 2),
		0,
		rl.RayWhite,
	)

	rl.DrawTexturePro(
		*tx,
		s.singleton.sheet.Frame(s.getFrameIndex()),
		s.singleton.sheet.Frame(s.getFrameIndex()),
		rl.Vector2Scale(rl.Vector2One(), 2),
		0,
		rl.RayWhite,
	)

	rl.DrawTexture(
		*tx,
		0, 0,
		rl.NewColor(255, 255, 255, 128),
	)

	rl.DrawText(
		fmt.Sprintf("(%0.2f, %0.2f)\nangle(%0.2f)\ndirection #%d.%d (%d)",
			s.singleton.rotation.X, s.singleton.rotation.Y,
			Vector2ToDegrees(s.singleton.rotation),
			s.getCurrentDirection(), s.singleton.animation.currentFrame,
			s.getFrameIndex(),
		),
		10, 450,
		20,
		rl.White,
	)
}
