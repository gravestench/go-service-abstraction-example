package texture

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
	"time"
)

type Service struct {
	render abstract.RenderService
}

func (s *Service) Init(possibleDependencies *[]interface{}) {
	s.populateDependencies(possibleDependencies)
}

func (s *Service) Name() string {
	return "Texture Manager"
}

func (s *Service) LoadImageToTexture(filepath string) rl.Texture2D {
	for !s.render.Initialized() {
		time.Sleep(time.Second * 10)
	}

	return rl.LoadTextureFromImage(rl.LoadImage(filepath))
}

func (s *Service) populateDependencies(others *[]interface{}) {
	for {
		if s.render != nil {
			break
		}

		for _, other := range *others {
			if r, ok := other.(abstract.RenderService); ok {
				s.render = r
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}
