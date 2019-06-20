package systems

import (
	"github.com/davidsbond/game/internal/component/components"
	"github.com/davidsbond/game/internal/system"
	"github.com/davidsbond/game/pkg/input"
	"github.com/davidsbond/game/pkg/scene"
)

type (
	wasdMovement struct {
	}
)

func NewWASDMovement() system.System {
	return &wasdMovement{}
}

func (m *wasdMovement) Run(state *input.State, scene *scene.Scene) {
	for _, e := range scene.GetEntities() {
		cmps := e.GetComponentsOfType(components.TypeWASDControl)

		for _, cmp := range cmps {
			ctrl := cmp.(*components.WASDControl)

			m.updatePosition(state.Keyboard, ctrl, scene)
		}
	}
}

func (m *wasdMovement) updatePosition(state input.KeyBoardState, ctrl *components.WASDControl, scene *scene.Scene) {
	if state["W"] {
		ctrl.Position.Y -= ctrl.Speed
	}

	if state["S"] {
		ctrl.Position.Y += ctrl.Speed
	}

	if state["A"] {
		ctrl.Position.Y -= ctrl.Speed
	}

	if state["D"] {
		ctrl.Position.Y += ctrl.Speed
	}
}
