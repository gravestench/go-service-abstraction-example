package spritesheet

import (
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

func (s *Service) populateDependencies(others *[]interface{}) {
	for {
		if s.texture != nil && s.render != nil {
			break
		}

		for _, other := range *others {
			if t, ok := other.(abstract.TextureService); ok {
				s.texture = t
			}

			if t, ok := other.(abstract.RenderService); ok {
				s.render = t
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}
