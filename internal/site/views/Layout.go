package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/components"
    . "maragu.dev/gomponents/html"
)

func Page(title string, body ...g.Node) g.Node {
    return HTML5(HTML5Props{
	Title:    title,
	Language: "en",
	Head: []g.Node{
	    Link(Rel("stylesheet"), Href("/static/custom.css")),
	    Link(Rel("stylesheet"), Href("/static/theme_black.css")),
	    Link(Rel("stylesheet"), Href("/static/modern.css")),
	    Link(Rel("stylesheet"), Href("/static/print.css"), g.Attr("media", "print")),
	    Script(Src("/static/custom.js"), Defer()),
	},
	Body: append(
	    []g.Node{
	        renderMenu(),
	    },
	    body...,
	),
    })
}

func renderMenu() g.Node {
    return Div(Class("Header"),
	Div(Class("Header-item-group"),
	    menuLink("/", "Home", ""),
	    menuLink("/netmail", "Netmail", ""),
	    menuLink("/echo", "Echomail", ""),
	    menuLink("/file", "Files", ""),
	    menuLink("/service", "Service", ""),
	    menuLink("/twit", "Address book", ""),
	    menuLink("/draft", "Draft", ""),
	),
	Div(Class("Header-item-group"),
	    menuLink("/settings", "", "settings"),
	),
    )
}

func menuLink(href, label, icon string) g.Node {
    return A(
	Class("nav-link"), Href(href),
	Div(Class("Header-item"),
	    g.If(icon != "", Icon(icon, 24)),
	    g.If(label != "", Span(Class("tab-label"), g.Text(label))),
	    // TODO: metric badge
	),
    )
}
