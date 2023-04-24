package config_file_manager

import (
	"fmt"
	"sort"
	"time"
)

func (s *Service) LastModified() time.Time {
	return s.lastModified
}

func (s *Service) GetKeys() (result []string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.store == nil {
		s.store = make(map[string]map[string]string)
	}

	for key := range s.store {
		result = append(result, key)
	}

	sort.Strings(result)

	return result
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

	if s.store == nil {
		s.store = make(map[string]map[string]string)
	}

	if data, found := s.store[group]; found {
		result = data[key]
	}

	s.mux.Unlock()

	return result
}

func (s *Service) Set(group, key string, value any) {
	s.mux.Lock()

	if _, found := s.store[group]; !found {
		s.store[group] = make(map[string]string)
	}

	s.store[group][key] = fmt.Sprintf("%v", value)

	s.mux.Unlock()

	s.lastModified = time.Now()
}

func (s *Service) Delete(args ...string) {
	s.mux.Lock()

	switch len(args) {
	case 1:
		s.log.Info().Msgf("deleting config for group '%s'", args[0])
		delete(s.store, args[0])
	case 2:
		s.log.Info().Msgf("deleting ['%s']['%s']", args[0], args[1])
		if data, found := s.store[args[0]]; found {
			delete(data, args[1])
		}
	}

	s.mux.Unlock()

	s.lastModified = time.Now()
}
