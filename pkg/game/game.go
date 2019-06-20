// Package game contains the entry point of the game, including the main
// update and drawing logic.
package game

import (
	"github.com/davidsbond/game/internal/system"
	"github.com/davidsbond/game/internal/system/systems"
	"github.com/davidsbond/game/pkg/config"
	"github.com/davidsbond/game/pkg/input"
	"github.com/davidsbond/game/pkg/scene"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type (
	// The Game type encompasses the entire game being ran.
	Game struct {
		config  *config.Config
		title   string
		scene   *scene.Scene
		systems []system.System
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
	g.addSystem(systems.NewWASDMovement())
	g.addSystem(systems.NewSpriteRenderer())

	ebiten.SetFullscreen(g.config.FullScreen)
	return ebiten.Run(g.run, g.config.Width, g.config.Height, 1, g.title)
}

func (g *Game) run(screen *ebiten.Image) error {
	state := input.GetState()

	for _, sys := range g.systems {
		if err := sys.Run(screen, state, g.scene); err != nil {
			if err := ebitenutil.DebugPrint(screen, err.Error()); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Game) addSystem(sys system.System) {
	g.systems = append(g.systems, sys)
}
