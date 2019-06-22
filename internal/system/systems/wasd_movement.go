package systems

import (
	"github.com/davidsbond/game/internal/component/components"
	"github.com/davidsbond/game/internal/input"
	"github.com/davidsbond/game/internal/scene"
	"github.com/hajimehoshi/ebiten"
)

func WASDMovement(_ *ebiten.Image, state *input.State, scene *scene.Scene) error {
	for _, e := range scene.GetEntities() {
		cmps := e.GetComponentsOfType(components.TypeWASDControl)

		for _, cmp := range cmps {
			ctrl := cmp.(*components.WASDControl)

			updatePosition(state.Keyboard, ctrl, scene)
		}
	}

	return nil
}

func updatePosition(state input.KeyBoardState, ctrl *components.WASDControl, scene *scene.Scene) {
	minX, minY, maxX, maxY := scene.GetBoundaries()

	if state["W"] && ctrl.Position.Y > minY {
		ctrl.Position.Y -= ctrl.Speed
	}

	if state["S"] && ctrl.Position.Y < maxY {
		ctrl.Position.Y += ctrl.Speed
	}

	if state["A"] && ctrl.Position.X > minX {
		ctrl.Position.X -= ctrl.Speed
	}

	if state["D"] && ctrl.Position.X < maxX {
		ctrl.Position.X += ctrl.Speed
	}
}
