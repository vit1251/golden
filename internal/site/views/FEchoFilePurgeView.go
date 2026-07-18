package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type FEchoFilePurgeData struct {
    Actions         []ToolbarAction
    AreaName        string
    ArchivedCount   int
    ActionURL       string
}

func FEchoFilePurgeView(data FEchoFilePurgeData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone("Purge file area",
            Form(Method("POST"), Action(data.ActionURL),
                P(g.Textf("Permanently delete %d archived files in '%s'?", data.ArchivedCount, data.AreaName)),
                Div(Class("toolbar"),
                    Button(Type("submit"), Class("toolbar-btn"), g.Text("Purge")),
                ),
            ),
        ),
    )
}
