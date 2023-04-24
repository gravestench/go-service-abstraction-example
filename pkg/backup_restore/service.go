package backup_restore

import (
	"time"

	"github.com/rs/zerolog"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

type Service struct {
	log abstract.Logger
	cfg abstract.ConfigurationManager

	backupers map[string]abstract.CanBackUp
	restorers map[string]abstract.CanRestore
}

func (s *Service) SetLogger(l abstract.Logger) {
	s.log = l
}

func (s *Service) Logger() *zerolog.Logger {
	return s.log.Logger()
}

func (s *Service) Init(manager abstract.ServiceManager) {
	s.backupers = make(map[string]abstract.CanBackUp)
	s.restorers = make(map[string]abstract.CanRestore)

	go s.lookForBackupers(manager.Services())
	go s.lookForRestorers(manager.Services())
}

func (s *Service) Name() string {
	return "Backup & Restore"
}

func (s *Service) lookForBackupers(i *[]interface{}) {
	for {
		for _, candidate := range *i {
			if b, ok := candidate.(abstract.CanBackUp); ok {
				s.backupers[b.BackupKey()] = b
			}
		}

		time.Sleep(time.Second)
	}
}

func (s *Service) lookForRestorers(i *[]interface{}) {
	for {
		for _, candidate := range *i {
			if r, ok := candidate.(abstract.CanRestore); ok {
				s.restorers[r.BackupKey()] = r
			}
		}

		time.Sleep(time.Second)
	}
}
