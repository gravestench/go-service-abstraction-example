package flags_distributor

import (
	"flag"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
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
	mux    sync.Mutex
}

func (s *Service) SetLogger(l abstract.Logger) {
	s.log = l
}

func (s *Service) Logger() *zerolog.Logger {
	return s.log.Logger()
}

func (s *Service) Init(app abstract.ServiceManager) {
	s.args = make([]string, len(os.Args))
	copy(s.args, os.Args)

	s.mux.Lock()
	s.parsed = make(map[string]struct{})
	s.mux.Unlock()

	go s.applyFlags(app.Services())
}

func (s *Service) Name() string {
	return "Flag Parser"
}

func (s *Service) applyFlags(services *[]interface{}) {
	for {
		for _, candidate := range *services {
			if svc, ok := candidate.(abstract.ServiceThatUsesFlags); ok {
				s.mux.Lock()
				if _, found := s.parsed[svc.Name()]; found {
					s.mux.Unlock()
					continue
				}

				args := s.getArgs()
				args = s.filterArgsForService(svc, args)

				s.log.Info().Msgf("parsing CLI flags for '%s' service", svc.Name())

				flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

				s.parsed[svc.Name()] = struct{}{}

				if err := svc.Parse(flagSet, args); err != nil {
					s.mux.Unlock()
					continue
				}

				s.mux.Unlock()
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

func (s *Service) filterArgsForService(fs abstract.ServiceThatUsesFlags, args []string) []string {
	matchArg := `--?[a-zA-Z0-9-]+( [^-][^ ]*)?`
	m := regexp.MustCompile(matchArg)
	args = m.FindAllString(strings.Join(args, " "), -1)
	filtered := make([]string, 0)

	for _, arg := range args {
		flagsTheServicesLooksFor := fs.Flags()
		flagsTheServicesLooksFor = append(flagsTheServicesLooksFor, "--help")
		if containsOneOfTheFlags(arg, flagsTheServicesLooksFor) {
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
