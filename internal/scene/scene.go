package scene

import (
	"github.com/davidsbond/game/internal/input"
	"github.com/hajimehoshi/ebiten"
)

type (
	// The Scene interface defines methods for updating and rendering a scene within
	// the game.
	Scene interface {
		// Draw should have all contents in the scene drawn to the provided
		// image.
		Draw(*ebiten.Image) error

		// Update to update all entities in the scene based on the provided parameters.
		// If an update requires a change of scene, it should be returned here. If an update
		// does not require a scene change, the return value of the first parameter should
		// be nil.
		Update(*input.State, *Info) (Scene, error)
	}

	// The Info type contains general information that all scenes should be aware of
	Info struct {
		Width  int // The screen width
		Height int // The screen height
	}
)
