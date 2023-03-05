package app

import (
	"math/rand"
	"time"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/logging"
)

func New() *App {
	a := &App{}
	a.init()
	return a
}

type App struct {
	quit     chan bool
	services []interface{}
	log      abstract.Logger
}

func (a *App) init() {
	if a.services != nil {
		return
	}

	rand.Seed(time.Now().UnixNano())

	a.quit = make(chan bool)
	a.services = make([]interface{}, 0)
	a.log = logging.New()
	a.services = append(a.services, a.log)

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

func (a *App) addService(s abstract.Service) {
	a.services = append(a.services, s)

	if a.log != nil {
		a.log.Info().Msgf("[APP] initializing '%s' service", s.Name())
	}

	go s.Init(&a.services)
}

func (a *App) Run() {
	_ = <-a.quit
	a.log.Info().Msg("[APP] exiting")
}
