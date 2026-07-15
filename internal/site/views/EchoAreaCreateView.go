package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type EchoAreaCreateData struct {
    Actions   []ToolbarAction
    ActionURL string
}

func EchoAreaCreateView(data EchoAreaCreateData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone("New area",
            Form(Class("form"), Action(data.ActionURL), Method("POST"),
                formField("echoname", "Area name", "RU.ANEKDOT"),
                formField("charset", "Charset", "CP866"),
                Button(Type("submit"), Class("btn btn-primary"), g.Text("Save")),
            ),
        ),
    )
}

func formField(name string, label string, placeholder string) g.Node {
    return Div(Class("form-group"),
        Label(Class("form-label"), g.Text(label)),
        Input(Type("text"), Name(name), Class("form-input"), Placeholder(placeholder)),
    )
}
