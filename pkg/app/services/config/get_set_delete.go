package config

import (
	"fmt"
)

func (s *Service) Get(group, key string) (result string) {
	if s.store == nil {
		s.store = make(map[string]map[string]string)
	}

	if data, found := s.store[group]; found {
		result = data[key]
	}

	return result
}

func (s *Service) Set(group, key string, value any) {
	if _, found := s.store[group]; !found {
		s.store[group] = make(map[string]string)
	}

	s.store[group][key] = fmt.Sprintf("%v", value)
}

func (s *Service) Delete(args ...string) {
	switch len(args) {
	case 1:
		s.Msgf("deleting config for group '%s'", args[0])
		delete(s.store, args[0])
	case 2:
		s.Msgf("deleting ['%s']['%s']", args[0], args[1])
		if data, found := s.store[args[0]]; found {
			delete(data, args[1])
		}
	}
}
