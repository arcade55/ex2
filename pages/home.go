package pages

import (
	"ex1/components"

	"github.com/arcade55/htma"
)

func Home() htma.Element {
	// Start with the root HTML component and add the others as its children.
	return components.Html().AddChild(
		components.Head(),
		components.Body().AddChild(
			components.Figure(),
		),
	)
}
