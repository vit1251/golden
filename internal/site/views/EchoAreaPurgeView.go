package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type EchoAreaPurgeData struct {
    Actions          []ToolbarAction
    AreaName         string
    ArchiveMsgCount  int
    ActionURL        string
}

func EchoAreaPurgeView(data EchoAreaPurgeData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone("Purge area",
            Form(Method("POST"), Action(data.ActionURL),
                P(g.Textf("Permanently delete %d archived messages in '%s'?", data.ArchiveMsgCount, data.AreaName)),
                Div(Class("toolbar"),
                    Button(Type("submit"), Class("toolbar-btn"), g.Text("Purge")),
                ),
            ),
        ),
    )
}