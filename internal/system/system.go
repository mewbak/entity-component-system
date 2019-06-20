// Package system contains types for managing systems used by
// the game.
package system

import (
	"github.com/davidsbond/game/pkg/input"
	"github.com/davidsbond/game/pkg/scene"
	"github.com/hajimehoshi/ebiten"
)

type (
	System interface {
		Run(screen *ebiten.Image, state *input.State, scene *scene.Scene) error
	}
)
