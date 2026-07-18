package handler

import (
    "fmt"
    "net/http"
    "sort"
    "time"

    commonfunc "github.com/vit1251/golden/internal/common"
    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/registry"
)

type ServiceHandler struct {
    registry *registry.Container
}

func NewServiceHandler(registry *registry.Container) *ServiceHandler {
    return &ServiceHandler{
	registry: registry,
    }
}

type dayAccum struct {
    sessions int
    duration time.Duration
    errors   int
    filesRX  int
    filesTX  int
}

func formatDuration(d time.Duration) string {
    h := int(d.Hours())
    m := int(d.Minutes()) % 60
    s := int(d.Seconds()) % 60
    if h > 0 {
	return fmt.Sprintf("%dh %dm", h, m)
    }
    if m > 0 {
	return fmt.Sprintf("%dm %ds", m, s)
    }
    return fmt.Sprintf("%ds", s)
}

func aggregateByDay(sessions []mapper.StatMailer) []views.DayStat {
    dayMap := make(map[string]*dayAccum)

    for _, s := range sessions {
	t := time.UnixMilli(s.SessionStart)
	key := t.Format("2006-01-02")

	if dayMap[key] == nil {
	    dayMap[key] = &dayAccum{}
	}

	d := dayMap[key]
	d.sessions++
	d.duration += time.Duration(s.SessionStop - s.SessionStart) * time.Millisecond
	d.filesRX += s.FileRXcount
	d.filesTX += s.FileTXcount

	// Для ошибок смотрим статус
	if s.Status == "Complete: RX = RxOk TX = TxOk" {
	} else {
	    d.errors++
	}
    }

    var result []views.DayStat
    for date, acc := range dayMap {
	var errStr string
	if acc.errors > 0 {
	    errStr = formatDuration(acc.duration)
	}
	_ = errStr

	result = append(result, views.DayStat{
	    Date:     date,
	    Sessions: acc.sessions,
	    Duration: formatDuration(acc.duration),
	    Volume:   "-", // TODO: добавить когда появится в MailerReport
	    MsgsRX:   0,   // TODO: добавить когда появится в MailerReport
	    MsgsTX:   0,
	    FilesRX:  acc.filesRX,
	    FilesTX:  acc.filesTX,
	    Errors:   acc.errors,
	})
    }

    sort.Slice(result, func(i, j int) bool {
	return result[i].Date < result[j].Date
    })

    return result
}

func formatUptime(d time.Duration) string {
    days := int(d.Hours()) / 24
    hours := int(d.Hours()) % 24
    mins := int(d.Minutes()) % 60
    return fmt.Sprintf("%dd %dh %dm", days, hours, mins)
}

func (h *ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    mapperManager := mapper.RestoreMapperManager(h.registry)
    statMailerMapper := mapperManager.GetStatMailerMapper()

    sessions, err := statMailerMapper.GetMailerSummary()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    daily := aggregateByDay(sessions)

    data := views.ServiceViewData{
        DailyStats: daily,
        SysVersion: commonfunc.GetVersion(),
        SysUptime: formatUptime(commonfunc.GetUptime()),
    }

    err = views.Page("Service", views.ServiceView(data)).Render(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
