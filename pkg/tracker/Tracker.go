package tracker

import (
	cmn "github.com/vit1251/golden/internal/common"
	"github.com/vit1251/golden/internal/utils"
	"github.com/vit1251/golden/pkg/queue"
	"log"
	"os"
	"path"
	"time"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type Tracker struct {
	registry.Service
}

func NewTracker(r *registry.Container) *Tracker {
	newTracker := new(Tracker)
	newTracker.SetRegistry(r)
	return newTracker
}

func (self Tracker) Track() {

	trackerStart := time.Now()
	log.Printf("Start tracker session")

	err1 := self.ProcessInbound()
	if err1 != nil {
		log.Printf("err = %+v", err1)
	}
	err2 := self.ProcessOutbound()
	if err2 != nil {
		log.Printf("err = %+v", err2)
	}

	log.Printf("Stop tracker session")
	elapsed := time.Since(trackerStart)

	log.Printf("Tracker session: %+v", elapsed)
}

func (self *Tracker) ProcessInbound() error {

	queueManager := queue.RestoreQueueManager(self.GetRegistry())

	/* New mailer inbound */
	mi := queueManager.GetMailerInbound()

	/* Scan inbound */
	items, err2 := mi.Scan()
	if err2 != nil {
		return err2
	}
	log.Printf("items = %+v", items)

	for _, item := range items {
		if item.Type == queue.TypeTICmail {
			log.Printf("Tracker: TIC packet: name = %s", item.Name)
			if err := self.processTICmail(item); err != nil {
				log.Printf("Tracker: process TIC with error: err = %+v", err)
			}
		} else {
			// TODO - message about skip ...
		}
	}

	return nil
}

func (self *Tracker) ProcessOutbound() error {
	return nil
}

func (self *Tracker) processTICmail(item queue.FileEntry) error {

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	fileMapper := mapperManager.GetFileMapper()
	fileAreaMapper := mapperManager.GetFileAreaMapper()
	//TODO - statMapper := mapperManager.GetStatMapper()

	/* Parse */
	newTicParser := NewTicParser(self.GetRegistry())
	tic, err1 := newTicParser.ParseFile(item.AbsolutePath)
	if err1 != nil {
		return err1
	}
	log.Printf("tic = %+v", tic)

	areaName := tic.GetArea()

	/* Search area */
	fa, err1 := fileAreaMapper.GetAreaByName(areaName)
	if err1 != nil {
		return err1
	}

	/* Prepare area directory */
	boxDirectory := cmn.GetFilesDirectory()
	inboundDirectory := cmn.GetInboundDirectory()

	areaLocation := path.Join(boxDirectory, areaName)
	err2 := os.MkdirAll(areaLocation, 0755)
	if err2 != nil {
		log.Printf("Fail on MkdirAll: err = %+v", err2)
	}

	/* Create area */
	if fa == nil {
		/* Prepare area */
		newFa := mapper.NewFileArea()
		newFa.SetName(areaName)
		newFa.Path = areaLocation
		/* Create area */
		if err := fileAreaMapper.CreateFileArea(newFa); err != nil {
			log.Printf("Fail CreateFileArea on fileMapper: area = %s err = %+v", areaName, err)
			return err
		}
	}

	/* Create new path */
	indexName := utils.IndexHelper_makeUUID()
	inboxTicLocation := path.Join(inboundDirectory, tic.GetFile())
	areaFileLocation := path.Join(areaLocation, indexName)
	log.Printf("inboxTicLocation = %s areaFileLocation = %s", inboxTicLocation, areaFileLocation)

	/* Move */
	if err := os.Rename(inboxTicLocation, areaFileLocation); err != nil {
		log.Printf("Fail on Rename: err = %+v", err)
	}

	/* Prepare ornginal name */
	var orig_name string = tic.GetLFile()
	if orig_name == "" {
		log.Printf("Tracker: no Long name is exists. Using DOS compatible name.")
		orig_name = tic.GetFile()
	}

	/* Register file */
	newFile := mapper.NewFile()
	newFile.SetArea(tic.GetArea())
	newFile.SetDesc(tic.GetDesc())
	newFile.SetUnixTime(tic.GetUnixTime())
	newFile.SetFile(indexName)
	newFile.SetOrigName(orig_name)
	fileMapper.RegisterFile(*newFile)

	/* Register status */
	//TODO - statMapper.RegisterInFile(tic.GetFile())

	/* Move TIC */
	areaTicLocation := path.Join(areaLocation, item.Name)
	log.Printf("Move %+v -> %+v", item.AbsolutePath, areaTicLocation)
	if err := os.Rename(item.AbsolutePath, areaTicLocation); err != nil {
		log.Printf("Fail on Rename: err = %+v", err)
	}

	return nil
}
