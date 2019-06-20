// Package system contains types for managing systems used by
// the game.
package system

import (
	"github.com/davidsbond/game/pkg/input"
	"github.com/davidsbond/game/pkg/scene"
)

type (
	System interface {
		Run(state *input.State, scene *scene.Scene)
	}
)
