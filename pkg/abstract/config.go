package abstract

type HasConfig interface {
	Config() Config
}

type Config interface {
	Default() Config
	ToBytes() ([]byte, error)
	FromBytes([]byte) error
}

type ConfigManager interface {
	Save() error
	Load() error
}
