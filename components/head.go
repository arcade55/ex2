package components

import "github.com/arcade55/htma"

func Head() htma.Element {

	return htma.Head().
		AddChild(
			htma.Meta().
				CharsetAttr("UTF-8"),
		).
		AddChild(
			htma.Link().
				RelAttr("stylesheet").
				HrefAttr("static/style.css"),
		)
}
