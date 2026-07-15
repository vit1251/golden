package views

import (
    "strconv"
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type EchoAreaUpdateData struct {
    Actions   []ToolbarAction
    ActionURL string
    AreaName  string
    Summary   string
    Charset   string
    SortOrder int64
}

func EchoAreaUpdateView(data EchoAreaUpdateData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone("Edit: " + data.AreaName,
            Form(Method("POST"), Action(data.ActionURL),
                Div(Class("form-group"),
                    Label(Class("form-label"), g.Text("Summary")),
                    Input(Class("form-input"), Type("text"), Name("summary"), Value(data.Summary)),
                ),
                Div(Class("form-group"),
                    Label(Class("form-label"), g.Text("Charset")),
                    Input(Class("form-input"), Type("text"), Name("charset"), Value(data.Charset)),
                ),
                Div(Class("form-group"),
                    Label(Class("form-label"), g.Text("Sort order")),
                    Input(Class("form-input"), Type("number"), Name("order"),
                        Value(strconv.FormatInt(data.SortOrder, 10))),
                ),
                Div(Class("toolbar"),
                    Button(Type("submit"), Class("toolbar-btn"), g.Text("Save")),
                ),
            ),
        ),
    )
}
