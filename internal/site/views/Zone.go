package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func Zone(header string, body ...g.Node) g.Node {
    return Div(Class("section"),
        Div(Class("section-header"), g.Text(header)),
        Div(Class("section-body"), g.Group(body)),
    )
}
