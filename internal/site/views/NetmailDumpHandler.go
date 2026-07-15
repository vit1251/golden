package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type NetmailDumpData struct {
    Actions []ToolbarAction
    Dump    string
}

func NetmailDumpView(data NetmailDumpData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone("Raw packet",
            Pre(Class("echo-msg-dump-preview"), g.Text(data.Dump)),
        ),
    )
}
