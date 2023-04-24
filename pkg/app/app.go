package app

import (
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/logging"
)

func New() *App {
	a := &App{}
	a.init()
	return a
}

type App struct {
	quit     chan os.Signal
	services []interface{}
	log      abstract.Logger
}

func (a *App) init() {

	if a.services != nil {
		return
	}

	a.log = logging.New("APP")

	rand.Seed(time.Now().UnixNano())

	a.quit = make(chan os.Signal, 1)
	signal.Notify(a.quit, os.Interrupt)

	a.services = make([]interface{}, 0)

	go func() { // prevent deadlock panic
		for {
			time.Sleep(time.Second)
		}
	}()
}

func (a *App) AddServices(ss ...abstract.Service) {
	for _, s := range ss {
		a.AddService(s)
	}
}

func (a *App) AddService(s abstract.Service) {
	a.init()
	a.addService(s)
}

func (a *App) Services() *[]interface{} {
	return &a.services
}

func (a *App) addService(s abstract.Service) {
	a.log.Info().Msgf("adding '%s' service", s.Name())

	a.services = append(a.services, s)

	// non-blocking call to resolve dependencies and then init
	go a.resolveDependenciesAndInit(s)
}

func (a *App) resolveDependenciesAndInit(s abstract.Service) {
	a.setupServiceLogger(s)

	if resolver, ok := s.(abstract.DependencyResolver); ok {
	attemptResolution:
		for {
			// attempt to resolve
			resolver.ResolveDependencies(a)

			// check the dependencies, if any are nil we try again
			for _, dependency := range resolver.Dependencies() {
				if dependency == nil {
					time.Sleep(time.Millisecond * 10)
					continue attemptResolution
				}
			}

			// everything is resolved we can now initialize
			if l, ok := s.(abstract.HasLogger); ok {
				l.Logger().Debug().Msg("dependencies satisfied ...")
			}

			// if we get here then all the deps are non-nil
			// and we know we have resolved everything
			break
		}
	}

	// everything is resolved we can now initialize
	if l, ok := s.(abstract.HasLogger); ok {
		l.Logger().Info().Msg("initializing ...")
	}

	s.Init(a)
}

func (a *App) Quit() {
	a.quit <- os.Interrupt
}

func (a *App) RemoveService(s abstract.Service) {
	for idx, svc := range a.services {
		if s != svc {
			continue
		}

		a.log.Info().Msgf("removing '%s' service", s.Name())
		a.services = append(a.services[:idx], a.services[idx+1:]...)

		break
	}
}

func (a *App) Run() {
	_ = <-a.quit
	a.log.Info().Msg("exiting")
}
