package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type FEchoAreaCreateData struct {
    Actions   []ToolbarAction
    ActionURL string
}

func FEchoAreaCreateView(data FEchoAreaCreateData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone("New file area",
            Form(Method("POST"), Action(data.ActionURL),
                Div(Class("form-group"),
                    Label(Class("form-label"), g.Text("File area name")),
                    Input(Type("text"), Name("fileecho"), Class("form-input")),
                ),
                Div(Class("toolbar"),
                    Button(Type("submit"), Class("toolbar-btn"), g.Text("Save")),
                ),
            ),
        ),
    )
}
