package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type EchoMsgViewData struct {
    Actions  []ToolbarAction
    AreaName string
    From     string
    To       string
    Subject  string
    Date     string
    Body     string  // pre-rendered HTML
}

func headerRow(name, value string) g.Node {
    return Div(
        Span(Class("echo-msg-view-header-name"), g.Text(name+" ")),
        Span(Class("echo-msg-view-header-value"), g.Text(value)),
    )
}

func EchoMsgViewView(data EchoMsgViewData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone(data.AreaName,
            Div(Class("echo-msg-view-header-wrapper"),
                Div(Class("echo-msg-view-header"),
                    headerRow("Area:", data.AreaName),
                    headerRow("From:", data.From),
                    headerRow("To:", data.To),
                    headerRow("Subject:", data.Subject),
                    headerRow("Date:", data.Date),
                ),
            ),
            Div(Class("echo-msg-view-body"), g.Raw(data.Body)),
        ),
    )
}
