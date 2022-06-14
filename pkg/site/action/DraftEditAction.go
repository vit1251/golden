package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type DraftEditAction struct {
	Action
}

func NewDraftEditAction() *DraftEditAction {
	return new(DraftEditAction)
}

func (self DraftEditAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	draftMapper := mapperManager.GetDraftMapper()

	/* Restore draftid */
	vars := mux.Vars(r)

	/* Restore area by "echoname" key */
	draftId := vars["draftid"]

	/* Restore draft index */
	newDraft, err1 := draftMapper.GetDraftByUUID(draftId)
	if err1 != nil {
		status := fmt.Sprintf("Fail in GetDraftById on draftMapper: err = %+v", err1)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}
	if newDraft == nil {
		status := fmt.Sprintf("Fail in GetDraftById on draftMapper: err = %+v", err1)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

	/* Prepare */
	if newDraft.IsEchoMail() {
		newTo := newDraft.GetTo()
		if newTo == "" {
			newDraft.SetTo("All")
		}
	}

	/* Render base view */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget().SetClass("container")
	vBox.Add(container)

	containerVBox := widgets.NewVBoxWidget()
	container.AddWidget(containerVBox)

	section := widgets.NewSectionWidget()

	mainView := self.makeMainView(newDraft)

	section.SetWidget(mainView)
	containerVBox.Add(section)

	bw.SetWidget(vBox)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
	}

}

func (self DraftEditAction) makeMainView(newDraft *mapper.Draft) widgets.IWidget {

	urlManager := self.restoreUrlManager()

	draftEditCompleteAddr := urlManager.CreateUrl("/draft/{draft_index}/edit/complete").
		SetParam("draft_index", newDraft.GetUUID()).
		Build()
	composeForm := widgets.NewFormWidget()
	composeForm.SetAction(draftEditCompleteAddr)
	composeForm.SetMethod("POST")

	items := widgets.NewVBoxWidget()

	items.Add(widgets.NewFormButtonWidget().SetType("submit").SetName("action").SetValue("save").SetTitle("Save"))
	items.Add(widgets.NewFormButtonWidget().SetType("submit").SetName("action").SetValue("delivery").SetTitle("Sent"))
	items.Add(widgets.NewFormButtonWidget().SetType("submit").SetName("action").SetValue("remove").SetTitle("Remove"))

	/* Area name */
	if newDraft.IsEchoMail() {
		areaInput := widgets.NewFormInputWidget()
		areaInput.SetTitle("Area")
		areaInput.SetName("area")
		areaInput.SetValue(newDraft.GetArea())
		areaInput.SetDisable(true)
		items.Add(areaInput)
	} else {
		items.Add(widgets.NewFormInputWidget().SetTitle("ToAddr").SetName("to_addr").SetValue(newDraft.GetToAddr()))
	}

	/* User name */
	if newDraft.IsEchoMail() {
		areaTo := widgets.NewFormInputWidget()
		areaTo.SetTitle("ToName")
		areaTo.SetName("to")
		areaTo.SetValue(newDraft.GetTo())
		items.Add(areaTo)
	} else {
		items.Add(widgets.NewFormInputWidget().SetTitle("ToName").SetName("to").SetValue(newDraft.GetTo()))
	}

	items.Add(widgets.NewFormInputWidget().SetTitle("Subject").SetName("subject").SetValue(newDraft.GetSubject()))

	items.Add(widgets.NewFormTextWidget().SetName("body").SetValue(newDraft.GetBody()))

	composeForm.SetWidget(items)

	return composeForm
}
