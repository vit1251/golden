package bbs

import (
	"fmt"
	"github.com/vit1251/golden/pkg/echomail"
	"github.com/vit1251/golden/pkg/msg"
	"strings"
)

type MessageView struct {
	View
	activeIndex int
}

func NewMessageView() *MessageView {
	mw := new(MessageView)
	mw.activeIndex = 0
	return mw
}

func (mw *MessageView) Render(cs *ConnState) {

	var areaManager *echomail.AreaManager
	var messageManager *echomail.MessageManager
	cs.container.Invoke(func(am *echomail.AreaManager, mm *echomail.MessageManager) {
		areaManager = am
		messageManager = mm
	})

	//
	area, err1 := areaManager.GetAreaByName(cs.activeArea)
	if err1 != nil {
		return
	}

	cs.t.ResetAttr()
	cs.t.cursorhome()
	cs.t.ED2()

	areaName := area.GetName()
	msgHeaders, err2 := messageManager.GetMessageHeaders(areaName)
	if err2 != nil {
		return
	}

	for i, msgHeader := range msgHeaders {

		if i == mw.activeIndex {

			if msgHeader.ViewCount > 0 {
				cs.t.SetAttr(F_WHITE)
			} else {
				cs.t.SetAttr(F_YELLOW)
			}
			if i == mw.activeIndex {
				cs.t.SetAttr(B_RED)
			} else {
				cs.t.SetAttr(B_BLACK)
			}

			cs.t.ResetAttr()

			cs.t.SetAttr(F_BLUE)
			cs.scr.DrawLineY(1, "─")
			cs.t.SetAttr(F_YELLOW)
			cs.scr.WriteStringXY(3, 1, fmt.Sprintf("%s", area.Summary ))
			cs.t.SetAttr(F_YELLOW)
			cs.scr.WriteStringXY(cs.t.Width - len(areaName) - 3, 1, fmt.Sprintf("%s", areaName ))

			marker := " "
			if msgHeader.ViewCount == 0 {
				marker = "*"
			}

			cs.t.SetAttr(F_WHITE)
			cs.scr.WriteStringXY(2, 2, fmt.Sprintf("Msg  : %d of %d %s %s", i+1, len(msgHeaders), msgHeader.ID, marker))
			cs.scr.WriteStringXY(2, 3, fmt.Sprintf("From : %s", msgHeader.From))
			cs.scr.WriteStringXY(2, 4, fmt.Sprintf("To   : %s", msgHeader.To))
			cs.scr.WriteStringXY(2, 5, fmt.Sprintf("Subj : %s", msgHeader.Subject))

			cs.scr.WriteStringXY(50, 3, fmt.Sprintf("%s", msgHeader.DateWritten))

			cs.t.SetAttr(F_BLUE)
			cs.scr.DrawLineY(6, "─")

			//row := fmt.Sprintf(pattern, newSign, msgHeader.Subject, msgHeader.From, msgHeader.To, msgHeader.DateWritten )

			cs.t.ResetAttr()

			msgHash := msgHeader.Hash
			msg2, _ := messageManager.GetMessageByHash(areaName, msgHash)

			content := msg2.GetContent()

			mtp := msg.NewMessageTextProcessor()
			mtp.Prepare(content)



			rows := strings.Split(content, "\r")
			for idx, row := range rows {
				msgLine := row
				var max = 7 + idx
				if max < cs.t.Height {

					mlp := msg.NewMessageLineParser()
					ml := mlp.Parse(msgLine)

					fmt.Printf("ml = %+v\n", ml)

					var level = ml.QuoteLevel
					if level > 0 {
						if level % 2 == 0 {
							cs.t.SetAttr(F_RED)
						} else {
							cs.t.SetAttr(F_GREEN)
						}
					} else {
						cs.t.ResetAttr()
					}

					cs.scr.WriteStringXY(0, max, msgLine)
				}
			}
		}

	}

}

func (mw *MessageView) markMessageByIndex(cs *ConnState, activeIndex int) error {

	var areaManager *echomail.AreaManager
	var messageManager *echomail.MessageManager
	cs.container.Invoke(func(am *echomail.AreaManager, mm *echomail.MessageManager) {
		areaManager = am
		messageManager = mm
	})

	//
	newArea, err1 := areaManager.GetAreaByName(cs.activeArea)
	if err1 != nil {
		return nil
	}

	areaName := newArea.GetName()
	msgHeaders, err2 := messageManager.GetMessageHeaders(areaName)
	if err2 != nil {
		return nil
	}

	for idx, msgHeader := range msgHeaders {
		if idx == activeIndex {
			msgHash := msgHeader.Hash
			messageManager.ViewMessageByHash(areaName, msgHash)
			return nil
		}
	}

	return nil
}

func (mw *MessageView) ProcessEvent(cs *ConnState, event *TerminalEvent) {

	if event.Type == TerminalKey && event.Key == "INSERT" {
		// TODO - activeView := NewComposeEchomailView()
	} else if event.Type == TerminalKey && event.Key == "DELETE" {
		// TODO - active overlay dialog widget confirm delete ...
		// Remove.Message
	} else if event.Type == TerminalKey && event.Key == "HOME" {
		mw.activeIndex = 0
	} else if event.Type == TerminalKey && event.Key == "LEFT" {
		mw.activeIndex -= 1
	} else if event.Type == TerminalKey && event.Key == "RIGHT" {
		mw.activeIndex += 1
	} else if event.Type == TerminalKey && event.Key == "ESC" {
		cs.activeView = NewAreaWidget()
	} else if event.Type == TerminalKey && event.Key == "SPACE" {
		mw.markMessageByIndex(cs, mw.activeIndex)
	}

}
