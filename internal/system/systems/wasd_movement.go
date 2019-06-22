package systems

import (
	"github.com/davidsbond/game/internal/component/components"
	"github.com/davidsbond/game/internal/scene"
	"github.com/davidsbond/game/pkg/input"
	"github.com/hajimehoshi/ebiten"
)

// The WASDMovement system finds components of type components.TypeWASDControl and
// updates their position based on the current state of the W, A, S and D keys on the
// keyboard.
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
