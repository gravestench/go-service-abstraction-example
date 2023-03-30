package sprite_test

import (
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	spriteKey  = "example"
	spritePath = "./pkg/scenes/sprite_test/data/graphics/sprites/iso_char.png"
)

func (s *Scene) prepareSprite() {
	s.spriteSheet.LoadImage(spriteKey, spritePath, 5, 8)

	// Set up the animation
	s.singleton.animation.currentFrame = 0
	s.singleton.animation.frameDuration = time.Second / 5 // Each frame will be displayed for 0.1 seconds
	s.singleton.animation.frameTimer = 0.0
}

func (s *Scene) getCurrentDirection() int {
	angle := float64(Vector2ToDegrees(s.singleton.rotation) + 0.1)
	_, numDirections := s.singleton.sheet.FrameDimensions()
	return quantizeDegrees(angle, numDirections)
}

func (s *Scene) getFrameOffset() int {
	direction := s.getCurrentDirection()
	framesPerDirection, _ := s.singleton.sheet.FrameDimensions()
	return direction * framesPerDirection
}

func (s *Scene) getFrameIndex() int {
	offset := s.getFrameOffset()
	return offset + s.singleton.animation.currentFrame
}

func (s *Scene) updateSpritePosition(t time.Duration) {
	s.singleton.position = rl.Vector2Add(s.singleton.position, s.singleton.velocity)

	ww, wh := s.render.WindowWidth(), s.render.WindowHeight()

	if ww < 1 || wh < 1 {
		return
	}

	s.singleton.position.X, s.singleton.position.Y = WrapCoordinates(s.singleton.position.X, s.singleton.position.Y, Rectangle{
		Width:  float32(ww),
		Height: float32(wh),
	})
}

func (s *Scene) updateSpriteRotation(t time.Duration) {
	if _, numDirections := s.singleton.sheet.FrameDimensions(); numDirections < 1 {
		return
	}

	if math.Abs(float64(s.singleton.velocity.X)) < 0.75 && math.Abs(float64(s.singleton.velocity.Y)) < 0.75 {
		return
	}

	// Set the rotation vector based on the velocity
	s.singleton.rotation = rl.Vector2{X: s.singleton.velocity.Y, Y: -s.singleton.velocity.X}
}

func (s *Scene) updateSpriteVelocity(t time.Duration) {
	const (
		speed    = 1
		maxSpeed = 4
	)

	if s.input.IsKeyDown(rl.KeyLeft) {
		s.singleton.velocity.X -= speed
	} else if s.input.IsKeyDown(rl.KeyRight) {
		s.singleton.velocity.X += speed
	} else {
		s.singleton.velocity.X -= s.singleton.velocity.X / 10
	}

	if s.input.IsKeyDown(rl.KeyUp) {
		s.singleton.velocity.Y -= speed
	} else if s.input.IsKeyDown(rl.KeyDown) {
		s.singleton.velocity.Y += speed
	} else {
		s.singleton.velocity.Y -= s.singleton.velocity.Y / 10
	}

	if s.singleton.velocity.X > maxSpeed {
		s.singleton.velocity.X = maxSpeed
	} else if s.singleton.velocity.X < -maxSpeed {
		s.singleton.velocity.X = -maxSpeed
	}

	if s.singleton.velocity.Y > maxSpeed {
		s.singleton.velocity.Y = maxSpeed
	} else if s.singleton.velocity.Y < -maxSpeed {
		s.singleton.velocity.Y = -maxSpeed
	}
}

func (s *Scene) updateSpriteFrame(t time.Duration) {
	if math.Abs(float64(s.singleton.velocity.X)) < 0.75 && math.Abs(float64(s.singleton.velocity.Y)) < 0.75 {
		s.singleton.animation.currentFrame = 0
		return
	}

	s.singleton.animation.frameTimer += t
	if s.singleton.animation.frameTimer >= s.singleton.animation.frameDuration {
		s.singleton.animation.frameTimer = 0
		s.singleton.animation.currentFrame++
		framesPerDirection, _ := s.singleton.sheet.FrameDimensions()
		if s.singleton.animation.currentFrame >= framesPerDirection {
			s.singleton.animation.currentFrame = 0
		}
	}
}

func getMousePosition() (x, y float64) {
	if !rl.IsWindowReady() {
		return 0, 0
	}

	rl.HideCursor()

	mousePos := rl.GetMousePosition()
	//windowPos := rl.GetWindowPosition()

	// Calculate the mouse position relative to the window
	mouseX := mousePos.X /*- windowPos.X*/
	mouseY := mousePos.Y /*- windowPos.Y*/

	return float64(mouseX), float64(mouseY)
}

func Vector2ToDegrees(vec rl.Vector2) float32 {
	return rl.Vector2Angle(rl.Vector2Zero(), vec) * 180 / math.Pi
}

func DegreesToVector2(degrees float32) rl.Vector2 {
	// Convert the angle from degrees to radians
	radians := degrees * math.Pi / 180

	// Calculate the x and y components of the vector
	x := float32(math.Sin(float64(radians)))
	y := -float32(math.Cos(float64(radians)))

	return rl.Vector2{X: x, Y: y}
}

func quantizeDegrees(degrees float64, quantization int) int {
	for degrees > 360 {
		degrees -= 360
	}

	// Convert the angle to a range of 0 to 360 degrees
	degrees = math.Mod(math.Abs(degrees), 360)

	// Calculate the size of each quantization interval
	interval := 360.0 / float64(quantization)

	// Quantize the angle to the nearest interval
	quantized := int(math.Floor((degrees + interval/2) / interval))

	return quantized
}
