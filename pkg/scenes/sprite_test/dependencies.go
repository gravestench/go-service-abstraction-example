package sprite_test

import (
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

func (s *Scene) populateDependencies(others *[]interface{}) {
	for {
		if s.input != nil && s.render != nil && s.mode != nil && s.texture != nil && s.spriteSheet != nil {
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

			if t, ok := other.(abstract.TextureService); ok {
				s.texture = t
			}

			if t, ok := other.(abstract.SpriteSheetService); ok {
				s.spriteSheet = t
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}
