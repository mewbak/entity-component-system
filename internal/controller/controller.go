// Package controller contains types for managing controllers used by
// the game.
package controller

import (
	"github.com/davidsbond/game/internal/input"
	"github.com/davidsbond/game/internal/scene"
)

type (
	// The Controller interface defines methods for a controller that can modify the state
	// of an entity based on the given input and scene information
	Controller interface {
		// Update the entity using this controller
		Update(*input.State, *scene.Info) error
	}
)
