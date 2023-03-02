package abstract

type ConfigurationManager interface {
	LoadConfig()
	SaveConfig()
	Get(group, key string) string
	Set(group, key string, value any)
	Delete(...string)
}
