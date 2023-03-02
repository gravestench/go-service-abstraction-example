package logging

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

var _ abstract.Logger = &Service{}

func New() *Service {
	// Create a console writer with a custom print format
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
	}

	s := &Service{
		l: zerolog.New(consoleWriter).With().Timestamp().Logger(),
	}

	return s
}

type Service struct {
	l zerolog.Logger
}

func (s Service) Init(_ *[]interface{}) {
	// noop
}

func (s Service) Name() string {
	return "Logging"
}

func (s *Service) Info() (m abstract.Messager) {
	return s.l.Info()
}

func (s *Service) Warn() (m abstract.Messager) {
	return s.l.Warn()
}

func (s *Service) Error() (m abstract.Messager) {
	return s.l.Error()
}

func (s *Service) Fatal() (m abstract.Messager) {
	return s.l.Fatal()
}
