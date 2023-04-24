package abstract

import (
	"time"
)

// ConfigurationManager is a super simple/naive data store
// that keeps track of grouped key=value pairs.
// the underlying storage looks like map[string]map[string]string
// but can really be anything as long it's conceptually configured like this
type ConfigurationManager interface {
	DirPath() string
	FilePath() string
	LoadConfig()
	SaveConfig()
	GetConfigGroup(groupName string) ConfigGroup
	GetKeys() []string
	Get(group, key string) string
	Set(group, key string, value any)
	Default(group, key string, val any)
	Delete(...string)
	LastModified() time.Time
}

// ConfigGroup is an interface for managing values
// inside a group being managed within the config manager
type ConfigGroup interface {
	GetKeys() []string
	Get(key string) string
	Set(key string, value any)
	Default(key string, val any)
	Delete(string)
	String() string
}
