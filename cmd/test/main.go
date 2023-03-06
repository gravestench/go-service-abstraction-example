package main

import (
	"github.com/gravestench/go-service-abstraction-example/pkg/app"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/config"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/database"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/flags_manager"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/router"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/session_middleware"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/static_assets"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/user"
	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/webserver"
)

func main() {
	a := app.New()

	a.AddServices(
		&flags_manager.Service{},
		&config.Service{},
		&user.Service{},
		&router.Service{},
		&webserver.Service{},
		&database.Service{},
		&session_middleware.Service{},
		&static_assets.Service{},
	)

	a.Run()
}
