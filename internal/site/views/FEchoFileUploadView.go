package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type FEchoFileUploadData struct {
    Actions   []ToolbarAction
    ActionURL string
    To        string
    AreaName  string
}

func FEchoFileUploadView(data FEchoFileUploadData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone("Upload to "+data.AreaName,
            Form(Method("POST"), Action(data.ActionURL),
                EncType("multipart/form-data"),
                Div(Class("form-group"),
                    Label(Class("form-label"), g.Text("File")),
                    Input(Type("file"), Name("file"), Class("form-input")),
                ),
                Div(Class("form-group"),
                    Label(Class("form-label"), g.Text("To")),
                    Input(Type("text"), Name("to"), Class("form-input"), Value(data.To)),
                ),
                Div(Class("form-group"),
                    Label(Class("form-label"), g.Text("Desc")),
                    Input(Type("text"), Name("desc"), Class("form-input")),
                ),
                Div(Class("form-group"),
                    Label(Class("form-label"), g.Text("LDesc")),
                    Textarea(Name("ldesc"), Class("form-input")),
                ),
                Div(Class("toolbar"),
                    Button(Type("submit"), Class("toolbar-btn"), g.Text("Send")),
                ),
            ),
        ),
    )
}
