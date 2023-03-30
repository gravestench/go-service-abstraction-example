package renderer

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

type Service struct {
	*config
	lookup            map[string]struct{}
	renderables       []abstract.Renderer
	initialized       bool
	windowShouldClose bool
}

func (s *Service) Init(possibleDependencies *[]interface{}) {
	s.lookup = make(map[string]struct{})
	s.renderables = make([]abstract.Renderer, 0)
	s.config = s.config.Default()

	{
		w := s.config.window
		rl.SetTraceLog(rl.LogNone)
		//rl.SetConfigFlags(rl.FlagFullscreenMode | rl.FlagWindowUndecorated | rl.FlagWindowMaximized)
		rl.InitWindow(800, 600, w.Title)
	}

	rl.SetTargetFPS(s.config.fps)

	go s.lookForRenderers(possibleDependencies)
	go s.renderForever()

	s.initialized = true
}

func (s *Service) Name() string {
	return "Renderer"
}
