package mode

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

const (
	ModeGame abstract.Mode = iota - 1
	ModeMenu
)

type Service struct {
	input       abstract.InputService
	currentMode abstract.Mode
	modals      []abstract.Modal
}

func (s *Service) Update(_ time.Duration) {
	if s.input.IsKeyPressed(rl.KeyEscape) {
		switch s.currentMode {
		case ModeGame:
			s.currentMode = ModeMenu
		case ModeMenu:
			s.currentMode = ModeGame
		}
	}

	s.SetMode(s.currentMode)
}

func (s *Service) Ready() bool {
	return s.input != nil
}

func (s *Service) Init(possibleDependencies *[]interface{}) {
	s.currentMode = ModeGame
	s.populateDependencies(possibleDependencies)
	go s.findModals(possibleDependencies)
}

func (s *Service) Name() string {
	return "Mode Select"
}

func (s *Service) Modals() []abstract.Modal {
	return s.modals
}

func (s *Service) SetMode(mode abstract.Mode) {
	s.currentMode = mode
	for _, m := range s.modals {
		m.SetMode(s.currentMode)
	}
}

func (s *Service) findModals(services *[]interface{}) {
	for {
		for _, service := range *services {
			if modal, ok := service.(abstract.Modal); ok {
				s.modals = append(s.modals, modal)
			}
		}

		time.Sleep(time.Second)
	}
}

func (s *Service) populateDependencies(others *[]interface{}) {
	for {
		if s.input != nil {
			break
		}

		for _, other := range *others {
			if input, ok := other.(abstract.InputService); ok {
				s.input = input
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}
