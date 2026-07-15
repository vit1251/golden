package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/config"
    "github.com/vit1251/golden/pkg/registry"
)

type SettingsHandler struct {
    registry *registry.Container
}

func NewSettingsHandler(registry *registry.Container) *SettingsHandler {
    return &SettingsHandler{
	registry: registry,
    }
}

func (h *SettingsHandler) makeSections(c *config.Config) []views.SettingsSection {
    return []views.SettingsSection{
        {
            Name: "Connection",
            Params: []views.SettingsParam{
		{Section: "main", Name: "Address", Value: c.Main.Address, Summary: "Your point address, e.g. 2:5020/1024.11" },
		{Section: "main", Name: "Password", Value: c.Main.Password, Summary: "Binkp password for uplink authentication" },
		{Section: "main", Name: "Link", Value: c.Main.Link, Summary: "Your uplink node address, e.g. 2:5020/1024" },
		{Section: "main", Name: "NetAddr", Value: c.Main.NetAddr, Summary: "Uplink Binkp host:port, e.g. f1024.n5020.z2.binkp.net:24554" },
            },
        },
        {
            Name: "Profile",
            Params: []views.SettingsParam{
                {Section: "main", Name: "RealName", Value: c.Main.RealName, Summary: "Your real name in Latin script, e.g. Dmitri Kamenski"},
    		{Section: "main", Name: "Country", Value: c.Main.Country, Summary: "Your country"},
	        {Section: "main", Name: "City", Value: c.Main.City, Summary: "Your city"},
	        {Section: "main", Name: "Origin", Value: c.Main.Origin, Summary: "Origin line appended to all outgoing messages"},
	        {Section: "main", Name: "TearLine", Value: c.Main.TearLine, Summary: "Personal tearline (signature) below message body"},
	        {Section: "main", Name: "StationName", Value: c.Main.StationName, Summary: "Your system name, announced during Binkp handshake"},
            },
        },
        {
            Name: "Messages",
	    Params: []views.SettingsParam{
	        {Section: "netmail", Name: "Charset", Value: c.Netmail.Charset, Summary: "Default charset for outgoing messages (CP866, UTF-8, KOI8-R, etc.)"},
		{Section: "echomail", Name: "Charset", Value: c.Echomail.Charset, Summary: "Default charset for outgoing messages (CP866, UTF-8, KOI8-R, etc.)"},
            },
	},
	{
            Name: "Polling",
	    Params: []views.SettingsParam{
    		{Section: "mailer", Name: "Interval", Value: c.Mailer.Interval, Summary: "Auto-poll interval in minutes (0 = manual only)"},
            },
        },
    }
}

func (h *SettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    // Шаг 0. Пролвайдеры данных
    configManager := config.RestoreConfigManager(h.registry)
    newConfig := configManager.GetConfig()

    // Шаг 1. Конструируем параметры конфигурации
    sections := h.makeSections(newConfig)

    // Шаг 2. Рендерим страницу
    err := views.Page("Settings", views.SettingsView(views.SettingsData{
        Actions: []views.ToolbarAction{},
        Sections: sections,
    })).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
