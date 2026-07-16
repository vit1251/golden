package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type FEchoAreaUpdateData struct {
    Actions   []ToolbarAction
    AreaName  string
    Summary   string
    Charset   string
    ActionURL string
}

func FEchoAreaUpdateView(data FEchoAreaUpdateData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone("Edit: "+data.AreaName,
	    Form(Method("POST"), Action(data.ActionURL),
                Div(Class("form-group"),
                    Label(Class("form-label"), g.Text("Summary")),
                    Input(Type("text"), Name("summary"), Class("form-input"), Value(data.Summary)),
                ),
                Div(Class("form-group"),
                    Label(Class("form-label"), g.Text("Charset")),
                    Input(Type("text"), Name("charset"), Class("form-input"), Value(data.Charset)),
                ),
                Div(Class("toolbar"),
                    Button(Type("submit"), Class("toolbar-btn"), g.Text("Save")),
                ),
            ),
        ),
    )
}
