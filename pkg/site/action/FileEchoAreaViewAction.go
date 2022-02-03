package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
	"strings"
)

type FileEchoAreaViewAction struct {
	Action
}

func NewFileEchoAreaViewAction() *FileEchoAreaViewAction {
	return new(FileEchoAreaViewAction)
}

func (self *FileEchoAreaViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	fileAreaMapper := mapperManager.GetFileAreaMapper()
	fileMapper := mapperManager.GetFileMapper()

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)
	newFile := vars["file"]
	log.Printf("file = %v", newFile)

	/* Get message area */
	area, err1 := fileAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	/**/
	file, err2 := fileMapper.GetFileByFileName(echoTag, newFile)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetFileByFileName on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Update view counter */
	err3 := fileMapper.ViewFileByFileName(area.GetName(), newFile)
	if err3 != nil {
		response := fmt.Sprintf("Fail on ViewFileByFileName on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render page */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	vBox.Add(container)

	containerVBox := widgets.NewVBoxWidget()

	container.AddWidget(containerVBox)

	/* Context actions */
	actionMenu := widgets.NewActionMenuWidget()
	actionMenu.Add(widgets.NewMenuAction().
		SetLink(fmt.Sprintf("/file/%s/tic/%s/download", area.GetName(), newFile)).
		SetClass("netmail-reply-action").
		SetIcon("icofont-edit").
		SetLabel("Download"))
	actionMenu.Add(widgets.NewMenuAction().
		SetLink(fmt.Sprintf("/file/%s/tic/%s/remove", area.GetName(), newFile)).
		SetClass("netmail-reply-action").
		SetIcon("icofont-edit").
		SetLabel("Remove"))

	containerVBox.Add(actionMenu)

	/* TODO - show meta here ... */
	if IsImage(file.GetFile()) {
		imageURL := fmt.Sprintf("/file/%s/tic/%s/download", area.GetName(), newFile)
		imageWidget := widgets.NewImageWidget()
		imageWidget.SetSource(imageURL)
		imageWidget.SetClass("preview")
		containerVBox.Add(imageWidget)
	}

	/* Done */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func IsImage(filename string) bool {
	upperName := strings.ToUpper(filename)
	return strings.HasSuffix(upperName, ".GIF") || strings.HasSuffix(upperName, ".JPG") || strings.HasSuffix(upperName, ".PNG")
}
