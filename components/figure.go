package components

import "github.com/arcade55/htma"

// Header creates the main site header component.
func Figure() htma.Element {
	return htma.Div().ClassAttr("portfolio").AddChild(
		htma.Figure().ClassAttr("featured").AddChild(
			htma.Img().SrcAttr("static/images/monkey.jpg").AltAttr("monkey"),
			htma.Figcaption().Text("Monkey")),
		htma.Figure().AddChild(
			htma.Img().SrcAttr("static/images/eagle.jpg").AltAttr("monkey"),
			htma.Figcaption().Text("Eagle")),
		htma.Figure().ClassAttr("featured").AddChild(
			htma.Img().SrcAttr("static/images/bird.jpg").AltAttr("bird"),
			htma.Figcaption().Text("Bird")),
		htma.Figure().AddChild(
			htma.Img().SrcAttr("static/images/bear.jpg").AltAttr("bear"),
			htma.Figcaption().Text("Bear")),
		htma.Figure().ClassAttr("featured").AddChild(
			htma.Img().SrcAttr("static/images/swan.jpg").AltAttr("swan"),
			htma.Figcaption().Text("Swan")),
		htma.Figure().AddChild(
			htma.Img().SrcAttr("static/images/elephants.jpg").AltAttr("elephants)"),
			htma.Figcaption().Text("Elephants")),
		htma.Figure().AddChild(
			htma.Img().SrcAttr("static/images/owl.jpg").AltAttr("owl)"),
			htma.Figcaption().Text("Owl")),
	)
}
