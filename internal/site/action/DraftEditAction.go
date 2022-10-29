package action

import (
	"fmt"
	"github.com/gorilla/mux"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/pkg/mapper"
	"net/http"
)

type DraftEditAction struct {
	Action
}

func NewDraftEditAction() *DraftEditAction {
	return new(DraftEditAction)
}

func (self DraftEditAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
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
	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget().SetClass("container")
	vBox.Add(container)

	containerVBox := widgets2.NewVBoxWidget()
	container.AddWidget(containerVBox)

	section := widgets2.NewSectionWidget()

	mainView := self.makeMainView(newDraft)

	section.SetWidget(mainView)
	containerVBox.Add(section)

	bw.SetWidget(vBox)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
	}

}

func (self DraftEditAction) makeMainView(newDraft *mapper.Draft) widgets2.IWidget {

	urlManager := um.RestoreUrlManager(self.GetRegistry())

	draftEditCompleteAddr := urlManager.CreateUrl("/draft/{draft_index}/edit/complete").
		SetParam("draft_index", newDraft.GetUUID()).
		Build()
	composeForm := widgets2.NewFormWidget()
	composeForm.SetAction(draftEditCompleteAddr)
	composeForm.SetMethod("POST")

	items := widgets2.NewVBoxWidget()

	items.Add(widgets2.NewFormButtonWidget().SetType("submit").SetName("action").SetValue("save").SetTitle("Save"))
	items.Add(widgets2.NewFormButtonWidget().SetType("submit").SetName("action").SetValue("delivery").SetTitle("Sent"))
	items.Add(widgets2.NewFormButtonWidget().SetType("submit").SetName("action").SetValue("remove").SetTitle("Remove"))

	/* Area name */
	if newDraft.IsEchoMail() {
		areaInput := widgets2.NewFormInputWidget()
		areaInput.SetTitle("Area")
		areaInput.SetName("area")
		areaInput.SetValue(newDraft.GetArea())
		areaInput.SetDisable(true)
		items.Add(areaInput)
	} else {
		items.Add(widgets2.NewFormInputWidget().SetTitle("ToAddr").SetName("to_addr").SetValue(newDraft.GetToAddr()))
	}

	/* User name */
	if newDraft.IsEchoMail() {
		areaTo := widgets2.NewFormInputWidget()
		areaTo.SetTitle("ToName")
		areaTo.SetName("to")
		areaTo.SetValue(newDraft.GetTo())
		items.Add(areaTo)
	} else {
		items.Add(widgets2.NewFormInputWidget().SetTitle("ToName").SetName("to").SetValue(newDraft.GetTo()))
	}

	items.Add(widgets2.NewFormInputWidget().SetTitle("Subject").SetName("subject").SetValue(newDraft.GetSubject()))

	items.Add(widgets2.NewFormTextWidget().SetName("body").SetValue(newDraft.GetBody()))

	composeForm.SetWidget(items)

	return composeForm
}
