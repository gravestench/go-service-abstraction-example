package main

import (
	"github.com/faiface/mainthread"

	"github.com/gravestench/go-service-abstraction-example/pkg/app"
	"github.com/gravestench/go-service-abstraction-example/pkg/scenes/menu"
	"github.com/gravestench/go-service-abstraction-example/pkg/scenes/sprite_test"
	"github.com/gravestench/go-service-abstraction-example/pkg/services/input"
	"github.com/gravestench/go-service-abstraction-example/pkg/services/mode"
	"github.com/gravestench/go-service-abstraction-example/pkg/services/renderer"
	"github.com/gravestench/go-service-abstraction-example/pkg/services/spritesheet"
	"github.com/gravestench/go-service-abstraction-example/pkg/services/texture"
	"github.com/gravestench/go-service-abstraction-example/pkg/services/update"
)

func main() {
	a := app.New()

	a.AddServices(
		//&flags_manager.Service{},
		&renderer.Service{},
		&input.Service{},
		&update.Service{},
		&mode.Service{},
		&texture.Service{},
		&spritesheet.Service{},

		&menu.Scene{},
		&sprite_test.Scene{},
	)

	mainthread.Run(a.Run)
}
