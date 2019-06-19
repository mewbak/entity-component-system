package controllers

import (
	"github.com/davidsbond/game/internal/controller"
	"github.com/davidsbond/game/internal/input"
	"github.com/davidsbond/game/internal/scene"
)

type (
	movementController struct {
		xPtr  *int
		yPtr  *int
		speed int
	}
)

func NewMovementController(xPtr, yPtr *int, speed int) controller.Controller {
	return &movementController{
		xPtr:  xPtr,
		yPtr:  yPtr,
		speed: speed,
	}
}

func (c *movementController) Update(state *input.State, info *scene.Info) error {
	if state.Keyboard["W"] && *c.yPtr > 0 {
		*c.yPtr -= c.speed
	}

	if state.Keyboard["A"] && *c.xPtr > 0 {
		*c.xPtr -= c.speed
	}

	if state.Keyboard["S"] && *c.yPtr < info.Height {
		*c.yPtr += c.speed
	}

	if state.Keyboard["D"] && *c.xPtr < info.Width {
		*c.xPtr += c.speed
	}

	return nil
}
