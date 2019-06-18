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

func (mm *mainMenu) Update(state *input.State, info *scene.Info) (scene.Scene, error) {
	if state.Keyboard["W"] && mm.squareY > 0 {
		mm.squareY--
	}

	if state.Keyboard["A"] && mm.squareX > 0 {
		mm.squareX--
	}

	if state.Keyboard["S"] && mm.squareY < info.Height {
		mm.squareY++
	}

	if state.Keyboard["D"] && mm.squareX < info.Width {
		mm.squareX++
	}

	return nil, nil
}
