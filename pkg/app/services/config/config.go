package config

import (
	"fmt"
	"sort"
)

func (s *Service) GetKeys(group string) (result []string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.store == nil {
		s.store = make(map[string]map[string]string)
	}

	for key := range s.store[group] {
		result = append(result, key)
	}

	sort.Strings(result)

	return result
}

func (s *Service) Touch(group, key string) {
	s.Set(group, key, s.Get(group, key))
}

func (s *Service) Default(group, key string, def any) {
	v := s.Get(group, key)

	if v != "" {
		return
	}

	s.Set(group, key, def)
}

func (s *Service) Get(group, key string) (result string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.store == nil {
		s.store = make(map[string]map[string]string)
	}

	if data, found := s.store[group]; found {
		result = data[key]
	}

	return result
}

func (s *Service) Set(group, key string, value any) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, found := s.store[group]; !found {
		s.store[group] = make(map[string]string)
	}

	s.store[group][key] = fmt.Sprintf("%v", value)
}

func (s *Service) Delete(args ...string) {
	s.mux.Lock()
	defer s.mux.Unlock()

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
