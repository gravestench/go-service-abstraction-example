package user

import (
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

func (s *Service) populateServices(dependencies *[]interface{}) {
	for {
		if s.log != nil && s.cfg != nil {
			break
		}

		for _, candidate := range *dependencies {
			if d, ok := candidate.(abstract.Logger); ok {
				s.log = d
			}

			if d, ok := candidate.(abstract.ConfigurationManager); ok {
				s.cfg = d
			}
		}

		time.Sleep(time.Millisecond * 10)
	}
}
