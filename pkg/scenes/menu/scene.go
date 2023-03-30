package menu

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
	"github.com/gravestench/go-service-abstraction-example/pkg/services/mode"
	"time"
)

type Scene struct {
	input  abstract.InputService
	render abstract.RenderService
	mode   abstract.ModeService
	active bool
}

func (s *Scene) Mode() abstract.Mode {
	return mode.ModeMenu
}

func (s *Scene) SetMode(m abstract.Mode) {
	s.active = m == s.Mode()
}

func (s *Scene) populateDependencies(others *[]interface{}) {
	for {
		if s.input != nil && s.render != nil {
			break
		}

		for _, other := range *others {
			if input, ok := other.(abstract.InputService); ok {
				s.input = input
			}

			if r, ok := other.(abstract.RenderService); ok {
				s.render = r
			}

			if m, ok := other.(abstract.ModeService); ok {
				s.mode = m
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}

func (s *Scene) Init(possibleDependencies *[]interface{}) {
	s.active = false
	s.populateDependencies(possibleDependencies)
}

func (s *Scene) Name() string {
	return "Menu"
}

func (s *Scene) IsActive() bool {
	return s.active && s.render != nil && s.input != nil && s.mode != nil
}

func (s *Scene) Render() {
	rl.DrawRectangleLines(
		100, 100,
		int32(-200+s.render.WindowWidth()),
		int32(-200+s.render.WindowHeight()),
		rl.NewColor(255, 0, 255, 255),
	)
}

func (s *Scene) Ready() bool {
	return s.input != nil
}

func (s *Scene) Update(duration time.Duration) {

}
