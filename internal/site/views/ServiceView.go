package views

import (
    commonfunc "github.com/vit1251/golden/internal/common"
    g "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

type DayStat struct {
    Date     string
    Sessions int
    Duration string
    Volume   string
    MsgsRX   int
    MsgsTX   int
    FilesRX  int
    FilesTX  int
    Errors   int
}

type ServiceViewData struct {
    DailyStats []DayStat
    SysVersion string
    SysUptime string
}

func ServiceView(data ServiceViewData) g.Node {
    return Div(
        mailerSection(data.DailyStats),
        sysSection(data.SysVersion, data.SysUptime),
    )
}

func mailerSection(stats []DayStat) g.Node {
    zoneText := "Передача почты (мейлер)" // TODO - использовать i18n вызов здесь позднее
    return Zone(zoneText,
        mailerDailyTable(stats),
    )
}

func mailerDailyTable(stats []DayStat) g.Node {
    return Table(Class("service-table"),
        THead(
            Tr(
                Th(g.Text("Date")),
                Th(Class("num"), g.Text("Sessions")),
                Th(Class("num"), g.Text("Duration")),
                Th(Class("num"), g.Text("Volume")),
                Th(Class("num"), g.Text("Msgs")),
                Th(Class("num"), g.Text("Files")),
                Th(Class("num"), g.Text("Errors")),
            ),
        ),
        TBody(
            g.Map(stats, func(s DayStat) g.Node {
                return Tr(
                    Td(g.Text(s.Date)),
                    Td(Class("num"), g.Textf("%d", s.Sessions)),
                    Td(Class("num"), g.Text(s.Duration)),
                    Td(Class("num"), g.Text(s.Volume)),
                    Td(Class("num"), g.Textf("%d/%d", s.MsgsRX, s.MsgsTX)),
                    Td(Class("num"), g.Textf("%d/%d", s.FilesRX, s.FilesTX)),
                    Td(Class("num"), g.If(s.Errors > 0, Class("error")), g.Textf("%d", s.Errors)),
                )
            }),
        ),
    )
}

func sysSection(version string, uptime string) g.Node {
    systemText := "Системная информация"
    return Zone(systemText,
        Div(g.Textf("Golden Point v%s", version)),
        Div(g.Textf("Platform: %s/%s", commonfunc.GetPlatform(), commonfunc.GetArch())),
        Div(g.Textf("Uptime: %s", uptime)),
    )
}
