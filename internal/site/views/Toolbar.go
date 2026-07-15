package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type ToolbarAction struct {
    Label string
    URL   string
    Icon  string       // ID из sprite.svg, пусто если без иконки
}

func Toolbar(actions ...ToolbarAction) g.Node {
    if len(actions) == 0 {
        return g.Group{}
    }
    return Div(Class("toolbar"),
        g.Map(actions, func(a ToolbarAction) g.Node {
            return A(Class("toolbar-btn"), Href(a.URL),
                g.If(a.Icon != "", Icon(a.Icon, 16)),
                Span(g.Text(a.Label)),
            )
        }),
    )
}
