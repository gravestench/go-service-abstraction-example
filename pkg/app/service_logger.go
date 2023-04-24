package app

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/rs/zerolog"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/logging"
)

func (a *App) setupServiceLogger(s abstract.Service) {
	// set up the logger so that it prints with its name prefixed
	if l, ok := s.(abstract.HasLogger); ok {
		l.SetLogger(logging.New(s.Name()))

		if ll := a.getLogLevelArg(); len(ll) > 1 {
			level, _ := strconv.ParseInt(ll[1], 10, 64)
			l.Logger().Level(zerolog.Level(level))
		}
	}
}

func (a *App) getLogLevelArg() []string {
	matchArg := `--?[a-zA-Z0-9-]+( [^-]?[^ ]*)?`
	m := regexp.MustCompile(matchArg)
	args := m.FindAllString(strings.Join(os.Args, " "), -1)
	filtered := make([]string, 0)

	for _, arg := range args {
		if containsOneOfTheFlags(arg, []string{"--log-level"}) {
			filtered = append(filtered, strings.Split(arg, " ")...)
		}
	}

	return filtered
}

func containsOneOfTheFlags(arg string, argsNeededByService []string) bool {
	for _, n := range argsNeededByService {
		if strings.Contains(arg, n) {
			return true
		}
	}

	return false
}
