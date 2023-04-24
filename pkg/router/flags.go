package router

import (
	"flag"
	"time"

	"github.com/rs/zerolog"
)

const (
	//flagSlug            = "router-slug"
	//flagSlugDefault     = "api"
	//flagSlugDescription = "root router slug"

	flagDisableSessions            = "session-disable"
	flagDisableSessionsDefault     = false
	flagDisableSessionsDescription = "disable session middleware"

	flagLogLevel            = "log-level"
	flagLogLevelDefault     = int(zerolog.InfoLevel)
	flagLogLevelDescription = "logging verbosity"
)

func (s *Service) Flags() []string {
	return []string{
		flagLogLevel,
		flagDisableSessions,
	}
}

func (s *Service) Parse(fs *flag.FlagSet, args []string) error {
	//fs.StringVar(&s.flags.slug, flagSlug, flagSlugDefault, flagSlugDescription)
	fs.BoolVar(&s.flags.disableSessions, flagDisableSessions, flagDisableSessionsDefault, flagDisableSessionsDescription)
	fs.IntVar(&s.flags.logLevel, flagLogLevel, flagLogLevelDefault, flagLogLevelDescription)

	s.flags.parsed = true

	return fs.Parse(args)
}

func (s *Service) dwellUntilFlagsParsed() {
	for {
		if s.flags.parsed {
			break
		}

		time.Sleep(time.Millisecond * 10)
	}
}
