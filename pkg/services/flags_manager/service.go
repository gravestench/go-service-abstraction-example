package flags_manager

import (
	"flag"
	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
	"os"
	"regexp"
	"strings"
	"time"
)

// Service is responsible for iterating over other services and
// filtering out the CLI args which dont apply to that particular service.
// this addresses an issue with the golang flag implementation that
// prevents flags from being defined per-module. If you try to parse
// flags which are not supposed to be parsed, it errors early and doesnt
// parse all of the flags.
type Service struct {
	log  abstract.Logger
	args []string

	parsed map[string]struct{}
}

func (s *Service) Init(possibleDependencies *[]interface{}) {
	s.args = make([]string, len(os.Args))
	copy(s.args, os.Args)
	s.parsed = make(map[string]struct{})
	s.populateDependencies(possibleDependencies)
	go s.applyFlags(possibleDependencies)
}

func (s *Service) Name() string {
	return "Flag Parser"
}

func (s *Service) applyFlags(services *[]interface{}) {
	for {
		for _, candidate := range *services {
			if flagUser, ok := candidate.(abstract.FlagService); ok {
				if _, found := s.parsed[flagUser.Name()]; found {
					continue
				}

				args := s.getArgs()
				args = s.filterArgsForService(flagUser, args)

				s.log.Info().Msgf("[%s] parsing CLI flags for '%s' service", s.Name(), flagUser.Name())
				flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
				err := flagUser.Parse(flagSet, args)
				if err != nil {
					continue
				}

				if _, found := s.parsed[flagUser.Name()]; !found {
					s.parsed[flagUser.Name()] = struct{}{}
				}
			}
		}

		time.Sleep(time.Second)
	}
}

func (s *Service) getArgs() []string {
	args := make([]string, len(s.args))
	copy(args, s.args)

	return args
}

func (s *Service) filterArgsForService(fs abstract.FlagService, args []string) []string {
	matchArg := `--?[a-zA-Z0-9-]+( [^-][^ ]+)`
	m := regexp.MustCompile(matchArg)
	args = m.FindAllString(strings.Join(args, " "), -1)
	filtered := make([]string, 0)

	for _, arg := range args {
		if containsOneOfTheFlags(arg, fs.Flags()) {
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
