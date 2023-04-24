package logging

import (
	"flag"
	"time"

	"github.com/rs/zerolog"
)

const (
	flagLogLevel            = "log-level"
	flagLogLevelDefault     = int(zerolog.InfoLevel)
	flagLogLevelDescription = "logging verbosity"
)

func (l *Logger) Flags() []string {
	return []string{
		flagLogLevel,
	}
}

func (l *Logger) Parse(fs *flag.FlagSet, args []string) error {
	fs.IntVar(&l.flags.level, flagLogLevel, flagLogLevelDefault, flagLogLevelDescription)

	l.flags.parsed = true

	return fs.Parse(args)
}

func (l *Logger) dwellUntilFlagsParsed() {
	for {
		if l.flags.parsed {
			break
		}

		time.Sleep(time.Millisecond * 10)
	}
}
