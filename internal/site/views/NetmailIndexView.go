package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type NetmailMsgHeader struct {
    Hash      string
    From      string
    Subject   string
    Date      string
    IsNew     bool
    ViewURL   string
}

type NetmailIndexData struct {
    Actions    []ToolbarAction
    Messages   []NetmailMsgHeader
    Pagination PaginationData
}

func NetmailIndexView(data NetmailIndexData) g.Node {
    return Div(Class("echo-list"),
        Toolbar(data.Actions...),
        g.Map(data.Messages, func(m NetmailMsgHeader) g.Node {
            classes := "echo-row"
            if m.IsNew { classes += " echo-row-new" }
            return A(Class(classes), Href(m.ViewURL),
                Div(Class("echo-row-indicator"),
                    g.If(m.IsNew, Icon("unread", 16)),
                ),
                Div(Class("echo-row-from"), g.Text(m.From)),
                Div(Class("echo-row-subject"), g.Text(m.Subject)),
                Div(Class("echo-row-date"), g.Text(m.Date)),
            )
        }),
        Pagination(data.Pagination),
    )
}