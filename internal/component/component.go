// Package component contains types that define component behaviour
// within the game.
package component

type (
	// Type denotes a kind of component. Each implementation of the
	// Component interface should have a unique type.
	Type string

	// The Component interface contains methods all components used
	// within the game must implement.
	Component interface {
		// Returns the component's unique type.
		GetType() Type
	}
)
