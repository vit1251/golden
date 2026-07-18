package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type ZipEntry struct {
    Name    string
    Comment string
    Size    string
}

type FEchoFileViewData struct {
    Actions    []ToolbarAction
    OrigName   string
    Desc       string
    Origin     string
    From       string
    To         string
    DiskSize   string
    Crc        string
    DiskPath   string
    ImageURL   string
    ZipComment string
    ZipFiles   []ZipEntry
}

func FEchoFileViewView(data FEchoFileViewData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone(data.OrigName,
            Div(Class("echo-msg-view-header-wrapper"),
                Div(Class("echo-msg-view-header"),
                    headerRow("Name:", data.OrigName),
		    headerRow("Description:", data.Desc),
		    headerRow("Origin:", data.Origin),
		    headerRow("From:", data.From),
		    headerRow("To:", data.To),
		    headerRow("Size:", data.DiskSize),
		    headerRow("CRC:", data.Crc),
                ),
            ),
            g.If(data.ImageURL != "", Div(Class("echo-msg-view-body"),
                Img(Class("preview"), Src(data.ImageURL)),
            )),
            g.If(data.ZipComment != "", Div(Class("echo-msg-view-body"),
                Pre(g.Text(data.ZipComment)),
            )),
            g.If(len(data.ZipFiles) > 0, Div(Class("echo-msg-view-body"),
                Div(g.Map(data.ZipFiles, func(z ZipEntry) g.Node {
                    return Div(g.Textf("%s - %s (%s)", z.Name, z.Comment, z.Size))
                })),
            )),
        ),
    )
}
