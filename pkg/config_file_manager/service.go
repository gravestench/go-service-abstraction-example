package config_file_manager

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

func New(name string, filePaths ...string) *Service {
	if len(filePaths) < 1 {
		return New(name, defaultConfigDir, defaultConfigFileName)
	}

	if len(filePaths) < 2 {
		return New(name, "", filePaths[0])
	}

	s := &Service{
		name:        name,
		cfgDirPath:  filePaths[0],
		cfgFilePath: filePaths[1],
	}

	return s
}

// each service will maintain its own config
type perServiceConfig = map[string]map[string]string

type Service struct {
	name         string
	store        perServiceConfig
	previous     string
	log          abstract.Logger
	mux          sync.Mutex
	cfgDirPath   string
	cfgFilePath  string
	lastModified time.Time
}

func (s *Service) SetLogger(l abstract.Logger) {
	s.log = l
}

func (s *Service) Logger() *zerolog.Logger {
	return s.log.Logger()
}

func (s *Service) Init(app abstract.ServiceManager) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.store = make(perServiceConfig)

	s.LoadConfig()
	go s.loopSaveConfig()
}

func (s *Service) Name() string {
	if s.name == "" {
		return "Config File Manager"
	}

	return fmt.Sprintf("Config File Manager (%s)", s.name)
}

func (s *Service) DirPath() string {
	if s.cfgDirPath == "" {
		s.cfgDirPath = defaultConfigDir
	}

	s.cfgDirPath, _ = expandTilde(s.cfgDirPath)

	return s.cfgDirPath
}

func (s *Service) FilePath() string {
	return s.cfgFilePath
}
