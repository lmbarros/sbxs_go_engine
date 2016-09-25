package main

import (
	"github.com/dradtke/go-allegro/allegro"

	engine "github.com/lmbarros/sbxs_go_engine"
)

func main() {
	allegro.Run(func() {
		engine.Initialize()
		defer engine.Shutdown()
	})
}
