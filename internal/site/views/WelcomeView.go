package views

import (
    "strings"

    commonfunc "github.com/vit1251/golden/internal/common"
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func WelcomeView() g.Node {
    return Div(Class("container"),
	renderVerpic(),
	renderProductVersion(),
	renderSourceCode(),
	renderContributors(),
    )
}

func renderVerpic() g.Node {
    version := "1_2_19" // TODO - перевести на версию из `commonfunc.GetVersion()`
    srcLogo := "/static/Dog_" + version + ".png"
    return Img(
	Src(srcLogo),
	Class("welcome-img"),
	Alt("Golden Point"),
    )
}

func renderProductVersion() g.Node {
    return Div(
	Style("padding-bottom: 32px"),
	Div(Class("welcome-header gold-text"),
	    Style("padding-bottom: 8px"),
	    g.Text("Golden Point"),
	),
	Div(Class("welcome-version"),
	    g.Textf("Version %s", commonfunc.GetVersion()),
	),
    )
}

func renderSourceCode() g.Node {
    return section("Source code and developing", "welcome-source",
	A(
	    Href("https://github.com/vit1251/golden"),
	    Class("welcome-source-link"),
	    g.Text("https://github.com/vit1251/golden"),
	),
    )
}

func renderContributors() g.Node {
    var names []string
    for _, c := range commonfunc.GetContributors() {
	names = append(names, c.Name)
    }
    return section("Contributors", "welcome-contributor-header",
	Div(Class("welcome-contributor-list"),
	    Style("text-align: center"),
	    g.Text(strings.Join(names, ", ")),
	),
    )
}

func section(title, titleClass string, body ...g.Node) g.Node {
    children := append(
	[]g.Node{
	    Style("padding-bottom: 32px"),
	    Div(Class(titleClass), Style("padding-bottom: 8px"), g.Text(title)),
	},
	body...,
    )
    return Div(children...)
}
