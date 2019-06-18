package scenes

import (
	"image/color"

	"github.com/davidsbond/game/internal/input"
	"github.com/davidsbond/game/internal/scene"
	"github.com/hajimehoshi/ebiten"
)

type (
	mainMenu struct {
		squareX int
		squareY int
	}
)

func MainMenu() (scene.Scene, error) {
	mm := &mainMenu{}
	return mm, mm.load()
}

func (mm *mainMenu) load() error {
	return nil
}

func (mm *mainMenu) Draw(screen *ebiten.Image) error {
	screen.Set(mm.squareX, mm.squareY, color.White)

	return nil
}

func (mm *mainMenu) Update(state *input.State) (scene.Scene, error) {
	mm.squareX = state.Mouse.X
	mm.squareY = state.Mouse.Y

	return nil, nil
}
