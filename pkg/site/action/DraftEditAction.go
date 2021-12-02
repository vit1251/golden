package action

import (
	"fmt"
	"github.com/gorilla/mux"
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

	/* Render base wiew */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget().SetClass("container")
	vBox.Add(container)

	containerVBox := widgets.NewVBoxWidget()
	container.AddWidget(containerVBox)

	var section *widgets.SectionWidget
	if newDraft.IsEchoMail() {
		newTitle := fmt.Sprintf("Edit conference message: %s", newDraft.GetArea())
		section = widgets.NewSectionWidget().SetTitle(newTitle)
	} else {
		section = widgets.NewSectionWidget().SetTitle("Edit direct message")
	}

	composeForm := widgets.NewFormWidget().
		SetAction(fmt.Sprintf("/draft/%s/edit/complete", newDraft.GetUUID())).
		SetMethod("POST")

	composeForm.SetWidget(widgets.NewVBoxWidget().
		Add(widgets.NewFormInputWidget().SetTitle("ToName").SetName("to").SetValue(newDraft.GetTo())).
		Add(widgets.NewFormInputWidget().SetTitle("ToAddr").SetName("to_addr").SetValue(newDraft.GetToAddr())).
		Add(widgets.NewFormInputWidget().SetTitle("Subject").SetName("subject").SetValue(newDraft.GetSubject())).
		Add(widgets.NewFormTextWidget().SetName("body").SetValue(newDraft.GetBody())).
		Add(widgets.NewFormButtonWidget().SetType("submit").SetName("action").SetValue("remove").SetTitle("Remove")).
		Add(widgets.NewFormButtonWidget().SetType("submit").SetName("action").SetValue("save").SetTitle("Save")).
		Add(widgets.NewFormButtonWidget().SetType("submit").SetName("action").SetValue("delivery").SetTitle("Delivery message")))

	section.SetWidget(composeForm)
	containerVBox.Add(section)

	bw.SetWidget(vBox)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
	}

}
