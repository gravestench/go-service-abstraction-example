package sprite_test

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
	"github.com/gravestench/go-service-abstraction-example/pkg/services/mode"
)

type Scene struct {
	input       abstract.InputService
	render      abstract.RenderService
	mode        abstract.ModeService
	texture     abstract.TextureService
	spriteSheet abstract.SpriteSheetService

	singleton struct {
		sheet     abstract.SpriteSheet
		position  rl.Vector2
		rotation  rl.Vector2
		velocity  rl.Vector2
		animation struct {
			currentFrame  int
			frameDuration time.Duration
			frameTimer    time.Duration
		}
	}

	active bool
}

func (s *Scene) Mode() abstract.Mode {
	return mode.ModeGame
}

func (s *Scene) SetMode(m abstract.Mode) {
	s.active = m == s.Mode()
}

func (s *Scene) Init(possibleDependencies *[]interface{}) {
	s.populateDependencies(possibleDependencies)
	s.singleton.position.X, s.singleton.position.Y = 400, 300
	s.prepareSprite()
	return // noop
}

func (s *Scene) Name() string {
	return "Sprite Test Scene"
}

func (s *Scene) IsActive() bool {
	return s.active && s.singleton.sheet != nil && s.singleton.sheet.Texture() != nil
}

func (s *Scene) Ready() bool {
	return s.input != nil
}
