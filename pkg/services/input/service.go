package input

import rl "github.com/gen2brain/raylib-go/raylib"

type Service struct {
}

func (s *Service) Init(possibleDependencies *[]interface{}) {

}

func (s *Service) Name() string {
	return "Input Service"
}

func (s *Service) IsKeyDown(keyCodes ...int) bool {
	if len(keyCodes) < 1 {
		return false
	}

	for _, code := range keyCodes {
		if !rl.IsKeyDown(int32(code)) {
			return false
		}
	}

	return true
}

func (s *Service) IsKeyPressed(keyCodes ...int) bool {
	if len(keyCodes) < 1 {
		return false
	}

	for _, code := range keyCodes {
		if !rl.IsKeyPressed(int32(code)) {
			return false
		}
	}

	return true
}

func (s *Service) IsKeyUp(keyCodes ...int) bool {
	if len(keyCodes) < 1 {
		return false
	}

	for _, code := range keyCodes {
		if !rl.IsKeyUp(int32(code)) {
			return false
		}
	}

	return true
}

func (s *Service) IsKeyReleased(keyCodes ...int) bool {
	if len(keyCodes) < 1 {
		return false
	}

	for _, code := range keyCodes {
		if !rl.IsKeyReleased(int32(code)) {
			return false
		}
	}

	return true
}

func (s *Service) IsMouseButtonDown(buttonCodes ...int) bool {
	if len(buttonCodes) < 1 {
		return false
	}

	for _, code := range buttonCodes {
		if !rl.IsMouseButtonDown(int32(code)) {
			return false
		}
	}

	return true
}
