package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type TreeNode struct {
    Subject string
    URL     string
    Items   []TreeNode
}

type EchoMsgTreeData struct {
    Actions  []ToolbarAction
    AreaName string
    Tree     []TreeNode
}

func EchoMsgTreeView(data EchoMsgTreeData) g.Node {
    return Div(
        Toolbar(data.Actions...),
        Zone(data.AreaName,
            Ul(treeItems(data.Tree)),
        ),
    )
}

func treeItems(nodes []TreeNode) g.Node {
    return g.Map(nodes, func(n TreeNode) g.Node {
        return Li(
            A(Href(n.URL), g.Text(n.Subject)),
            g.If(len(n.Items) > 0, Ul(treeItems(n.Items))),
        )
    })
}
