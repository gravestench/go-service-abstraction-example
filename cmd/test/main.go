package main

import (
	"github.com/gravestench/go-service-abstraction-example/pkg/app"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/config"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/router"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/user"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/webserver"
)

func main() {
	a := app.New()

	a.AddServices(
		&config.Service{},
		&user.Service{},
		&router.Service{},
		&webserver.Service{},
	)

	a.Run()
}