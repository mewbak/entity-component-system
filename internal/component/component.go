package component

type (
	Type string

	Component interface {
		GetType() Type
	}
)
