package config

import (
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

func (s *Service) Group(name string) abstract.ValueControl {
	return &tx{
		fnGet:     func(key string) string { return s.Get(name, key) },
		fnSet:     func(key string, val any) { s.Set(name, key, val) },
		fnTouch:   func(key string) { s.Touch(name, key) },
		fnDefault: func(key string, val any) { s.Default(name, key, val) },
	}
}

type tx struct {
	fnGet     func(key string) string
	fnSet     func(key string, val any)
	fnTouch   func(key string)
	fnDefault func(key string, val any)
}

func (t tx) Get(key string) string       { return t.fnGet(key) }
func (t tx) Set(key string, val any)     { t.fnSet(key, val) }
func (t tx) Touch(key string)            { t.fnTouch(key) }
func (t tx) Default(key string, val any) { t.fnDefault(key, val) }
