// Package game contains the entry point of the game, including the main
// update and drawing logic.
package game

import (
	"github.com/davidsbond/game/internal/config"
	"github.com/davidsbond/game/internal/input"
	"github.com/davidsbond/game/internal/scene"
	"github.com/davidsbond/game/internal/scene/scenes"
	"github.com/hajimehoshi/ebiten"
)

type (
	// The Game type encompasses the entire game being ran.
	Game struct {
		config *config.Config
		title  string
		scene  scene.Scene
	}
)

// New creates a new game based on the provided configuration.
func New(cnf *config.Config) *Game {
	return &Game{
		config: cnf,
		title:  "A game",
	}
}

// Start the game.
func (g *Game) Start() error {
	// On start, load the main menu
	sc, err := scenes.MainMenu()

	if err != nil {
		return err
	}

	g.scene = sc

	ebiten.SetFullscreen(g.config.FullScreen)
	return ebiten.Run(g.run, g.config.Width, g.config.Height, 1, g.title)
}

func (g *Game) run(screen *ebiten.Image) error {
	if err := g.update(); err != nil {
		return err
	}

	return g.draw(screen)
}

func (g *Game) update() error {
	// Update the current scene using the latest input state
	newScene, err := g.scene.Update(
		input.GetState(),
		&scene.Info{
			Width:  g.config.Width,
			Height: g.config.Height,
		},
	)

	if err != nil {
		return err
	}

	// If the update method returned a new scene, set it as the
	// current scene.
	if newScene != nil {
		g.scene = newScene
	}

	return nil
}

func (g *Game) draw(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return g.scene.Draw(screen)
}
