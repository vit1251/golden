package views

import (
    "fmt"

    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type PaginationData struct {
    CurrentPage int
    TotalPages  int
    BaseURL     string
}

func Pagination(p PaginationData) g.Node {
    if p.TotalPages <= 1 {
        return g.Group{} // не показываем если одна страница
    }

    return Div(Class("pagination"),
        A(Class("page-link"),
            g.If(p.CurrentPage > 1, Href(fmt.Sprintf("%s?page=%d", p.BaseURL, p.CurrentPage-1))),
            g.Text("← Newer"),
        ),
        Span(Class("page-info"),
            g.Textf("Page %d of %d", p.CurrentPage, p.TotalPages),
        ),
        A(Class("page-link"),
            g.If(p.CurrentPage < p.TotalPages, Href(fmt.Sprintf("%s?page=%d", p.BaseURL, p.CurrentPage+1))),
            g.Text("Older →"),
        ),
    )
}