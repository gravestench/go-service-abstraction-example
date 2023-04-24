package logging

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

var _ abstract.Logger = &Logger{}

func New(name string) *Logger {
	// Create a console writer with a custom print format
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return fmt.Sprintf("| %5s | %-20s |", strings.ToUpper(fmt.Sprintf("%s", i)), name)
		},
		NoColor: false,
	}

	l := zerolog.New(consoleWriter).With().Timestamp().Logger()
	s := &Logger{
		l: &l,
	}

	return s
}

type Logger struct {
	l     *zerolog.Logger
	flags struct {
		parsed bool
		level  int
	}
}

func (l *Logger) Init() {
	l.SetLevel(l.flags.level)
}

func (l *Logger) Logger() *zerolog.Logger {
	return l.l
}

func (l *Logger) SetLevel(level int) abstract.Logger {
	l.l.Info().Msgf("setting log level: %s", zerolog.Level(level).String())
	newLogger := l.l.Level(zerolog.Level(level))
	l.l = &newLogger
	return l
}

func (l *Logger) Trace() (m abstract.Messager) {
	return l.l.Trace()
}

func (l *Logger) Debug() (m abstract.Messager) {
	return l.l.Debug()
}

func (l *Logger) Info() (m abstract.Messager) {
	return l.l.Info()
}

func (l *Logger) Warn() (m abstract.Messager) {
	return l.l.Warn()
}

func (l *Logger) Error() (m abstract.Messager) {
	return l.l.Error()
}

func (l *Logger) Fatal() (m abstract.Messager) {
	return l.l.Fatal()
}
