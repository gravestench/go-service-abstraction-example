package update

import (
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

type Service struct {
	updaters   map[string]abstract.Updater
	lastUpdate time.Time
	scalar     float64
}

func (s *Service) Init(possibleDependencies *[]interface{}) {
	s.updaters = make(map[string]abstract.Updater)
	s.scalar = 1.0
	s.lastUpdate = time.Now()
	go s.updateForever()
	go s.lookForUpdaters(possibleDependencies)

}

func (s *Service) Name() string {
	return "Update Service"
}

func (s *Service) updateForever() {
	for {
		s.Update()
		time.Sleep(time.Second / 60)
	}
}

func (s *Service) Update() {
	for _, updater := range s.updaters {
		if !updater.Ready() {
			continue
		}

		timeSinceLastFrame := time.Since(s.lastUpdate)
		s.lastUpdate = time.Now()

		updater.Update(time.Duration(float64(timeSinceLastFrame) * s.scalar))
	}
}

func (s *Service) lookForUpdaters(possibleUpdater *[]interface{}) {
	for {
		for _, other := range *possibleUpdater {
			if u, ok := other.(abstract.Updater); ok {
				if _, found := s.updaters[u.Name()]; !found {
					s.updaters[u.Name()] = u
				}
			}
		}

		time.Sleep(time.Second)
	}
}
