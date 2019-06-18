// Package input contains methods and types for obtaining the current state
// of the user's input, including keyboard and mouse.
package input

import "github.com/hajimehoshi/ebiten"

type (
	// The State type represents the current state of the keyboard
	// and mouse.
	State struct {
		Keyboard KeyBoardState // The current state of the keyboard
		Mouse    MouseState    // The current state of the mouse
	}

	// The MouseState type represents the current state of the mouse
	MouseState struct {
		Left   bool // Whether or not the left mouse button is pressed
		Right  bool // Whether or not the right mouse button is pressed
		Middle bool // Whether or not the middle mouse button is pressed
		X      int  // The location of the cursor on the X-axis
		Y      int  // The location of the cursor on the Y-axis
	}

	KeyBoardState map[string]bool
)

// GetState obtains the current state of the keyboard and mouse buttons.
func GetState() *State {
	state := &State{
		Keyboard: make(map[string]bool),
		Mouse:    getMouseState(),
	}

	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		state.Keyboard[k.String()] = ebiten.IsKeyPressed(k)
	}

	return state
}

func getMouseState() MouseState {
	x, y := ebiten.CursorPosition()

	return MouseState{
		Left:   ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft),
		Right:  ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight),
		Middle: ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle),
		X:      x,
		Y:      y,
	}
}
