package bbs

import (
	"fmt"
	"github.com/vit1251/golden/pkg/echomail"
	"unicode"
)

type AreaListView struct {
	Widget
	activeIndex int
}

func NewAreaWidget() *AreaListView {
	return new(AreaListView)
}

type Col struct {
	DefaultWidth int
	Width        int
	Auto         bool
	Type         rune
}

type AreaListWidget struct {
	Cols []*Col
}

/**
 * Отрисовываем меню выбора области
 */
func (av *AreaListView) Render(cs *ConnState) {

	cs.t.ResetAttr()
	cs.t.cursorhome()
	cs.t.ED2()

	cs.t.ResetAttr()
	cs.t.SetAttr(B_BLUE)
	cs.scr.DrawLineY( 1,"─")

	var areaManager *echomail.AreaManager
	cs.container.Invoke(func(am *echomail.AreaManager) {
		areaManager = am
	})

	/* Get message area */
	areas, err1 := areaManager.GetAreas()
	if err1 != nil {
		cs.activeView = nil
		return
	}

	// Search maximum area name size
	var areaListEchoMax int
	for _, areas := range areas {
		areaName := areas.Name()
		areaNameSize := len(areaName)
		if areaListEchoMax < areaNameSize {
			areaListEchoMax = areaNameSize
		}
	}

	// Prepare pattern
	aw := new(AreaListWidget)

	areaListFormat := "AM D CPUN E G " // Default
	//areaListFormat := "AM E CPUN G " // Layout without descs
	//areaListFormat := "ME D CPUN G" // Alternative layout

	for _, code := range areaListFormat {
		if unicode.IsDigit(code) {
			// TODO
		}
		if unicode.IsSpace(code) {
			if code == ' ' {
				newColumn := Col{
					DefaultWidth: 1,
					Width:        1,
					Type:         code,
					//Allign: Right,
				}
				aw.Cols = append(aw.Cols, &newColumn)
			}
		}
		if unicode.IsLetter(code) {
			//Определитель  Описание                                    Ширина поля
			//по умолчанию
			//A             Номер области                               4
			//M             Символ маркировки                           1
			//D             Описание                                    динамическая
			//C             Количество сообщений                        6
			//P             Символ личной почты (+)                     1
			//U             Количество непрочитанных/новых сообщений    6
			//N             Изменение после последнего сканирования (*) 1
			//E             Имя области                                 AreaListEchoMax
			//G             Имя группы                                  {0,1,3}
			if code == 'A' {
				newColumn := Col{
					DefaultWidth: 4,
					Width:        4,
					Type:         code,
					//Allign: Right,
				}
				aw.Cols = append(aw.Cols, &newColumn)
			}
			if code == 'M' {
				newColumn := Col{
					DefaultWidth: 1,
					Width:        1,
					Type:         code,
					//Allign: Right,
				}
				aw.Cols = append(aw.Cols, &newColumn)
			}
			if code == 'D' {
				newColumn := Col{
					DefaultWidth: 0,
					Width:        0,
					Auto:         true,
					Type:         code,
					//Allign: Right,
				}
				aw.Cols = append(aw.Cols, &newColumn)
			}
			if code == 'C' {
				newColumn := Col{
					DefaultWidth: 6,
					Width:        6,
					Type:         code,
					//Allign: Right,
				}
				aw.Cols = append(aw.Cols, &newColumn)
			}
			if code == 'U' {
				newColumn := Col{
					DefaultWidth: 6,
					Width:        6,
					Type:         code,
					//Allign: Right,
				}
				aw.Cols = append(aw.Cols, &newColumn)
			}
			if code == 'E' {
				newColumn := Col{
					DefaultWidth: areaListEchoMax,
					Width:        areaListEchoMax,
					Type:         code,
					//Allign: Right,
				}
				aw.Cols = append(aw.Cols, &newColumn)
			}
		}
	}

	// Precalculate dynamic
	var newWidth = cs.t.Width
	for _, col := range aw.Cols {
		newWidth -= col.DefaultWidth
	}
	fmt.Printf("auto = %d\n", newWidth)
	for _, col := range aw.Cols {
		if col.Auto {
			col.Width = newWidth
		}
	}

	for areaIndex, area := range areas {

		var row string
		for _, col := range aw.Cols {
			if col.Type == ' ' {
				for i := 0; i < col.Width; i++ {
					row += " "
				}
			}
			if col.Type == 'A' {
				// Номер области
				row += fmt.Sprintf("%*d", col.Width, areaIndex)
			}
			if col.Type == 'M' {
				// Символ маркировки
				if area.NewMessageCount > 0 {
					row += "+"
				} else {
					row += " "
				}
			}
			if col.Type == 'E' {
				// Имя области
				row += fmt.Sprintf("%-*s", col.Width, area.Name())
			}
			if col.Type == 'D' {
				// Имя области
				row += fmt.Sprintf("%-*s", col.Width, area.Summary)
			}
			if col.Type == 'C' {
				// C             Количество сообщений                        6
				row += fmt.Sprintf("%*d", col.Width, area.MessageCount)
			}
			if col.Type == 'P' {
				// P             Символ личной почты (+)                     1
			}
			if col.Type == 'U' {
				//U             Количество непрочитанных/новых сообщений    6
				row += fmt.Sprintf("%*d", col.Width, area.NewMessageCount)
			}

		}

		if av.activeIndex == areaIndex {
			cs.t.SetAttr(B_RED)
		} else {
			cs.t.SetAttr(B_BLACK)
		}
		cs.scr.WriteStringXY(1, 2 + areaIndex, row)

		cs.t.ResetAttr()
		cs.t.SetAttr(B_BLUE)
		cs.scr.DrawLineY( cs.t.Height, "─")

	}


}

func (av *AreaListView) getAreaByIndex(cs *ConnState, idx int) *echomail.Area {
	var areaManager *echomail.AreaManager
	cs.container.Invoke(func(am *echomail.AreaManager) {
		areaManager = am
	})

	/* Get message area */
	areas, err1 := areaManager.GetAreas()
	if err1 != nil {
		cs.activeView = nil
		return nil
	}

	for areaIndex, area := range areas {
		if idx == areaIndex {
			return area
		}
	}

	return nil
}

func (av *AreaListView) ProcessEvent(cs *ConnState, event *TerminalEvent) {

	if event.Type == TerminalKey && event.Key == "ESC" {
		cs.activeView = NewWelcomeView()
	} else
	if event.Type == TerminalKey && event.Key == "HOME" {
		av.activeIndex = 0
	} else
	if event.Type == TerminalKey && event.Key == "END" {
		av.activeIndex = 0
	} else
	if event.Type == TerminalKey && event.Key == "UP" {
		if av.activeIndex > 0 {
			av.activeIndex -= 1
		}
	} else if event.Type == TerminalKey && event.Key == "DOWN" {
		av.activeIndex += 1
	} else if event.Type == TerminalKey && event.Key == "ENTER" {
		area := av.getAreaByIndex(cs, av.activeIndex)
		cs.activeArea = area.Name()
		cs.activeView = NewMessageView()
	} else {
		fmt.Printf("AreaListView: event = %v\n", event)
	}

}
