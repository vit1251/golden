package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type FEchoAreaIndexData struct {
    Actions []ToolbarAction
    Areas   []AreaHeader
}

func FEchoAreaIndexView(data FEchoAreaIndexData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Div(Class("area-list"),
            g.Map(data.Areas, func(a AreaHeader) g.Node {
                IsUnread := a.NewMsgCount > 0
                classes := "area-row"
                if IsUnread { classes += " area-row-new"}
                return A(Class(classes), Href(a.IndexURL),
                    Div(Class("area-row-indicator"),
                        g.If(IsUnread, Icon("unread", 16)),
                    ),
                    Div(Class("area-row-name"), g.Text(a.Name)),
                    Div(Class("area-row-summary"), g.Text(a.Summary)),
                    g.If(a.NewMsgCount > 0,
                        Span(Class("area-row-count"), g.Textf("%d", a.NewMsgCount)),
                    ),
                )
            }),
        ),
    )
}
