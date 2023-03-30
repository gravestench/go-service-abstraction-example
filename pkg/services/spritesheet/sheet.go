package spritesheet

import rl "github.com/gen2brain/raylib-go/raylib"

type sheet struct {
	rows, columns int
	sequences     map[string][]int
	frames        []rl.Rectangle
	texture       *rl.Texture2D
}

func (s *sheet) NumFrames() int {
	return s.rows * s.columns
}

func (s *sheet) NumSequences() int {
	return len(s.sequences)
}

func (s *sheet) Sequences() map[string][]int {
	return s.sequences
}

func (s *sheet) Texture() *rl.Texture2D {
	return s.texture
}

func (s *sheet) Frame(i int) rl.Rectangle {
	return s.frames[i]
}

func (s *sheet) SetSequence(name string, frames []int) {
	s.sequences[name] = frames
}

func (s *sheet) FrameDimensions() (w, h int) {
	if s.NumFrames() == 0 {
		return 0, 0
	}

	f := s.Frame(0)

	return int(f.X), int(f.Y)
}
