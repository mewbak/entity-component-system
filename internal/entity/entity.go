// Package entity contains types for managing entities within the game.
package entity

import (
	"github.com/davidsbond/game/internal/input"
	"github.com/davidsbond/game/internal/scene"
	"github.com/hajimehoshi/ebiten"
)

type (
	// The Entity interface defines methods for updating and rendering an entity within
	// the game.
	Entity interface {
		// Draw should draw the entity to the provided screen
		Draw(*ebiten.Image) error

		// Update the current state of the entity
		Update(*input.State, *scene.Info) error
	}
)
