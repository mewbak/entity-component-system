// Package system contains types for managing systems used by
// the game.
package system

import (
	"github.com/davidsbond/game/internal/scene"
	"github.com/davidsbond/game/pkg/input"
	"github.com/hajimehoshi/ebiten"
)

type (
	// The System type is a function that updates all entities within the scene
	// based on things like user-input.
	System func(screen *ebiten.Image, state *input.State, scene *scene.Scene) error
)
