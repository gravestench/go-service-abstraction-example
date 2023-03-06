package abstract

// ConfigurationManager is a super simple/naive data store
// that keeps track of grouped key=value pairs.
// the underlying storage looks like map[string]map[string]string
// but can really be anything as long it's conceptually configured like this
type ConfigurationManager interface {
	LoadConfig()
	SaveConfig()
	Group(name string) ValueControl
	GetKeys(group string) []string
	Get(group, key string) string
	Set(group, key string, value any)
	Touch(group, key string)
	Default(group, key string, val any)
	Delete(...string)
}

type ValueControl interface {
	Get(key string) string
	Set(key string, value any)
	Touch(key string)
	Default(key string, val any)
	Delete(string)
}
