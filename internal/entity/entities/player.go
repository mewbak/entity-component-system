package entities

import (
	"github.com/davidsbond/game/internal/component/components"
	"github.com/davidsbond/game/internal/entity"
)

func NewPlayer() *entity.Entity {
	pl := entity.New("player")

	pos := components.NewPosition()
	wasd := components.NewWASDControl(pos, 1)

	pl.AddComponent("position", pos)
	pl.AddComponent("wasd-control", wasd)

	return pl
}
