package renderer

import (
	"time"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

func (s *Service) lookForRenderers(possibleRenderer *[]interface{}) {
	for {
		for _, other := range *possibleRenderer {
			if r, ok := other.(abstract.Renderer); ok {
				if _, found := s.lookup[r.Name()]; !found {
					s.lookup[r.Name()] = struct{}{}
					s.renderables = append(s.renderables, r)
				}
			}
		}

		time.Sleep(time.Second)
	}
}

func (s *Service) renderForever() {
	mainthread.Call(func() {
		for !s.windowShouldClose {
			s.render()
		}

		rl.CloseWindow()
	})
}

func (s *Service) render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for _, r := range s.renderables {
		if r.IsActive() {
			r.Render()
		}
	}

	rl.EndDrawing()
}

func (s *Service) WindowWidth() int {
	return rl.GetScreenWidth()
}

func (s *Service) WindowHeight() int {
	return rl.GetScreenHeight()
}

func (s *Service) Close() {
	s.windowShouldClose = true
}

func (s *Service) Initialized() bool {
	return s.initialized
}
