package scenes

import (
	"github.com/davidsbond/game/internal/entity/entities"
	"github.com/davidsbond/game/internal/scene"
)

// TestScene generates a scene for testing
func TestScene() (*scene.Scene, error) {
	sc := scene.New("test_scene", 300, 300)

	pl, err := entities.NewPlayer()

	if err != nil {
		return sc, err
	}

	sc.AddEntity(pl)
	return sc, nil
}
