package config_file_manager

import (
	"encoding/json"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

func (s *Service) GetConfigGroup(name string) abstract.ConfigGroup {
	return &tx{
		fnGetKeys: func() []string { return s.GetKeys() },
		fnGet:     func(key string) string { return s.Get(name, key) },
		fnSet:     func(key string, val any) { s.Set(name, key, val) },
		fnDefault: func(key string, val any) { s.Default(name, key, val) },
		fnDelete:  func(key string) { s.Delete(name, key) },
		fnString: func() string {
			bytes, _ := json.Marshal(s.store[name])
			return string(bytes)
		},
	}
}

type tx struct {
	fnGetKeys func() []string
	fnGet     func(key string) string
	fnSet     func(key string, val any)
	fnDefault func(key string, val any)
	fnDelete  func(key string)
	fnString  func() string
}

func (t tx) GetKeys() []string           { return t.fnGetKeys() }
func (t tx) Get(key string) string       { return t.fnGet(key) }
func (t tx) Set(key string, val any)     { t.fnSet(key, val) }
func (t tx) Default(key string, val any) { t.fnDefault(key, val) }
func (t tx) Delete(key string)           { t.fnDelete(key) }
func (t tx) String() string              { return t.fnString() }
