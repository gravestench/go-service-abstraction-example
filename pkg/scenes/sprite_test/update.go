package sprite_test

import "time"

func (s *Scene) Update(t time.Duration) {
	if s.input == nil {
		return
	}

	if s.singleton.sheet == nil {
		return
	}

	s.updateSprite(t)
}

func (s *Scene) updateSprite(t time.Duration) {
	//s.updateSpritePosition(t)
	s.updateSpriteRotation(t)
	s.updateSpriteVelocity(t)
	s.updateSpriteFrame(t)
}
