package scenes

import (
	"github.com/davidsbond/game/internal/entity"
	"github.com/davidsbond/game/internal/entity/entities"
	"github.com/davidsbond/game/internal/input"
	"github.com/davidsbond/game/internal/scene"
	"github.com/hajimehoshi/ebiten"
)

type (
	mainMenu struct {
		player entity.Entity
	}
)

func MainMenu() (scene.Scene, error) {
	mm := &mainMenu{}
	return mm, mm.load()
}

func (mm *mainMenu) load() error {
	mm.player = entities.NewPlayer()

	return nil
}

func (mm *mainMenu) Draw(screen *ebiten.Image) error {
	if err := mm.player.Draw(screen); err != nil {
		return err
	}

	return nil
}

func (mm *mainMenu) Update(state *input.State, info *scene.Info) (scene.Scene, error) {
	if err := mm.player.Update(state, info); err != nil {
		return nil, err
	}

	return nil, nil
}
