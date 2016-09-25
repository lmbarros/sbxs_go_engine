package engine

import (
	"log"

	"github.com/lmbarros/sbxs_go_engine/backend"
)

type EngineSettings struct {
	logger *log.Logger
}

// Initialize initializes the engine.
func Initialize(settings EngineSettings) {
	backend.InitializeAllegro()
}

// Shutdown shutds
func Shutdown() {
	backend.ShutdownAllegro()
}
