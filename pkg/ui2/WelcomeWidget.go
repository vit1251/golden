package ui2

type WelcomeView struct {

}

func NewWelcomeView() *WelcomeView {
	return new(WelcomeView)
}

func (ww *WelcomeView) Render(cs *ConnState) {

	cs.t.ResetAttr()
	cs.t.cursorhome()
	cs.t.ED2()

	var margin = (cs.t.Width - 66) / 2

	cs.t.SetAttr(F_BLUE)
	cs.scr.DrawLineY( 1,"─")

	cs.t.SetAttr(F_YELLOW)

	cs.scr.WriteStringXY(margin, 3, " ::::::::   ::::::::  :::        :::::::::  :::::::::: ::::    :::")
	cs.scr.WriteStringXY(margin, 4, ":+:    :+: :+:    :+: :+:        :+:    :+: :+:        :+:+:   :+:")
	cs.scr.WriteStringXY(margin, 5, "+:+        +:+    +:+ +:+        +:+    +:+ +:+        :+:+:+  +:+")
	cs.scr.WriteStringXY(margin, 6, ":#:        +#+    +:+ +#+        +#+    +:+ +#++:++#   +#+ +:+ +#+")
	cs.scr.WriteStringXY(margin, 7, "+#+   +#+# +#+    +#+ +#+        +#+    +#+ +#+        +#+  +#+#+#")
	cs.scr.WriteStringXY(margin, 8, "#+#    #+# #+#    #+# #+#        #+#    #+# #+#        #+#   #+#+#")
	cs.scr.WriteStringXY(margin, 9, " ########   ########  ########## #########  ########## ###    ####")

	var margin2 = (cs.t.Width - 56) / 2

	cs.t.SetAttr(F_RED)

	cs.scr.WriteStringXY(margin2, 12,":::::::::   :::::::: ::::::::::: ::::    ::: :::::::::::")
	cs.scr.WriteStringXY(margin2, 13,":+:    :+: :+:    :+:    :+:     :+:+:   :+:     :+:    ")
	cs.scr.WriteStringXY(margin2, 14,"+:+    +:+ +:+    +:+    +:+     :+:+:+  +:+     +:+    ")
	cs.scr.WriteStringXY(margin2, 15,"+#++:++#+  +#+    +:+    +#+     +#+ +:+ +#+     +#+    ")
	cs.scr.WriteStringXY(margin2, 16,"+#+        +#+    +#+    +#+     +#+  +#+#+#     +#+    ")
	cs.scr.WriteStringXY(margin2, 17,"#+#        #+#    #+#    #+#     #+#   #+#+#     #+#    ")
	cs.scr.WriteStringXY(margin2, 18,"###         ######## ########### ###    ####     ###    ")

	cs.t.SetAttr(F_BLUE)
	cs.scr.DrawLineY( 20,"─")

//	authorWidget.AddLine(fmt.Sprintf("Version: %s", "1.2.13"))
//	authorWidget.AddLine("Contributors: \n")
//	authorWidget.AddLine(" Vitold Sedyshev \n")
//	authorWidget.AddLine(" Sergey Anohin \n")

//	authorWidget.RenderXY(5, 15, cs.scr)

//	cs.scr.DrawLine("▒")

	var margin4 int = (cs.t.Width - 23) / 2
	cs.scr.WriteStringXY(margin4, 23, "Press ENTER to continue")

	cs.t.SetAttr(F_BLUE)
	cs.scr.DrawLineY( 25,"─")

}

func (ww *WelcomeView) ProcessEvent(cs *ConnState, event *TerminalEvent) {

	if event.Type == TerminalKey && event.Key == "ENTER" {
		cs.activeView = NewAreaWidget()
	}

}
