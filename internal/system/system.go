// Package system contains types for managing systems used by
// the game.
package system

import (
	"github.com/davidsbond/game/internal/scene"
	"github.com/davidsbond/game/pkg/input"
	"github.com/hajimehoshi/ebiten"
)

type (
	System func(screen *ebiten.Image, state *input.State, scene *scene.Scene) error
)
