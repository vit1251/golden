package views

import (
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type SettingsParam struct {
    Section string
    Name    string
    Value   string
    Summary string
}

type SettingsSection struct {
    Name   string
    Params []SettingsParam
}

type SettingsData struct {
    Actions  []ToolbarAction
    Sections []SettingsSection
}

func SettingsView(data SettingsData) g.Node {
    return Div(
        Toolbar(data.Actions...),
	Form(Method("POST"), Action("/settings"),
    	    g.Map(data.Sections, func(s SettingsSection) g.Node {
        	return Zone(s.Name,
            	    g.Map(s.Params, func(p SettingsParam) g.Node {
                	return Div(Class("form-group"),
                    	    Label(Class("form-label"),
        		    	    g.Attr("for", p.Section+"."+p.Name),
        	                g.Text(p.Name),
            	            ),
                	    Input(Class("form-input"), Type("text"),
                    	        g.Attr("id", p.Section+"."+p.Name),
                		g.Attr("name", p.Section+"."+p.Name),
                	        Value(p.Value),
                	        Title(p.Summary),
                    	    ),
                	)
            	    }),
        	)
    	    }),
    	    Div(Class("toolbar"),
        	Button(Type("submit"), Class("toolbar-btn"), g.Text("Save")),
                Button(Type("reset"), Class("toolbar-btn"), g.Text("Discard")),
	    ),
	),
    )
}
