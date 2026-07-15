package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type MsgHeader struct {
    Hash     string
    From     string
    To       string
    Subject  string
    Date     string     // уже отформатированная дата
    IsNew    bool
    ViewURL  string
}

type EchoMsgIndexData struct {
    Actions    []ToolbarAction
    AreaName   string
    Messages   []MsgHeader
    Pagination PaginationData
}

func EchoMsgIndexView(data EchoMsgIndexData) g.Node {
    return Div(Class("echo-list"),
        Toolbar(data.Actions...),
        g.Map(data.Messages, func(m MsgHeader) g.Node {
            return echoMsgRow(m)
        }),
        Pagination(data.Pagination),
    )
}

func echoMsgRow(m MsgHeader) g.Node {
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
}
