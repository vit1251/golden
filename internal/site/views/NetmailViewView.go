package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type NetmailAttachment struct {
    Name string
    URL  string
    Size string
}

type NetmailViewData struct {
    Actions     []ToolbarAction
    From        string
    To          string
    Subject     string
    Date        string
    Body        string      // HTML from doc.HTML()
    Attachments []NetmailAttachment
}

func attachmentsSection(atts []NetmailAttachment) g.Node {
    return Div(
        Span(Class("netmail-msg-view-header-name"), g.Text("Attachments: ")),
        g.Map(atts, func(a NetmailAttachment) g.Node {
            return A(Href(a.URL), g.Text(a.Name + " (" + a.Size + ")"), Br())
        }),
    )
}

func NetmailViewView(data NetmailViewData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone("Netmail",
            Div(Class("netmail-msg-view-header-wrapper"),
                Div(Class("netmail-msg-view-header"),
                    headerRow("From:", data.From),
                    headerRow("To:", data.To),
                    headerRow("Subject:", data.Subject),
                    headerRow("Date:", data.Date),
                    g.If(len(data.Attachments) > 0, attachmentsSection(data.Attachments)),
                ),
            ),
            Div(Class("netmail-msg-view-body"), g.Raw(data.Body)),
        ),
    )
}