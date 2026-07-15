package views

import (
    "strconv"
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func Icon(id string, size int) g.Node {
    return SVG(
        g.Attr("viewBox", "0 0 24 24"),
        Width(strconv.Itoa(size)),
        Height(strconv.Itoa(size)),
        g.El("use", Href("/static/sprite.svg#"+id)),
    )
}
