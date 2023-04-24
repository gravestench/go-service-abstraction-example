package main

import (
	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
	"github.com/gravestench/go-service-abstraction-example/pkg/app"
	"github.com/gravestench/go-service-abstraction-example/pkg/backup_restore"
	config "github.com/gravestench/go-service-abstraction-example/pkg/config_file_manager"
	"github.com/gravestench/go-service-abstraction-example/pkg/flags_distributor"
	"github.com/gravestench/go-service-abstraction-example/pkg/router"
	"github.com/gravestench/go-service-abstraction-example/pkg/webserver"
)

func main() {
	a := app.New()

	for _, service := range []abstract.Service{
		// distributes flags to services, only gives the flags which the service is interested in
		&flags_distributor.Service{},

		// knows how to back-up/restore other services which implement the requisite interfaces.
		&backup_restore.Service{},

		// common config service; a single file with a doubly-nested json object.
		&config.Service{},

		// routing service, looks for services that implement the requisite interfaces for initializing routes
		&router.Service{},

		// The web server, works together with the router and config service
		// eg. can check if the port has been changed in the config at runtime and restart the server on that
		// new port without having to rebuild the routes or anything, since there is a separate service that does that.
		&webserver.Service{},
	} {
		a.AddServices(service)
	}

	a.Run()
}
