package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type FEchoFileRow struct {
    OrigName  string
    Desc      string
    Date      string
    ViewURL   string
    IsNew     bool
    IsMissing bool
}

type FEchoFileIndexData struct {
    Actions  []ToolbarAction
    AreaName string
    Files    []FEchoFileRow
}

func FEchoFileIndexView(data FEchoFileIndexData) g.Node {
    return Div(Class("echo-list"),
        Toolbar(data.Actions...),
        g.Map(data.Files, func(f FEchoFileRow) g.Node {
            classes := "echo-row"
            if f.IsNew { classes += " echo-row-new" }
            if f.IsMissing { classes += " echo-row-missing" }
            return A(Class(classes), Href(f.ViewURL),
                Div(Class("echo-row-indicator"),
                    g.If(f.IsNew, Icon("unread", 16)),
                    g.If(f.IsMissing, Span(g.Text("[!]"))),
                ),
                Div(Class("echo-row-from"), g.Text(f.OrigName)),
                Div(Class("echo-row-subject"), g.Text(f.Desc)),
                Div(Class("echo-row-date"), g.Text(f.Date)),
            )
        }),
    )
}
